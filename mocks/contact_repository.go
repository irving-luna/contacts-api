// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	model "accelone-api/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ContactRepository is an autogenerated mock type for the ContactRepository type
type ContactRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, contact
func (_m *ContactRepository) Create(ctx context.Context, contact model.Contact) (*model.Contact, error) {
	ret := _m.Called(ctx, contact)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Contact) (*model.Contact, error)); ok {
		return rf(ctx, contact)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Contact) *model.Contact); ok {
		r0 = rf(ctx, contact)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Contact) error); ok {
		r1 = rf(ctx, contact)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, contactID
func (_m *ContactRepository) Delete(ctx context.Context, contactID string) error {
	ret := _m.Called(ctx, contactID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, contactID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, contactID
func (_m *ContactRepository) Get(ctx context.Context, contactID string) (*model.Contact, error) {
	ret := _m.Called(ctx, contactID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Contact, error)); ok {
		return rf(ctx, contactID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Contact); ok {
		r0 = rf(ctx, contactID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, contactID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, contactID, data
func (_m *ContactRepository) Update(ctx context.Context, contactID string, data model.Contact) (*model.Contact, error) {
	ret := _m.Called(ctx, contactID, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *model.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Contact) (*model.Contact, error)); ok {
		return rf(ctx, contactID, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Contact) *model.Contact); ok {
		r0 = rf(ctx, contactID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, model.Contact) error); ok {
		r1 = rf(ctx, contactID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewContactRepository creates a new instance of ContactRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContactRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContactRepository {
	mock := &ContactRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
