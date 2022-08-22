package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"car-record/domain/Vehicle/usecase"
	"car-record/entities"
	"car-record/entities/mocks"
)

var timeout = time.Duration(2) * time.Second
var mockUserID, mockVehicleID uint = 4, 6
var mockName, mockLicense, mockCompany, mockModel = "mockName", "mock License", "mock Company", "mock Model"
var mockVehicle = entities.VehicleDetail{
	ID:                 6,
	Name:               mockName,
	License:            mockLicense,
	Company:            mockCompany,
	Model:              mockModel,
	EngineDisplacement: decimal.NewFromFloat(148.8),
	PurchaseDate:       "2022/02/18",
}
var queryFailedReason = "record not found"

func TestAdd(t *testing.T) {
	var _mockVehicleRepo = mocks.NewVehicleRepository(t)
	var _mockVehicleUsecase = usecase.NewVehicleUsecase(_mockVehicleRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockVehicleRepo.On("IsLicenseExist",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // user id
			mock.AnythingOfType("string"),            // license
		).Return(false).Once()
		_mockVehicleRepo.On("Add",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // user id
			mock.AnythingOfType("string"),            // name
			mock.AnythingOfType("string"),            // license
			mock.AnythingOfType("string"),            // company
			mock.AnythingOfType("string"),            // model
		).Return(nil).Once()

		err := _mockVehicleUsecase.Add(context.TODO(), mockUserID, mockName, mockLicense, mockCompany, mockModel)

		assert.NoError(t, err)
		_mockVehicleRepo.AssertExpectations(t)
	})

	t.Run("failed:_license_aleady_exist", func(t *testing.T) {
		_mockVehicleRepo.On("IsLicenseExist",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // user id
			mock.AnythingOfType("string"),            // license
		).Return(true).Once()

		err := _mockVehicleUsecase.Add(context.TODO(), mockUserID, mockName, mockLicense, mockCompany, mockModel)

		assert.EqualError(t, err, `license already exist`)
		_mockVehicleRepo.AssertExpectations(t)
	})
	t.Run("failed:_add_failed", func(t *testing.T) {
		_mockVehicleRepo.On("IsLicenseExist",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // user id
			mock.AnythingOfType("string"),            // license
		).Return(false).Once()
		_mockVehicleRepo.On("Add",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // user id
			mock.AnythingOfType("string"),            // name
			mock.AnythingOfType("string"),            // license
			mock.AnythingOfType("string"),            // company
			mock.AnythingOfType("string"),            // model
		).Return(errors.New(`insert failed`)).Once()

		err := _mockVehicleUsecase.Add(context.TODO(), mockUserID, mockName, mockLicense, mockCompany, mockModel)

		assert.EqualError(t, err, `insert failed`)
		_mockVehicleRepo.AssertExpectations(t)
	})
}

func TestGetList(t *testing.T) {
	var _mockVehicleRepo = mocks.NewVehicleRepository(t)
	var _mockVehicleUsecase = usecase.NewVehicleUsecase(_mockVehicleRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockVehicleRepo.On("GetList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // user id
		).Return([]entities.Vehicle{{ID: 1}, {ID: 2}}, nil).Once()

		list, err := _mockVehicleUsecase.GetList(context.TODO(), mockUserID)

		assert.NotEmpty(t, list)
		assert.NoError(t, err)
		_mockVehicleRepo.AssertExpectations(t)
	})

	t.Run("failed:_query_failed", func(t *testing.T) {
		_mockVehicleRepo.On("GetList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // user id
		).Return([]entities.Vehicle{}, errors.New(queryFailedReason)).Once()

		list, err := _mockVehicleUsecase.GetList(context.TODO(), mockUserID)

		assert.Empty(t, list)
		assert.EqualError(t, err, queryFailedReason)
		_mockVehicleRepo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	var _mockVehicleRepo = mocks.NewVehicleRepository(t)
	var _mockVehicleUsecase = usecase.NewVehicleUsecase(_mockVehicleRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockVehicleRepo.On("Get",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // vehicle id
		).Return(entities.VehicleDetail{ID: 2}, nil).Once()

		vehicle, err := _mockVehicleUsecase.Get(context.TODO(), mockVehicleID)

		assert.NotEqual(t, vehicle.ID, 2)
		assert.NoError(t, err)
		_mockVehicleRepo.AssertExpectations(t)
	})

	t.Run("failed:_query_failed", func(t *testing.T) {
		_mockVehicleRepo.On("Get",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // vehicle id
		).Return(entities.VehicleDetail{}, errors.New(queryFailedReason)).Once()

		_, err := _mockVehicleUsecase.Get(context.TODO(), mockVehicleID)

		assert.EqualError(t, err, queryFailedReason)
		_mockVehicleRepo.AssertExpectations(t)
	})
}

func TestEdit(t *testing.T) {
	var _mockVehicleRepo = mocks.NewVehicleRepository(t)
	var _mockVehicleUsecase = usecase.NewVehicleUsecase(_mockVehicleRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockVehicleRepo.On("Edit",
			mock.AnythingOfType("*context.emptyCtx"),      // ctx
			mock.AnythingOfType("uint"),                   // user id
			mock.AnythingOfType("entities.VehicleDetail"), // vehicle detail
		).Return(nil).Once()

		err := _mockVehicleUsecase.Edit(context.TODO(), mockUserID, mockVehicle)

		assert.NoError(t, err)
		_mockVehicleRepo.AssertExpectations(t)
	})

	t.Run("failed:_save_failed", func(t *testing.T) {
		_mockVehicleRepo.On("Edit",
			mock.AnythingOfType("*context.emptyCtx"),      // ctx
			mock.AnythingOfType("uint"),                   // user id
			mock.AnythingOfType("entities.VehicleDetail"), // vehicle detail
		).Return(errors.New(queryFailedReason)).Once()

		err := _mockVehicleUsecase.Edit(context.TODO(), mockUserID, mockVehicle)

		assert.EqualError(t, err, queryFailedReason)
		_mockVehicleRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	var _mockVehicleRepo = mocks.NewVehicleRepository(t)
	var _mockVehicleUsecase = usecase.NewVehicleUsecase(_mockVehicleRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockVehicleRepo.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // vehicle id
		).Return(nil).Once()

		err := _mockVehicleUsecase.Delete(context.TODO(), mockVehicleID)

		assert.NoError(t, err)
		_mockVehicleRepo.AssertExpectations(t)
	})

	t.Run("failed:_delete_failed", func(t *testing.T) {
		_mockVehicleRepo.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // vehicle id
		).Return(errors.New(`delete failed`)).Once()

		err := _mockVehicleUsecase.Delete(context.TODO(), mockVehicleID)

		assert.EqualError(t, err, `delete failed`)
		_mockVehicleRepo.AssertExpectations(t)
	})
}
