// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SCRestAPIHandler is an autogenerated mock type for the SCRestAPIHandler type
type SCRestAPIHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: response, numSharders, err
func (_m *SCRestAPIHandler) Execute(response map[string][]byte, numSharders int, err error) {
	_m.Called(response, numSharders, err)
}

type mockConstructorTestingTNewSCRestAPIHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewSCRestAPIHandler creates a new instance of SCRestAPIHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSCRestAPIHandler(t mockConstructorTestingTNewSCRestAPIHandler) *SCRestAPIHandler {
	mock := &SCRestAPIHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
