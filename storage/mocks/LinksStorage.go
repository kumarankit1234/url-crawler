// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LinksStorage is an autogenerated mock type for the LinksStorage type
type LinksStorage struct {
	mock.Mock
}

// Add provides a mock function with given fields: link
func (_m *LinksStorage) Add(link string) {
	_m.Called(link)
}

// IsPresent provides a mock function with given fields: link
func (_m *LinksStorage) IsPresent(link string) bool {
	ret := _m.Called(link)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewLinksStorage creates a new instance of LinksStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLinksStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *LinksStorage {
	mock := &LinksStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
