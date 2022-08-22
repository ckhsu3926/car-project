package usecase

import (
	"context"
	"time"

	"car-record/entities"
)

type userGasStationUsecase struct {
	repo           entities.UserGasStationRepository
	contextTimeout time.Duration
}

func NewUserGasStationUsecase(u entities.UserGasStationRepository, timeout time.Duration) entities.UserGasStationUsecase {
	return &userGasStationUsecase{
		repo:           u,
		contextTimeout: timeout,
	}
}

func (uc *userGasStationUsecase) Add(ctx context.Context, userID uint, name string) (err error) {
	err = uc.repo.Add(ctx, userID, name)

	return
}
func (uc *userGasStationUsecase) GetList(ctx context.Context, userID uint) (list []entities.UserGasStation, err error) {
	list, err = uc.repo.GetList(ctx, userID)

	return
}
func (uc *userGasStationUsecase) Delete(ctx context.Context, userID uint, id uint) (err error) {
	err = uc.repo.Delete(ctx, userID, id)

	return
}
