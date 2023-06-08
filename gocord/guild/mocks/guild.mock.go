// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	guild "github.com/fathalfath30/gocord/gocord/guild"
	mock "github.com/stretchr/testify/mock"
)

// GuildMock is an autogenerated mock type for the IGuild type
type GuildMock struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *GuildMock) Create(_a0 *guild.Guild) (*error, error) {
	ret := _m.Called(_a0)

	var r0 *error
	var r1 error
	if rf, ok := ret.Get(0).(func(*guild.Guild) (*error, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*guild.Guild) *error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*error)
		}
	}

	if rf, ok := ret.Get(1).(func(*guild.Guild) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *GuildMock) Get(id string) (*guild.Guild, error) {
	ret := _m.Called(id)

	var r0 *guild.Guild
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*guild.Guild, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *guild.Guild); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*guild.Guild)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGuildMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewGuildMock creates a new instance of GuildMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGuildMock(t mockConstructorTestingTNewGuildMock) *GuildMock {
	mock := &GuildMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
