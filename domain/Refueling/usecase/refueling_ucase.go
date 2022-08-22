package usecase

import (
	"context"
	"time"

	"car-record/entities"
)

type refuelingUsecase struct {
	repo           entities.RefuelingRepository
	contextTimeout time.Duration
}

func NewRefuelingUsecase(r entities.RefuelingRepository, timeout time.Duration) entities.RefuelingUsecase {
	return &refuelingUsecase{
		repo:           r,
		contextTimeout: timeout,
	}
}

func (uc *refuelingUsecase) Add(ctx context.Context, record entities.RefuelingRecord) (err error) {
	err = uc.repo.Add(ctx, record)

	return
}
func (uc *refuelingUsecase) GetList(ctx context.Context, vehicleID uint) (list []entities.RefuelingRecord, err error) {
	list, err = uc.repo.GetList(ctx, vehicleID)

	return
}

func (uc *refuelingUsecase) Update(ctx context.Context, record entities.RefuelingRecord) (err error) {
	err = uc.repo.Update(ctx, record)

	return
}
func (uc *refuelingUsecase) Delete(ctx context.Context, recordID uint) (err error) {
	err = uc.repo.Delete(ctx, recordID)

	return
}
