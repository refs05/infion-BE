// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	users "infion-BE/businesses/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateNewUser provides a mock function with given fields: domain, ctx
func (_m *Repository) CreateNewUser(domain users.DomainUser, ctx context.Context) (users.DomainUser, error) {
	ret := _m.Called(domain, ctx)

	var r0 users.DomainUser
	if rf, ok := ret.Get(0).(func(users.DomainUser, context.Context) users.DomainUser); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(users.DomainUser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.DomainUser, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: userId, ctx
func (_m *Repository) FindById(userId int, ctx context.Context) (users.DomainUser, error) {
	ret := _m.Called(userId, ctx)

	var r0 users.DomainUser
	if rf, ok := ret.Get(0).(func(int, context.Context) users.DomainUser); ok {
		r0 = rf(userId, ctx)
	} else {
		r0 = ret.Get(0).(users.DomainUser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, context.Context) error); ok {
		r1 = rf(userId, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeaderboard provides a mock function with given fields: ctx
func (_m *Repository) GetLeaderboard(ctx context.Context) ([]users.DomainUser, error) {
	ret := _m.Called(ctx)

	var r0 []users.DomainUser
	if rf, ok := ret.Get(0).(func(context.Context) []users.DomainUser); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.DomainUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsername provides a mock function with given fields: domain, ctx
func (_m *Repository) GetUsername(domain users.DomainUser, ctx context.Context) (users.DomainUser, error) {
	ret := _m.Called(domain, ctx)

	var r0 users.DomainUser
	if rf, ok := ret.Get(0).(func(users.DomainUser, context.Context) users.DomainUser); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(users.DomainUser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.DomainUser, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: domain, ctx
func (_m *Repository) Login(domain users.DomainUser, ctx context.Context) (users.DomainUser, error) {
	ret := _m.Called(domain, ctx)

	var r0 users.DomainUser
	if rf, ok := ret.Get(0).(func(users.DomainUser, context.Context) users.DomainUser); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(users.DomainUser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.DomainUser, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userDomain, ctx
func (_m *Repository) Update(userDomain *users.DomainUser, ctx context.Context) (users.DomainUser, error) {
	ret := _m.Called(userDomain, ctx)

	var r0 users.DomainUser
	if rf, ok := ret.Get(0).(func(*users.DomainUser, context.Context) users.DomainUser); ok {
		r0 = rf(userDomain, ctx)
	} else {
		r0 = ret.Get(0).(users.DomainUser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.DomainUser, context.Context) error); ok {
		r1 = rf(userDomain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}