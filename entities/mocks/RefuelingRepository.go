// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entities "car-record/entities"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RefuelingRepository is an autogenerated mock type for the RefuelingRepository type
type RefuelingRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, record
func (_m *RefuelingRepository) Add(ctx context.Context, record entities.RefuelingRecord) error {
	ret := _m.Called(ctx, record)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.RefuelingRecord) error); ok {
		r0 = rf(ctx, record)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, recordID
func (_m *RefuelingRepository) Delete(ctx context.Context, recordID uint) error {
	ret := _m.Called(ctx, recordID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, recordID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetList provides a mock function with given fields: ctx, vehicleID
func (_m *RefuelingRepository) GetList(ctx context.Context, vehicleID uint) ([]entities.RefuelingRecord, error) {
	ret := _m.Called(ctx, vehicleID)

	var r0 []entities.RefuelingRecord
	if rf, ok := ret.Get(0).(func(context.Context, uint) []entities.RefuelingRecord); ok {
		r0 = rf(ctx, vehicleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.RefuelingRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, vehicleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, record
func (_m *RefuelingRepository) Update(ctx context.Context, record entities.RefuelingRecord) error {
	ret := _m.Called(ctx, record)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.RefuelingRecord) error); ok {
		r0 = rf(ctx, record)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRefuelingRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRefuelingRepository creates a new instance of RefuelingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRefuelingRepository(t mockConstructorTestingTNewRefuelingRepository) *RefuelingRepository {
	mock := &RefuelingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
