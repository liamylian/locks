// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/study-only/go-locks (interfaces: ExpiryLockFactory)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	go_locks "github.com/study-only/go-locks"
	reflect "reflect"
	time "time"
)

// MockExpiryLockFactory is a mock of ExpiryLockFactory interface
type MockExpiryLockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockExpiryLockFactoryMockRecorder
}

// MockExpiryLockFactoryMockRecorder is the mock recorder for MockExpiryLockFactory
type MockExpiryLockFactoryMockRecorder struct {
	mock *MockExpiryLockFactory
}

// NewMockExpiryLockFactory creates a new mock instance
func NewMockExpiryLockFactory(ctrl *gomock.Controller) *MockExpiryLockFactory {
	mock := &MockExpiryLockFactory{ctrl: ctrl}
	mock.recorder = &MockExpiryLockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExpiryLockFactory) EXPECT() *MockExpiryLockFactoryMockRecorder {
	return m.recorder
}

// NewLock mocks base method
func (m *MockExpiryLockFactory) NewLock(arg0 string, arg1 time.Duration) go_locks.TryLocker {
	ret := m.ctrl.Call(m, "NewLock", arg0, arg1)
	ret0, _ := ret[0].(go_locks.TryLocker)
	return ret0
}

// NewLock indicates an expected call of NewLock
func (mr *MockExpiryLockFactoryMockRecorder) NewLock(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLock", reflect.TypeOf((*MockExpiryLockFactory)(nil).NewLock), arg0, arg1)
}