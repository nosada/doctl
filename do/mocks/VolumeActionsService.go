// Generated: please do not edit by hand

package mocks

import do "github.com/digitalocean/doctl/do"
import godo "github.com/digitalocean/godo"
import mock "github.com/stretchr/testify/mock"

// VolumeActionsService is an autogenerated mock type for the VolumeActionsService type
type VolumeActionsService struct {
	mock.Mock
}

// Attach provides a mock function with given fields: _a0, _a1
func (_m *VolumeActionsService) Attach(_a0 string, _a1 int) (*do.Action, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.Action
	if rf, ok := ret.Get(0).(func(string, int) *do.Action); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Action)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Detach provides a mock function with given fields: _a0, _a1
func (_m *VolumeActionsService) Detach(_a0 string, _a1 int) (*do.Action, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.Action
	if rf, ok := ret.Get(0).(func(string, int) *do.Action); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Action)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *VolumeActionsService) Get(_a0 string, _a1 int) (*do.Action, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.Action
	if rf, ok := ret.Get(0).(func(string, int) *do.Action); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Action)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *VolumeActionsService) List(_a0 string, _a1 *godo.ListOptions) ([]do.Action, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []do.Action
	if rf, ok := ret.Get(0).(func(string, *godo.ListOptions) []do.Action); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]do.Action)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.ListOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Resize provides a mock function with given fields: _a0, _a1, _a2
func (_m *VolumeActionsService) Resize(_a0 string, _a1 int, _a2 string) (*do.Action, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *do.Action
	if rf, ok := ret.Get(0).(func(string, int, string) *do.Action); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Action)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
