// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	followUsers "infion-BE/businesses/followUsers"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CountByFollowedID provides a mock function with given fields: ctx, id
func (_m *Repository) CountByFollowedID(ctx context.Context, id uint) (int, error) {
	ret := _m.Called(ctx, id)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, uint) int); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, followUsersDomain
func (_m *Repository) Delete(ctx context.Context, followUsersDomain *followUsers.Domain) (followUsers.Domain, error) {
	ret := _m.Called(ctx, followUsersDomain)

	var r0 followUsers.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *followUsers.Domain) followUsers.Domain); ok {
		r0 = rf(ctx, followUsersDomain)
	} else {
		r0 = ret.Get(0).(followUsers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *followUsers.Domain) error); ok {
		r1 = rf(ctx, followUsersDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, followUsersId
func (_m *Repository) GetByID(ctx context.Context, followUsersId int) (followUsers.Domain, error) {
	ret := _m.Called(ctx, followUsersId)

	var r0 followUsers.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) followUsers.Domain); ok {
		r0 = rf(ctx, followUsersId)
	} else {
		r0 = ret.Get(0).(followUsers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, followUsersId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDuplicate provides a mock function with given fields: ctx, followedID, followerID
func (_m *Repository) GetDuplicate(ctx context.Context, followedID int, followerID int) (followUsers.Domain, error) {
	ret := _m.Called(ctx, followedID, followerID)

	var r0 followUsers.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, int) followUsers.Domain); ok {
		r0 = rf(ctx, followedID, followerID)
	} else {
		r0 = ret.Get(0).(followUsers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, followedID, followerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, followUsersDomain
func (_m *Repository) Store(ctx context.Context, followUsersDomain *followUsers.Domain) (followUsers.Domain, error) {
	ret := _m.Called(ctx, followUsersDomain)

	var r0 followUsers.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *followUsers.Domain) followUsers.Domain); ok {
		r0 = rf(ctx, followUsersDomain)
	} else {
		r0 = ret.Get(0).(followUsers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *followUsers.Domain) error); ok {
		r1 = rf(ctx, followUsersDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, followUsersDomain
func (_m *Repository) Update(ctx context.Context, followUsersDomain *followUsers.Domain) (followUsers.Domain, error) {
	ret := _m.Called(ctx, followUsersDomain)

	var r0 followUsers.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *followUsers.Domain) followUsers.Domain); ok {
		r0 = rf(ctx, followUsersDomain)
	} else {
		r0 = ret.Get(0).(followUsers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *followUsers.Domain) error); ok {
		r1 = rf(ctx, followUsersDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
