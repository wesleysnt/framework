// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	seeder "github.com/wesleysnt/framework/contracts/database/seeder"
	mock "github.com/stretchr/testify/mock"
)

// Facade is an autogenerated mock type for the Facade type
type Facade struct {
	mock.Mock
}

// Call provides a mock function with given fields: seeders
func (_m *Facade) Call(seeders []seeder.Seeder) error {
	ret := _m.Called(seeders)

	var r0 error
	if rf, ok := ret.Get(0).(func([]seeder.Seeder) error); ok {
		r0 = rf(seeders)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CallOnce provides a mock function with given fields: seeders
func (_m *Facade) CallOnce(seeders []seeder.Seeder) error {
	ret := _m.Called(seeders)

	var r0 error
	if rf, ok := ret.Get(0).(func([]seeder.Seeder) error); ok {
		r0 = rf(seeders)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSeeder provides a mock function with given fields: name
func (_m *Facade) GetSeeder(name string) seeder.Seeder {
	ret := _m.Called(name)

	var r0 seeder.Seeder
	if rf, ok := ret.Get(0).(func(string) seeder.Seeder); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(seeder.Seeder)
		}
	}

	return r0
}

// GetSeeders provides a mock function with given fields:
func (_m *Facade) GetSeeders() []seeder.Seeder {
	ret := _m.Called()

	var r0 []seeder.Seeder
	if rf, ok := ret.Get(0).(func() []seeder.Seeder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]seeder.Seeder)
		}
	}

	return r0
}

// Register provides a mock function with given fields: seeders
func (_m *Facade) Register(seeders []seeder.Seeder) {
	_m.Called(seeders)
}

// NewFacade creates a new instance of Facade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFacade(t interface {
	mock.TestingT
	Cleanup(func())
}) *Facade {
	mock := &Facade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
