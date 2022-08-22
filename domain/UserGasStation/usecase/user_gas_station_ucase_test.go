package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"car-record/domain/UserGasStation/usecase"
	"car-record/entities"
	"car-record/entities/mocks"
)

var timeout = time.Duration(2) * time.Second
var mockID, mockUserID uint = 2, 3
var mockName = "test gas station"
var mockList []entities.UserGasStation = []entities.UserGasStation{
	{ID: 1, UserID: mockUserID, Name: mockName + "1"},
	{ID: 2, UserID: mockUserID, Name: mockName + "2"},
	{ID: 3, UserID: mockUserID, Name: mockName + "3"},
}

func TestAdd(t *testing.T) {
	var _mockRepo = mocks.NewUserGasStationRepository(t)
	var _mockUsecase = usecase.NewUserGasStationUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("Add",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // userID
			mock.AnythingOfType("string"),            // name
		).Return(nil).Once()

		err := _mockUsecase.Add(context.TODO(), mockUserID, mockName)

		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_add_failed", func(t *testing.T) {
		_mockRepo.On("Add",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // userID
			mock.AnythingOfType("string"),            // name
		).Return(errors.New(`insert failed`)).Once()

		err := _mockUsecase.Add(context.TODO(), mockUserID, mockName)

		assert.EqualError(t, err, `insert failed`)
		_mockRepo.AssertExpectations(t)
	})
}

func TestGetList(t *testing.T) {
	var _mockRepo = mocks.NewUserGasStationRepository(t)
	var _mockUsecase = usecase.NewUserGasStationUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("GetList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // userID
		).Return(mockList, nil).Once()

		list, err := _mockUsecase.GetList(context.TODO(), mockUserID)

		assert.ObjectsAreEqual(list, mockList)
		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_get_failed", func(t *testing.T) {
		_mockRepo.On("GetList",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // userID
		).Return([]entities.UserGasStation{}, errors.New(`get failed`)).Once()

		list, err := _mockUsecase.GetList(context.TODO(), mockUserID)

		assert.Empty(t, list)
		assert.EqualError(t, err, `get failed`)
		_mockRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	var _mockRepo = mocks.NewUserGasStationRepository(t)
	var _mockUsecase = usecase.NewUserGasStationUsecase(_mockRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockRepo.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // userID
			mock.AnythingOfType("uint"),              // record id
		).Return(nil).Once()

		err := _mockUsecase.Delete(context.TODO(), mockUserID, mockID)

		assert.NoError(t, err)
		_mockRepo.AssertExpectations(t)
	})

	t.Run("failed:_delete_failed", func(t *testing.T) {
		_mockRepo.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("uint"),              // userID
			mock.AnythingOfType("uint"),              // record id
		).Return(errors.New(`delete failed`)).Once()

		err := _mockUsecase.Delete(context.TODO(), mockUserID, mockID)

		assert.EqualError(t, err, `delete failed`)
		_mockRepo.AssertExpectations(t)
	})
}
