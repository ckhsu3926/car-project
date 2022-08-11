package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"car-record/entities"
)

type userUsecase struct {
	userRepo       entities.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u entities.UserRepository, timeout time.Duration) entities.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (uc *userUsecase) Register(ctx context.Context, username string, password string) (token string, err error) {
	usernameExist := uc.userRepo.IsUsernameExist(ctx, username)
	if usernameExist {
		return "", errors.New(`username already exist`)
	}
	token = createToken(username, password)
	err = uc.userRepo.Create(ctx, username, password, token)

	return
}

func (uc *userUsecase) Login(ctx context.Context, username string, password string) (token string, err error) {
	user, err := uc.userRepo.Get(ctx, username, password)
	if err != nil {
		return
	}

	token = createToken(username, password)
	err = uc.userRepo.UpdateToken(ctx, user.Username, token)

	return
}

func (uc *userUsecase) Logout(ctx context.Context, token string) (err error) {
	user, err := uc.userRepo.GetByToken(ctx, token)
	if err != nil {
		return
	}

	err = uc.userRepo.UpdateToken(ctx, user.Username, "")

	return
}

func (uc *userUsecase) Authorize(ctx context.Context, token string) (user entities.User, err error) {
	user, err = uc.userRepo.GetByToken(ctx, token)

	return
}

func createToken(username, password string) string {
	t := time.Now()
	hash := md5.Sum([]byte(username + password + t.Format("2006-01-02 15:04:05")))
	return hex.EncodeToString(hash[:])
}
