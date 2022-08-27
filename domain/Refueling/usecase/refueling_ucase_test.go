package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"car-record/domain/Refueling/usecase"
	"car-record/entities"
	"car-record/entities/mocks"
)

var timeout = time.Duration(2) * time.Second
var mockVehicleID, mockRecordID uint = 6, 12
var mockRecord entities.RefuelingRecord = entities.RefuelingRecord{
	VehicleID:         mockVehicleID,
	Date:              "2022/08/18",
	Station:           "test station",
	OctaneNumber:      95,
	UnitPrice:         decimal.NewFromFloat(28.3),
	Count:             decimal.NewFromFloat(28.3),
	Value:             283,
	Mileage:           decimal.NewFromFloat(3456.7),
	MonitorFuelRecord: decimal.NewFromFloat(18.3),
}
var mockUpdatedRecord = entities.RefuelingRecord{
	VehicleID:         mockVehicleID,
	Date:              "2022/09/19",
	Station:           "new station",
	OctaneNumber:      95,
	UnitPrice:         decimal.NewFromFloat(28.3),
	Count:             decimal.NewFromFloat(28.3),
	Value:             283,
	Mileage:           decimal.NewFromFloat(3456.7),
	MonitorFuelRecord: decimal.NewFromFloat(18.3),
}

func TestAdd(t *testing.T) {
	var _mockRepo = mocks.NewRefuelingRepository(t)
	var _mockUsecase = usecase.NewRefuelingUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("Add",
			mock.AnythingOfType("*context.emptyCtx"),        // ctx
			mock.AnythingOfType("entities.RefuelingRecord"), // RefuelingRecord
		).Return(nil).Once()

		err := _mockUsecase.Add(context.TODO(), mockRecord)

		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_add_failed", func(t *testing.T) {
		_mockRepo.On("Add",
			mock.AnythingOfType("*context.emptyCtx"),        // ctx
			mock.AnythingOfType("entities.RefuelingRecord"), // RefuelingRecord
		).Return(errors.New(`insert failed`)).Once()

		err := _mockUsecase.Add(context.TODO(), mockRecord)

		assert.EqualError(t, err, `insert failed`)
		_mockRepo.AssertExpectations(t)
	})
}

func TestGetList(t *testing.T) {
	var _mockRepo = mocks.NewRefuelingRepository(t)
	var _mockUsecase = usecase.NewRefuelingUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("GetList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // vehicleID
		).Return([]entities.RefuelingRecord{mockRecord}, nil).Once()

		list, err := _mockUsecase.GetList(context.TODO(), mockVehicleID)

		assert.NotEmpty(t, list)
		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_query_failed", func(t *testing.T) {
		_mockRepo.On("GetList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // vehicleID
		).Return([]entities.RefuelingRecord{}, errors.New(`query failed`)).Once()

		list, err := _mockUsecase.GetList(context.TODO(), mockVehicleID)

		assert.Empty(t, list)
		assert.EqualError(t, err, "query failed")
		_mockRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	var _mockRepo = mocks.NewRefuelingRepository(t)
	var _mockUsecase = usecase.NewRefuelingUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("Update",
			mock.AnythingOfType("*context.emptyCtx"),        // ctx
			mock.AnythingOfType("entities.RefuelingRecord"), // RefuelingRecord
		).Return(nil).Once()

		err := _mockUsecase.Update(context.TODO(), mockUpdatedRecord)

		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_update_failed", func(t *testing.T) {
		_mockRepo.On("Update",
			mock.AnythingOfType("*context.emptyCtx"),        // ctx
			mock.AnythingOfType("entities.RefuelingRecord"), // RefuelingRecord
		).Return(errors.New(`save failed`)).Once()

		err := _mockUsecase.Update(context.TODO(), mockUpdatedRecord)

		assert.EqualError(t, err, "save failed")
		_mockRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	var _mockRepo = mocks.NewRefuelingRepository(t)
	var _mockUsecase = usecase.NewRefuelingUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // recordID
		).Return(nil).Once()

		err := _mockUsecase.Delete(context.TODO(), mockRecordID)

		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_delete_failed", func(t *testing.T) {
		_mockRepo.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // recordID
		).Return(errors.New(`delete failed`)).Once()

		err := _mockUsecase.Delete(context.TODO(), mockRecordID)

		assert.EqualError(t, err, `delete failed`)
		_mockRepo.AssertExpectations(t)
	})
}
