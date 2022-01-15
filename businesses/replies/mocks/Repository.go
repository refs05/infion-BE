// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	replies "infion-BE/businesses/replies"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CountByCommentID provides a mock function with given fields: ctx, id
func (_m *Repository) CountByCommentID(ctx context.Context, id int) (int, error) {
	ret := _m.Called(ctx, id)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, int) int); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, repliesDomain
func (_m *Repository) Delete(ctx context.Context, repliesDomain *replies.Domain) (replies.Domain, error) {
	ret := _m.Called(ctx, repliesDomain)

	var r0 replies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *replies.Domain) replies.Domain); ok {
		r0 = rf(ctx, repliesDomain)
	} else {
		r0 = ret.Get(0).(replies.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *replies.Domain) error); ok {
		r1 = rf(ctx, repliesDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, repliesId
func (_m *Repository) GetByID(ctx context.Context, repliesId int) (replies.Domain, error) {
	ret := _m.Called(ctx, repliesId)

	var r0 replies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) replies.Domain); ok {
		r0 = rf(ctx, repliesId)
	} else {
		r0 = ret.Get(0).(replies.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, repliesId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplies provides a mock function with given fields: ctx
func (_m *Repository) GetReplies(ctx context.Context) ([]replies.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []replies.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []replies.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]replies.Domain)
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

// GetRepliesByCommentID provides a mock function with given fields: ctx, commentId
func (_m *Repository) GetRepliesByCommentID(ctx context.Context, commentId int) ([]replies.Domain, error) {
	ret := _m.Called(ctx, commentId)

	var r0 []replies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) []replies.Domain); ok {
		r0 = rf(ctx, commentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]replies.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, commentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, repliesDomain
func (_m *Repository) Store(ctx context.Context, repliesDomain *replies.Domain) (replies.Domain, error) {
	ret := _m.Called(ctx, repliesDomain)

	var r0 replies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *replies.Domain) replies.Domain); ok {
		r0 = rf(ctx, repliesDomain)
	} else {
		r0 = ret.Get(0).(replies.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *replies.Domain) error); ok {
		r1 = rf(ctx, repliesDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, repliesDomain
func (_m *Repository) Update(ctx context.Context, repliesDomain *replies.Domain) (replies.Domain, error) {
	ret := _m.Called(ctx, repliesDomain)

	var r0 replies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *replies.Domain) replies.Domain); ok {
		r0 = rf(ctx, repliesDomain)
	} else {
		r0 = ret.Get(0).(replies.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *replies.Domain) error); ok {
		r1 = rf(ctx, repliesDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
