package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"car-record/domain/User/usecase"
	"car-record/entities"
	"car-record/entities/mocks"
)

var timeout = time.Duration(2) * time.Second
var mockUsername, mockPassword, mockToken = "user", "password", "token"

func TestRegister(t *testing.T) {
	var _mockUserRepo = mocks.NewUserRepository(t)
	var _mockUserUsecase = usecase.NewUserUsecase(_mockUserRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockUserRepo.On("IsUsernameExist",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
		).Return(false).Once()
		_mockUserRepo.On("Create",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
			mock.AnythingOfType("string"),            // password
			mock.AnythingOfType("string"),            // token
		).Return(nil).Once()

		token, err := _mockUserUsecase.Register(context.TODO(), mockUsername, mockPassword)

		assert.NotEmpty(t, token)
		assert.NoError(t, err)
		_mockUserRepo.AssertExpectations(t)
	})

	t.Run("failed: user_aleady_exist", func(t *testing.T) {
		_mockUserRepo.On("IsUsernameExist",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
		).Return(true).Once()

		token, err := _mockUserUsecase.Register(context.TODO(), mockUsername, mockPassword)

		assert.Empty(t, token)
		assert.EqualError(t, err, `username already exist`)
		_mockUserRepo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	var _mockUserRepo = mocks.NewUserRepository(t)
	var _mockUserUsecase = usecase.NewUserUsecase(_mockUserRepo, timeout)
	mockUsername, mockPassword := "user", "password"

	t.Run("success", func(t *testing.T) {
		_mockUserRepo.On("Get",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
			mock.AnythingOfType("string"),            // password
		).Return(entities.User{Username: mockUsername}, nil).Once()
		_mockUserRepo.On("UpdateToken",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
			mock.AnythingOfType("string"),            // token
		).Return(nil).Once()

		token, err := _mockUserUsecase.Login(context.TODO(), mockUsername, mockPassword)

		assert.NotEmpty(t, token)
		assert.NoError(t, err)
		_mockUserRepo.AssertExpectations(t)
	})

	t.Run("failed: user not exist", func(t *testing.T) {
		_mockUserRepo.On("Get",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
			mock.AnythingOfType("string"),            // password
		).Return(entities.User{}, errors.New(`record not found`)).Once()

		token, err := _mockUserUsecase.Login(context.TODO(), mockUsername, mockPassword)

		assert.Empty(t, token)
		assert.EqualError(t, err, `record not found`)
		_mockUserRepo.AssertExpectations(t)
	})
	t.Run("failed: update token failed", func(t *testing.T) {
		_mockUserRepo.On("Get",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
			mock.AnythingOfType("string"),            // password
		).Return(entities.User{Username: mockUsername}, nil).Once()
		_mockUserRepo.On("UpdateToken",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
			mock.AnythingOfType("string"),            // token
		).Return(errors.New(`update failed`)).Once()

		token, err := _mockUserUsecase.Login(context.TODO(), mockUsername, mockPassword)

		assert.Empty(t, token)
		assert.EqualError(t, err, `update failed`)
		_mockUserRepo.AssertExpectations(t)
	})
}

func TestLogout(t *testing.T) {
	var _mockUserRepo = mocks.NewUserRepository(t)
	var _mockUserUsecase = usecase.NewUserUsecase(_mockUserRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockUserRepo.On("GetByToken",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // token
		).Return(entities.User{LoginToken: mockToken}, nil).Once()
		_mockUserRepo.On("UpdateToken",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
			mock.AnythingOfType("string"),            // token
		).Return(nil).Once()

		err := _mockUserUsecase.Logout(context.TODO(), mockToken)

		assert.NoError(t, err)
		_mockUserRepo.AssertExpectations(t)
	})

	t.Run("failed: token not exist", func(t *testing.T) {
		_mockUserRepo.On("GetByToken",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // token
		).Return(entities.User{}, errors.New(`record not found`)).Once()

		err := _mockUserUsecase.Logout(context.TODO(), mockToken)

		assert.EqualError(t, err, `record not found`)
		_mockUserRepo.AssertExpectations(t)
	})
	t.Run("failed: update token failed", func(t *testing.T) {
		_mockUserRepo.On("GetByToken",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // token
		).Return(entities.User{Username: mockUsername}, nil).Once()
		_mockUserRepo.On("UpdateToken",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // username
			mock.AnythingOfType("string"),            // token
		).Return(errors.New(`update failed`)).Once()

		err := _mockUserUsecase.Logout(context.TODO(), mockToken)

		assert.EqualError(t, err, `update failed`)
		_mockUserRepo.AssertExpectations(t)
	})
}

func TestAuthorize(t *testing.T) {
	var _mockUserRepo = mocks.NewUserRepository(t)
	var _mockUserUsecase = usecase.NewUserUsecase(_mockUserRepo, timeout)

	t.Run("success", func(t *testing.T) {
		_mockUserRepo.On("GetByToken",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // token
		).Return(entities.User{Username: mockUsername, LoginToken: mockToken}, nil).Once()

		user, err := _mockUserUsecase.Authorize(context.TODO(), mockToken)

		assert.Equal(t, user.Username, mockUsername)
		assert.NoError(t, err)
		_mockUserRepo.AssertExpectations(t)
	})

	t.Run("failed: token not exist", func(t *testing.T) {
		_mockUserRepo.On("GetByToken",
			mock.AnythingOfType("*context.emptyCtx"), // ctx
			mock.AnythingOfType("string"),            // token
		).Return(entities.User{}, errors.New(`record not found`)).Once()

		err := _mockUserUsecase.Logout(context.TODO(), mockToken)

		assert.EqualError(t, err, `record not found`)
		_mockUserRepo.AssertExpectations(t)
	})
}
