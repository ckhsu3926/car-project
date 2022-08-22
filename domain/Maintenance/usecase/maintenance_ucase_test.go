package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"car-record/domain/Maintenance/usecase"
	"car-record/entities"
	"car-record/entities/mocks"
)

var timeout = time.Duration(2) * time.Second
var mockRecordID, mockVehicleID uint = 6, 12
var mockRecord = entities.MaintenanceRecord{
	ID:        mockRecordID,
	VehicleID: mockVehicleID,
	Date:      "2022/09/19",
	Mileage:   decimal.NewFromFloat(3456.7),
	Shop:      "the great dealer",
	Amount:    350,
}
var mockUpdatedRecord = entities.MaintenanceRecord{
	ID:        mockRecordID,
	VehicleID: mockVehicleID,
	Date:      "2022/09/19",
	Mileage:   decimal.NewFromFloat(3456.7),
	Shop:      "the great dealer",
	Amount:    600,
}
var mockRecordDetail = entities.MaintenanceRecordDetail{
	Name:    "change oil",
	Value:   350,
	Content: "FUCHS 10W40 SILKOLENE PRO 4T XP 10w-40 *2",
}

// record
func TestCreateRecord(t *testing.T) {
	var _mockRepo = mocks.NewMaintenanceRepository(t)
	var _mockUsecase = usecase.NewMaintenanceUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("CreateRecord",
			mock.AnythingOfType("*context.emptyCtx"),          // ctx
			mock.AnythingOfType("entities.MaintenanceRecord"), // MaintenanceRecord
		).Return(nil).Once()

		err := _mockUsecase.CreateRecord(context.TODO(), mockRecord)

		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_create_failed", func(t *testing.T) {
		_mockRepo.On("CreateRecord",
			mock.AnythingOfType("*context.emptyCtx"),          // ctx
			mock.AnythingOfType("entities.MaintenanceRecord"), // MaintenanceRecord
		).Return(errors.New(`insert failed`)).Once()

		err := _mockUsecase.CreateRecord(context.TODO(), mockRecord)

		assert.EqualError(t, err, `insert failed`)
		_mockRepo.AssertExpectations(t)
	})
}

func TestGetRecordList(t *testing.T) {
	var _mockRepo = mocks.NewMaintenanceRepository(t)
	var _mockUsecase = usecase.NewMaintenanceUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("GetRecordList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // vehicleID
		).Return([]entities.MaintenanceRecord{mockRecord}, nil).Once()

		list, err := _mockUsecase.GetRecordList(context.TODO(), mockVehicleID)

		assert.NotEmpty(t, list)
		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_query_failed", func(t *testing.T) {
		_mockRepo.On("GetRecordList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // vehicleID
		).Return([]entities.MaintenanceRecord{}, errors.New(`query failed`)).Once()

		list, err := _mockUsecase.GetRecordList(context.TODO(), mockVehicleID)

		assert.Empty(t, list)
		assert.EqualError(t, err, "query failed")
		_mockRepo.AssertExpectations(t)
	})
}

func TestUpdateRecord(t *testing.T) {
	var _mockRepo = mocks.NewMaintenanceRepository(t)
	var _mockUsecase = usecase.NewMaintenanceUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("UpdateRecord",
			mock.AnythingOfType("*context.emptyCtx"),          // ctx
			mock.AnythingOfType("entities.MaintenanceRecord"), // MaintenanceRecord
		).Return(nil).Once()

		err := _mockUsecase.UpdateRecord(context.TODO(), mockUpdatedRecord)

		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_update_failed", func(t *testing.T) {
		_mockRepo.On("UpdateRecord",
			mock.AnythingOfType("*context.emptyCtx"),          // ctx
			mock.AnythingOfType("entities.MaintenanceRecord"), // MaintenanceRecord
		).Return(errors.New(`save failed`)).Once()

		err := _mockUsecase.UpdateRecord(context.TODO(), mockUpdatedRecord)

		assert.EqualError(t, err, "save failed")
		_mockRepo.AssertExpectations(t)
	})
}

func TestDeleteRecord(t *testing.T) {
	var _mockRepo = mocks.NewMaintenanceRepository(t)
	var _mockUsecase = usecase.NewMaintenanceUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("DeleteRecord",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // recordID
		).Return(nil).Once()

		err := _mockUsecase.DeleteRecord(context.TODO(), mockRecordID)

		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_delete_failed", func(t *testing.T) {
		_mockRepo.On("DeleteRecord",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // recordID
		).Return(errors.New(`delete failed`)).Once()

		err := _mockUsecase.DeleteRecord(context.TODO(), mockRecordID)

		assert.EqualError(t, err, `delete failed`)
		_mockRepo.AssertExpectations(t)
	})
}

// detail
func TestGetDetailList(t *testing.T) {
	var _mockRepo = mocks.NewMaintenanceRepository(t)
	var _mockUsecase = usecase.NewMaintenanceUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("GetDetailList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // recordID
		).Return([]entities.MaintenanceRecordDetail{mockRecordDetail}, nil).Once()

		list, err := _mockUsecase.GetDetailList(context.TODO(), mockRecordID)

		assert.NotEmpty(t, list)
		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_get_detail_list_failed", func(t *testing.T) {
		_mockRepo.On("GetDetailList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // recordID
		).Return([]entities.MaintenanceRecordDetail{}, errors.New(`query failed`)).Once()

		list, err := _mockUsecase.GetDetailList(context.TODO(), mockRecordID)

		assert.Empty(t, list)
		assert.EqualError(t, err, `query failed`)
		_mockRepo.AssertExpectations(t)
	})
}

func TestSetDetailList(t *testing.T) {
	var _mockRepo = mocks.NewMaintenanceRepository(t)
	var _mockUsecase = usecase.NewMaintenanceUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("SetDetailList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // recordID
			mock.AnythingOfType("[]entities.MaintenanceRecordDetail"),
		).Return(nil).Once()

		err := _mockUsecase.SetDetailList(context.TODO(), mockRecordID, []entities.MaintenanceRecordDetail{mockRecordDetail})

		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_set_detail_list_failed", func(t *testing.T) {
		_mockRepo.On("SetDetailList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // recordID
			mock.AnythingOfType("[]entities.MaintenanceRecordDetail"),
		).Return(errors.New(`set failed`)).Once()

		err := _mockUsecase.SetDetailList(context.TODO(), mockRecordID, []entities.MaintenanceRecordDetail{mockRecordDetail})

		assert.EqualError(t, err, `set failed`)
		_mockRepo.AssertExpectations(t)
	})
}
