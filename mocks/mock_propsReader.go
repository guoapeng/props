// Code generated by MockGen. DO NOT EDIT.
// Source: PropsReader.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAppConfigProperties is a mock of AppConfigProperties interface
type MockAppConfigProperties struct {
	ctrl     *gomock.Controller
	recorder *MockAppConfigPropertiesMockRecorder
}

// MockAppConfigPropertiesMockRecorder is the mock recorder for MockAppConfigProperties
type MockAppConfigPropertiesMockRecorder struct {
	mock *MockAppConfigProperties
}

// NewMockAppConfigProperties creates a new mock instance
func NewMockAppConfigProperties(ctrl *gomock.Controller) *MockAppConfigProperties {
	mock := &MockAppConfigProperties{ctrl: ctrl}
	mock.recorder = &MockAppConfigPropertiesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppConfigProperties) EXPECT() *MockAppConfigPropertiesMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockAppConfigProperties) Get(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockAppConfigPropertiesMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAppConfigProperties)(nil).Get), key)
}
