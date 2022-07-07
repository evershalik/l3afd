// Code generated by MockGen. DO NOT EDIT.
// Source: mock_interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockplatformInterface is a mock of platformInterface interface.
type MockplatformInterface struct {
	ctrl     *gomock.Controller
	recorder *MockplatformInterfaceMockRecorder
}

// MockplatformInterfaceMockRecorder is the mock recorder for MockplatformInterface.
type MockplatformInterfaceMockRecorder struct {
	mock *MockplatformInterface
}

// NewMockplatformInterface creates a new mock instance.
func NewMockplatformInterface(ctrl *gomock.Controller) *MockplatformInterface {
	mock := &MockplatformInterface{ctrl: ctrl}
	mock.recorder = &MockplatformInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockplatformInterface) EXPECT() *MockplatformInterfaceMockRecorder {
	return m.recorder
}

// GetPlatform mocks base method.
func (m *MockplatformInterface) GetPlatform() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlatform")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlatform indicates an expected call of GetPlatform.
func (mr *MockplatformInterfaceMockRecorder) GetPlatform() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlatform", reflect.TypeOf((*MockplatformInterface)(nil).GetPlatform))
}
