// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jbell-circle/hermes/internal/controllers (interfaces: AddressController)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/jbell-circle/hermes/internal/models"
)

// MockAddressController is a mock of AddressController interface.
type MockAddressController struct {
	ctrl     *gomock.Controller
	recorder *MockAddressControllerMockRecorder
}

// MockAddressControllerMockRecorder is the mock recorder for MockAddressController.
type MockAddressControllerMockRecorder struct {
	mock *MockAddressController
}

// NewMockAddressController creates a new mock instance.
func NewMockAddressController(ctrl *gomock.Controller) *MockAddressController {
	mock := &MockAddressController{ctrl: ctrl}
	mock.recorder = &MockAddressControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddressController) EXPECT() *MockAddressControllerMockRecorder {
	return m.recorder
}

// CreateAddress mocks base method.
func (m *MockAddressController) CreateAddress(arg0 string, arg1 models.Address) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddress", arg0, arg1)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAddress indicates an expected call of CreateAddress.
func (mr *MockAddressControllerMockRecorder) CreateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddress", reflect.TypeOf((*MockAddressController)(nil).CreateAddress), arg0, arg1)
}
