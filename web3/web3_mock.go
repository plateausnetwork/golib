// Code generated by MockGen. DO NOT EDIT.
// Source: web3.go

// Package web3 is a generated GoMock package.
package web3

import (
	big "math/big"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
)

// MockWeb3 is a mock of Web3 interface.
type MockWeb3 struct {
	ctrl     *gomock.Controller
	recorder *MockWeb3MockRecorder
}

// MockWeb3MockRecorder is the mock recorder for MockWeb3.
type MockWeb3MockRecorder struct {
	mock *MockWeb3
}

// NewMockWeb3 creates a new mock instance.
func NewMockWeb3(ctrl *gomock.Controller) *MockWeb3 {
	mock := &MockWeb3{ctrl: ctrl}
	mock.recorder = &MockWeb3MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeb3) EXPECT() *MockWeb3MockRecorder {
	return m.recorder
}

// GetBlockNumber mocks base method.
func (m *MockWeb3) GetBlockNumber() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockNumber")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockNumber indicates an expected call of GetBlockNumber.
func (mr *MockWeb3MockRecorder) GetBlockNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockNumber", reflect.TypeOf((*MockWeb3)(nil).GetBlockNumber))
}

// GetNonce mocks base method.
func (m *MockWeb3) GetNonce(addr common.Address, blockNumber *big.Int) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", addr, blockNumber)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNonce indicates an expected call of GetNonce.
func (mr *MockWeb3MockRecorder) GetNonce(addr, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockWeb3)(nil).GetNonce), addr, blockNumber)
}

// SendRawTransaction mocks base method.
func (m *MockWeb3) SendRawTransaction(to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRawTransaction", to, amount, gasLimit, gasPrice, data)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawTransaction indicates an expected call of SendRawTransaction.
func (mr *MockWeb3MockRecorder) SendRawTransaction(to, amount, gasLimit, gasPrice, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawTransaction", reflect.TypeOf((*MockWeb3)(nil).SendRawTransaction), to, amount, gasLimit, gasPrice, data)
}

// SetAccount mocks base method.
func (m *MockWeb3) SetAccount(privateKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccount", privateKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockWeb3MockRecorder) SetAccount(privateKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockWeb3)(nil).SetAccount), privateKey)
}

// SetChainId mocks base method.
func (m *MockWeb3) SetChainId(chainId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetChainId", chainId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetChainId indicates an expected call of SetChainId.
func (mr *MockWeb3MockRecorder) SetChainId(chainId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetChainId", reflect.TypeOf((*MockWeb3)(nil).SetChainId), chainId)
}
