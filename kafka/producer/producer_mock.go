// Code generated by MockGen. DO NOT EDIT.
// Source: producer.go

// Package producer is a generated GoMock package.
package producer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProducer is a mock of Producer interface.
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer.
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance.
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockProducer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProducer)(nil).Close))
}

// Flush mocks base method.
func (m *MockProducer) Flush(timeoutMs int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", timeoutMs)
	ret0, _ := ret[0].(int)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockProducerMockRecorder) Flush(timeoutMs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockProducer)(nil).Flush), timeoutMs)
}

// Produce mocks base method.
func (m *MockProducer) Produce(msg Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce.
func (mr *MockProducerMockRecorder) Produce(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockProducer)(nil).Produce), msg)
}

// Responses mocks base method.
func (m *MockProducer) Responses() chan Delivery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Responses")
	ret0, _ := ret[0].(chan Delivery)
	return ret0
}

// Responses indicates an expected call of Responses.
func (mr *MockProducerMockRecorder) Responses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Responses", reflect.TypeOf((*MockProducer)(nil).Responses))
}