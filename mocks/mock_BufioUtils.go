// Code generated by MockGen. DO NOT EDIT.
// Source: utils/BufioUtils.go

// Package mocks is a generated GoMock package.
package mocks

import (
	bufio "bufio"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockBufioUtils is a mock of BufioUtils interface
type MockBufioUtils struct {
	ctrl     *gomock.Controller
	recorder *MockBufioUtilsMockRecorder
}

// MockBufioUtilsMockRecorder is the mock recorder for MockBufioUtils
type MockBufioUtilsMockRecorder struct {
	mock *MockBufioUtils
}

// NewMockBufioUtils creates a new mock instance
func NewMockBufioUtils(ctrl *gomock.Controller) *MockBufioUtils {
	mock := &MockBufioUtils{ctrl: ctrl}
	mock.recorder = &MockBufioUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBufioUtils) EXPECT() *MockBufioUtilsMockRecorder {
	return m.recorder
}

// NewScanner mocks base method
func (m *MockBufioUtils) NewScanner(r io.Reader) *bufio.Scanner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewScanner", r)
	ret0, _ := ret[0].(*bufio.Scanner)
	return ret0
}

// NewScanner indicates an expected call of NewScanner
func (mr *MockBufioUtilsMockRecorder) NewScanner(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewScanner", reflect.TypeOf((*MockBufioUtils)(nil).NewScanner), r)
}
