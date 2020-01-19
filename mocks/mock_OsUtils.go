// Code generated by MockGen. DO NOT EDIT.
// Source: utils/OsUtils.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	os "os"
	reflect "reflect"
)

// MockOsUtils is a mock of OsUtils interface
type MockOsUtils struct {
	ctrl     *gomock.Controller
	recorder *MockOsUtilsMockRecorder
}

// MockOsUtilsMockRecorder is the mock recorder for MockOsUtils
type MockOsUtilsMockRecorder struct {
	mock *MockOsUtils
}

// NewMockOsUtils creates a new mock instance
func NewMockOsUtils(ctrl *gomock.Controller) *MockOsUtils {
	mock := &MockOsUtils{ctrl: ctrl}
	mock.recorder = &MockOsUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOsUtils) EXPECT() *MockOsUtilsMockRecorder {
	return m.recorder
}

// Getenv mocks base method
func (m *MockOsUtils) Getenv(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getenv", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Getenv indicates an expected call of Getenv
func (mr *MockOsUtilsMockRecorder) Getenv(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getenv", reflect.TypeOf((*MockOsUtils)(nil).Getenv), key)
}

// Open mocks base method
func (m *MockOsUtils) Open(name string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", name)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockOsUtilsMockRecorder) Open(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockOsUtils)(nil).Open), name)
}