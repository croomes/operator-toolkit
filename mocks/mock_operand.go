// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/darkowlzz/composite-reconciler/operator/v1/operand (interfaces: Operand)

// Package mocks is a generated GoMock package.
package mocks

import (
	v1 "github.com/darkowlzz/composite-reconciler/event/v1"
	operand "github.com/darkowlzz/composite-reconciler/operator/v1/operand"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOperand is a mock of Operand interface
type MockOperand struct {
	ctrl     *gomock.Controller
	recorder *MockOperandMockRecorder
}

// MockOperandMockRecorder is the mock recorder for MockOperand
type MockOperandMockRecorder struct {
	mock *MockOperand
}

// NewMockOperand creates a new mock instance
func NewMockOperand(ctrl *gomock.Controller) *MockOperand {
	mock := &MockOperand{ctrl: ctrl}
	mock.recorder = &MockOperandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOperand) EXPECT() *MockOperandMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockOperand) Delete(arg0 interface{}) (v1.ReconcilerEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(v1.ReconcilerEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockOperandMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOperand)(nil).Delete), arg0)
}

// Ensure mocks base method
func (m *MockOperand) Ensure(arg0 interface{}) (v1.ReconcilerEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ensure", arg0)
	ret0, _ := ret[0].(v1.ReconcilerEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ensure indicates an expected call of Ensure
func (mr *MockOperandMockRecorder) Ensure(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ensure", reflect.TypeOf((*MockOperand)(nil).Ensure), arg0)
}

// Name mocks base method
func (m *MockOperand) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockOperandMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockOperand)(nil).Name))
}

// ReadyCheck mocks base method
func (m *MockOperand) ReadyCheck(arg0 interface{}) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadyCheck", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadyCheck indicates an expected call of ReadyCheck
func (mr *MockOperandMockRecorder) ReadyCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadyCheck", reflect.TypeOf((*MockOperand)(nil).ReadyCheck), arg0)
}

// RequeueStrategy mocks base method
func (m *MockOperand) RequeueStrategy() operand.RequeueStrategy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequeueStrategy")
	ret0, _ := ret[0].(operand.RequeueStrategy)
	return ret0
}

// RequeueStrategy indicates an expected call of RequeueStrategy
func (mr *MockOperandMockRecorder) RequeueStrategy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequeueStrategy", reflect.TypeOf((*MockOperand)(nil).RequeueStrategy))
}

// Requires mocks base method
func (m *MockOperand) Requires() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Requires")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Requires indicates an expected call of Requires
func (mr *MockOperandMockRecorder) Requires() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Requires", reflect.TypeOf((*MockOperand)(nil).Requires))
}
