// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import zk "github.com/samuel/go-zookeeper/zk"

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Driver) Close() {
	_m.Called()
}

// Connect provides a mock function with given fields: url
func (_m *Driver) Connect(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateNode provides a mock function with given fields: path, data
func (_m *Driver) CreateNode(path string, data []byte) error {
	ret := _m.Called(path, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(path, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNode provides a mock function with given fields: path
func (_m *Driver) DeleteNode(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetConnection provides a mock function with given fields:
func (_m *Driver) GetConnection() *zk.Conn {
	ret := _m.Called()

	var r0 *zk.Conn
	if rf, ok := ret.Get(0).(func() *zk.Conn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zk.Conn)
		}
	}

	return r0
}

// GetData provides a mock function with given fields: path
func (_m *Driver) GetData(path string) ([]byte, error) {
	ret := _m.Called(path)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LockEntity provides a mock function with given fields: name
func (_m *Driver) LockEntity(name string) (string, bool) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ReleaseEntity provides a mock function with given fields: lock
func (_m *Driver) ReleaseEntity(lock string) {
	_m.Called(lock)
}

// WatchForNode provides a mock function with given fields: path
func (_m *Driver) WatchForNode(path string) (<-chan zk.Event, error) {
	ret := _m.Called(path)

	var r0 <-chan zk.Event
	if rf, ok := ret.Get(0).(func(string) <-chan zk.Event); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan zk.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
