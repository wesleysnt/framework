// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cache "github.com/wesleysnt/framework/contracts/cache"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

// Add provides a mock function with given fields: key, value, t
func (_m *Driver) Add(key string, value interface{}, t time.Duration) bool {
	ret := _m.Called(key, value, t)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) bool); ok {
		r0 = rf(key, value, t)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Decrement provides a mock function with given fields: key, value
func (_m *Driver) Decrement(key string, value ...int) (int, error) {
	_va := make([]interface{}, len(value))
	for _i := range value {
		_va[_i] = value[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, value...)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...int) error); ok {
		r1 = rf(key, value...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Flush provides a mock function with given fields:
func (_m *Driver) Flush() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Forever provides a mock function with given fields: key, value
func (_m *Driver) Forever(key string, value interface{}) bool {
	ret := _m.Called(key, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, interface{}) bool); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Forget provides a mock function with given fields: key
func (_m *Driver) Forget(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields: key, def
func (_m *Driver) Get(key string, def ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, def...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(key, def...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetBool provides a mock function with given fields: key, def
func (_m *Driver) GetBool(key string, def ...bool) bool {
	_va := make([]interface{}, len(def))
	for _i := range def {
		_va[_i] = def[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...bool) bool); ok {
		r0 = rf(key, def...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetInt provides a mock function with given fields: key, def
func (_m *Driver) GetInt(key string, def ...int) int {
	_va := make([]interface{}, len(def))
	for _i := range def {
		_va[_i] = def[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, def...)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetInt64 provides a mock function with given fields: key, def
func (_m *Driver) GetInt64(key string, def ...int64) int64 {
	_va := make([]interface{}, len(def))
	for _i := range def {
		_va[_i] = def[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, ...int64) int64); ok {
		r0 = rf(key, def...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetString provides a mock function with given fields: key, def
func (_m *Driver) GetString(key string, def ...string) string {
	_va := make([]interface{}, len(def))
	for _i := range def {
		_va[_i] = def[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, def...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Has provides a mock function with given fields: key
func (_m *Driver) Has(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Increment provides a mock function with given fields: key, value
func (_m *Driver) Increment(key string, value ...int) (int, error) {
	_va := make([]interface{}, len(value))
	for _i := range value {
		_va[_i] = value[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, value...)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...int) error); ok {
		r1 = rf(key, value...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lock provides a mock function with given fields: key, t
func (_m *Driver) Lock(key string, t ...time.Duration) cache.Lock {
	_va := make([]interface{}, len(t))
	for _i := range t {
		_va[_i] = t[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 cache.Lock
	if rf, ok := ret.Get(0).(func(string, ...time.Duration) cache.Lock); ok {
		r0 = rf(key, t...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Lock)
		}
	}

	return r0
}

// Pull provides a mock function with given fields: key, def
func (_m *Driver) Pull(key string, def ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, def...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(key, def...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Put provides a mock function with given fields: key, value, t
func (_m *Driver) Put(key string, value interface{}, t time.Duration) error {
	ret := _m.Called(key, value, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) error); ok {
		r0 = rf(key, value, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remember provides a mock function with given fields: key, ttl, callback
func (_m *Driver) Remember(key string, ttl time.Duration, callback func() (interface{}, error)) (interface{}, error) {
	ret := _m.Called(key, ttl, callback)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, time.Duration, func() (interface{}, error)) interface{}); ok {
		r0 = rf(key, ttl, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Duration, func() (interface{}, error)) error); ok {
		r1 = rf(key, ttl, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RememberForever provides a mock function with given fields: key, callback
func (_m *Driver) RememberForever(key string, callback func() (interface{}, error)) (interface{}, error) {
	ret := _m.Called(key, callback)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, func() (interface{}, error)) interface{}); ok {
		r0 = rf(key, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, func() (interface{}, error)) error); ok {
		r1 = rf(key, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithContext provides a mock function with given fields: ctx
func (_m *Driver) WithContext(ctx context.Context) cache.Driver {
	ret := _m.Called(ctx)

	var r0 cache.Driver
	if rf, ok := ret.Get(0).(func(context.Context) cache.Driver); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Driver)
		}
	}

	return r0
}

type mockConstructorTestingTNewDriver interface {
	mock.TestingT
	Cleanup(func())
}

// NewDriver creates a new instance of Driver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDriver(t mockConstructorTestingTNewDriver) *Driver {
	mock := &Driver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
