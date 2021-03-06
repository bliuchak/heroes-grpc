// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import storage "github.com/bliuchak/heroes/internal/storage"

// Storager is an autogenerated mock type for the Storager type
type Storager struct {
	mock.Mock
}

// CreateHero provides a mock function with given fields: id, name
func (_m *Storager) CreateHero(id string, name string) error {
	ret := _m.Called(id, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(id, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteHero provides a mock function with given fields: id
func (_m *Storager) DeleteHero(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetHero provides a mock function with given fields: name
func (_m *Storager) GetHero(name string) (storage.Hero, error) {
	ret := _m.Called(name)

	var r0 storage.Hero
	if rf, ok := ret.Get(0).(func(string) storage.Hero); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(storage.Hero)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHeroes provides a mock function with given fields:
func (_m *Storager) GetHeroes() ([]storage.Hero, error) {
	ret := _m.Called()

	var r0 []storage.Hero
	if rf, ok := ret.Get(0).(func() []storage.Hero); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.Hero)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields:
func (_m *Storager) Status() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
