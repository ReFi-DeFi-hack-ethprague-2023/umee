// Code generated by MockGen. DO NOT EDIT.
// Source: ./x/uibc/expected_keepers.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/bank/types"
	gomock "github.com/golang/mock/gomock"
	types1 "github.com/umee-network/umee/v4/x/leverage/types"
)

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// GetDenomMetaData mocks base method.
func (m *MockBankKeeper) GetDenomMetaData(ctx types.Context, denom string) (types0.Metadata, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDenomMetaData", ctx, denom)
	ret0, _ := ret[0].(types0.Metadata)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDenomMetaData indicates an expected call of GetDenomMetaData.
func (mr *MockBankKeeperMockRecorder) GetDenomMetaData(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDenomMetaData", reflect.TypeOf((*MockBankKeeper)(nil).GetDenomMetaData), ctx, denom)
}

// IterateAllDenomMetaData mocks base method.
func (m *MockBankKeeper) IterateAllDenomMetaData(ctx types.Context, cb func(types0.Metadata) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateAllDenomMetaData", ctx, cb)
}

// IterateAllDenomMetaData indicates an expected call of IterateAllDenomMetaData.
func (mr *MockBankKeeperMockRecorder) IterateAllDenomMetaData(ctx, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateAllDenomMetaData", reflect.TypeOf((*MockBankKeeper)(nil).IterateAllDenomMetaData), ctx, cb)
}

// SetDenomMetaData mocks base method.
func (m *MockBankKeeper) SetDenomMetaData(ctx types.Context, denomMetaData types0.Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDenomMetaData", ctx, denomMetaData)
}

// SetDenomMetaData indicates an expected call of SetDenomMetaData.
func (mr *MockBankKeeperMockRecorder) SetDenomMetaData(ctx, denomMetaData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDenomMetaData", reflect.TypeOf((*MockBankKeeper)(nil).SetDenomMetaData), ctx, denomMetaData)
}

// MockLeverageKeeper is a mock of LeverageKeeper interface.
type MockLeverageKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockLeverageKeeperMockRecorder
}

// MockLeverageKeeperMockRecorder is the mock recorder for MockLeverageKeeper.
type MockLeverageKeeperMockRecorder struct {
	mock *MockLeverageKeeper
}

// NewMockLeverageKeeper creates a new mock instance.
func NewMockLeverageKeeper(ctrl *gomock.Controller) *MockLeverageKeeper {
	mock := &MockLeverageKeeper{ctrl: ctrl}
	mock.recorder = &MockLeverageKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeverageKeeper) EXPECT() *MockLeverageKeeperMockRecorder {
	return m.recorder
}

// DeriveExchangeRate mocks base method.
func (m *MockLeverageKeeper) DeriveExchangeRate(ctx types.Context, denom string) types.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeriveExchangeRate", ctx, denom)
	ret0, _ := ret[0].(types.Dec)
	return ret0
}

// DeriveExchangeRate indicates an expected call of DeriveExchangeRate.
func (mr *MockLeverageKeeperMockRecorder) DeriveExchangeRate(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeriveExchangeRate", reflect.TypeOf((*MockLeverageKeeper)(nil).DeriveExchangeRate), ctx, denom)
}

// ExchangeUToken mocks base method.
func (m *MockLeverageKeeper) ExchangeUToken(ctx types.Context, uToken types.Coin) (types.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeUToken", ctx, uToken)
	ret0, _ := ret[0].(types.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExchangeUToken indicates an expected call of ExchangeUToken.
func (mr *MockLeverageKeeperMockRecorder) ExchangeUToken(ctx, uToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeUToken", reflect.TypeOf((*MockLeverageKeeper)(nil).ExchangeUToken), ctx, uToken)
}

// GetTokenSettings mocks base method.
func (m *MockLeverageKeeper) GetTokenSettings(ctx types.Context, baseDenom string) (types1.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenSettings", ctx, baseDenom)
	ret0, _ := ret[0].(types1.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenSettings indicates an expected call of GetTokenSettings.
func (mr *MockLeverageKeeperMockRecorder) GetTokenSettings(ctx, baseDenom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenSettings", reflect.TypeOf((*MockLeverageKeeper)(nil).GetTokenSettings), ctx, baseDenom)
}

// MockOracle is a mock of Oracle interface.
type MockOracle struct {
	ctrl     *gomock.Controller
	recorder *MockOracleMockRecorder
}

// MockOracleMockRecorder is the mock recorder for MockOracle.
type MockOracleMockRecorder struct {
	mock *MockOracle
}

// NewMockOracle creates a new mock instance.
func NewMockOracle(ctrl *gomock.Controller) *MockOracle {
	mock := &MockOracle{ctrl: ctrl}
	mock.recorder = &MockOracleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOracle) EXPECT() *MockOracleMockRecorder {
	return m.recorder
}

// Price mocks base method.
func (m *MockOracle) Price(ctx types.Context, denom string) (types.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Price", ctx, denom)
	ret0, _ := ret[0].(types.Dec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Price indicates an expected call of Price.
func (mr *MockOracleMockRecorder) Price(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Price", reflect.TypeOf((*MockOracle)(nil).Price), ctx, denom)
}