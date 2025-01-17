package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v5/util/sdkutil"
	"github.com/umee-network/umee/v5/x/refileverage/types"
)

// DeriveBorrowAPY derives the current borrow interest rate on a token denom
// using its supply utilization and token-specific params. Returns zero on
// invalid asset.
func (k Keeper) DeriveBorrowAPY(ctx sdk.Context, denom string) sdk.Dec {
	token, err := k.GetTokenSettings(ctx, denom)
	if err != nil {
		return sdk.ZeroDec()
	}
	return token.BaseBorrowRate
}

// AccrueAllInterest is called by EndBlock to update borrow positions.
// It accrues interest on all open borrows, increase reserves, funds
// oracle rewards, and sets LastInterestTime to BlockTime.
func (k Keeper) AccrueAllInterest(ctx sdk.Context) error {
	currentTime := ctx.BlockTime().Unix()
	prevInterestTime := k.getLastInterestTime(ctx)
	if prevInterestTime <= 0 {
		// if stored LastInterestTime is zero (or negative), either the chain has just started
		// or the genesis file has been modified intentionally. In either case, proceed as if
		// 0 seconds have passed since the last block, thus accruing no interest and setting
		// the current BlockTime as the new starting point.
		prevInterestTime = currentTime
	}

	// calculate time elapsed since last interest accrual (measured in years for APR math)
	if currentTime < prevInterestTime {
		// precaution against this and similar issues: https://github.com/tendermint/tendermint/issues/8773
		k.Logger(ctx).With("AccrueAllInterest will wait for block time > prevInterestTime").Error(
			types.ErrNegativeTimeElapsed.Error(),
			"current", currentTime,
			"prev", prevInterestTime,
		)

		// if LastInterestTime appears to be in the future, do nothing (besides logging) and leave
		// LastInterestTime at its stored value. This will repeat every block until BlockTime exceeds
		// LastInterestTime.
		return nil
	}

	yearsElapsed := sdk.NewDec(currentTime - prevInterestTime).QuoInt64(types.SecondsPerYear)
	if yearsElapsed.GTE(sdk.OneDec()) {
		// this safeguards primarily against misbehaving block time or incorrectly modified genesis states
		// which would accrue significant interest on borrows instantly. Chain will halt.
		return types.ErrExcessiveTimeElapsed.Wrapf("BlockTime: %d, LastInterestTime: %d",
			currentTime, prevInterestTime)
	}

	tokens := k.GetAllRegisteredTokens(ctx)
	newReserves := sdk.NewCoins()
	totalInterest := sdk.NewCoins()

	for _, token := range tokens {
		if token.Blacklist {
			// skip accruing interest on blacklisted assets
			continue
		}

		// interest is accrued by continuous compound interest on each denom's Interest Scalar
		scalar := k.getInterestScalar(ctx, token.BaseDenom)

		// multiply interest scalar by e^(APY*time)
		exponential := ApproxExponential(k.DeriveBorrowAPY(ctx, token.BaseDenom).Mul(yearsElapsed))
		if err := k.setInterestScalar(ctx, token.BaseDenom, scalar.Mul(exponential)); err != nil {
			return err
		}

		prevTotalBorrowed := k.getAdjustedTotalBorrowed(ctx).Mul(scalar)
		interestAccrued := prevTotalBorrowed.Mul(exponential.Sub(sdk.OneDec()))
		totalInterest = totalInterest.Add(sdk.NewCoin(
			token.BaseDenom,
			interestAccrued.TruncateInt(),
		))

		// if interest accrued on this denom is at least one base token
		if interestAccrued.GT(sdk.OneDec()) {
			// calculate new reserves gained for this denom, rounding up
			newReserves = newReserves.Add(sdk.NewCoin(
				token.BaseDenom,
				interestAccrued.Mul(token.ReserveFactor).Ceil().TruncateInt(),
			))
		}
	}

	// apply all reserve increases accumulated when iterating over denoms
	for _, coin := range newReserves {
		if err := k.setReserves(ctx, coin.Add(k.GetReserves(ctx, coin.Denom))); err != nil {
			return err
		}
	}

	// set LastInterestTime
	err := k.setLastInterestTime(ctx, currentTime)
	if err != nil {
		return err
	}

	// Because this action is not caused by a message, logging and
	// events are here instead of msg_server.go
	k.Logger(ctx).Debug(
		"interest accrued",
		"block_height", fmt.Sprintf("%d", ctx.BlockHeight()),
		"unix_time", fmt.Sprintf("%d", currentTime),
		"interest", totalInterest.String(),
		"reserved", newReserves.String(),
	)
	sdkutil.Emit(&ctx, &types.EventInterestAccrual{
		BlockHeight:   uint64(ctx.BlockHeight()),
		Timestamp:     uint64(currentTime),
		TotalInterest: totalInterest,
		Reserved:      newReserves,
	})
	return nil
}
