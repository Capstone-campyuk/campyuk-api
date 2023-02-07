// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	image "campyuk-api/features/image"

	mock "github.com/stretchr/testify/mock"
)

// ImageData is an autogenerated mock type for the ImageData type
type ImageData struct {
	mock.Mock
}

// Add provides a mock function with given fields: userID, core
func (_m *ImageData) Add(userID uint, core image.Core) error {
	ret := _m.Called(userID, core)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, image.Core) error); ok {
		r0 = rf(userID, core)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: usesrID, imageID
func (_m *ImageData) Delete(usesrID uint, imageID uint) error {
	ret := _m.Called(usesrID, imageID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(usesrID, imageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: usesrID, imageID, core
func (_m *ImageData) Update(usesrID uint, imageID uint, core image.Core) error {
	ret := _m.Called(usesrID, imageID, core)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, image.Core) error); ok {
		r0 = rf(usesrID, imageID, core)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewImageData interface {
	mock.TestingT
	Cleanup(func())
}

// NewImageData creates a new instance of ImageData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImageData(t mockConstructorTestingTNewImageData) *ImageData {
	mock := &ImageData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
