// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/mock/go/context (interfaces: WithTimeout,CancelFunc)

// Package context is a generated GoMock package.
package context

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockWithTimeout is a mock of WithTimeout interface.
type MockWithTimeout struct {
	ctrl     *gomock.Controller
	recorder *MockWithTimeoutMockRecorder
}

// MockWithTimeoutMockRecorder is the mock recorder for MockWithTimeout.
type MockWithTimeoutMockRecorder struct {
	mock *MockWithTimeout
}

// NewMockWithTimeout creates a new mock instance.
func NewMockWithTimeout(ctrl *gomock.Controller) *MockWithTimeout {
	mock := &MockWithTimeout{ctrl: ctrl}
	mock.recorder = &MockWithTimeoutMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWithTimeout) EXPECT() *MockWithTimeoutMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockWithTimeout) Do(arg0 context.Context, arg1 time.Duration) (context.Context, context.CancelFunc) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(context.CancelFunc)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockWithTimeoutMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockWithTimeout)(nil).Do), arg0, arg1)
}

// MockCancelFunc is a mock of CancelFunc interface.
type MockCancelFunc struct {
	ctrl     *gomock.Controller
	recorder *MockCancelFuncMockRecorder
}

// MockCancelFuncMockRecorder is the mock recorder for MockCancelFunc.
type MockCancelFuncMockRecorder struct {
	mock *MockCancelFunc
}

// NewMockCancelFunc creates a new mock instance.
func NewMockCancelFunc(ctrl *gomock.Controller) *MockCancelFunc {
	mock := &MockCancelFunc{ctrl: ctrl}
	mock.recorder = &MockCancelFuncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCancelFunc) EXPECT() *MockCancelFuncMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockCancelFunc) Do() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Do")
}

// Do indicates an expected call of Do.
func (mr *MockCancelFuncMockRecorder) Do() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockCancelFunc)(nil).Do))
}