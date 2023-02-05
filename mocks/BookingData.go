// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	booking "campyuk-api/features/booking"

	mock "github.com/stretchr/testify/mock"
)

// BookingData is an autogenerated mock type for the BookingData type
type BookingData struct {
	mock.Mock
}

// Cancel provides a mock function with given fields: userID, bookingID
func (_m *BookingData) Cancel(userID uint, bookingID uint) (booking.Core, error) {
	ret := _m.Called(userID, bookingID)

	var r0 booking.Core
	if rf, ok := ret.Get(0).(func(uint, uint) booking.Core); ok {
		r0 = rf(userID, bookingID)
	} else {
		r0 = ret.Get(0).(booking.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, bookingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: userID, virtualNumber
func (_m *BookingData) Create(userID uint, virtualNumber string) (booking.Core, error) {
	ret := _m.Called(userID, virtualNumber)

	var r0 booking.Core
	if rf, ok := ret.Get(0).(func(uint, string) booking.Core); ok {
		r0 = rf(userID, virtualNumber)
	} else {
		r0 = ret.Get(0).(booking.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(userID, virtualNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: userID, bookingID
func (_m *BookingData) GetByID(userID uint, bookingID uint) (booking.Core, error) {
	ret := _m.Called(userID, bookingID)

	var r0 booking.Core
	if rf, ok := ret.Get(0).(func(uint, uint) booking.Core); ok {
		r0 = rf(userID, bookingID)
	} else {
		r0 = ret.Get(0).(booking.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, bookingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: userID
func (_m *BookingData) List(userID uint) ([]booking.Core, error) {
	ret := _m.Called(userID)

	var r0 []booking.Core
	if rf, ok := ret.Get(0).(func(uint) []booking.Core); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: bookingID
func (_m *BookingData) Update(bookingID uint) (booking.Core, error) {
	ret := _m.Called(bookingID)

	var r0 booking.Core
	if rf, ok := ret.Get(0).(func(uint) booking.Core); ok {
		r0 = rf(bookingID)
	} else {
		r0 = ret.Get(0).(booking.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(bookingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookingData interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookingData creates a new instance of BookingData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookingData(t mockConstructorTestingTNewBookingData) *BookingData {
	mock := &BookingData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
