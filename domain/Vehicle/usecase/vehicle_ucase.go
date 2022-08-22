package usecase

import (
	"context"
	"errors"
	"time"

	"car-record/entities"
)

type vehicleUsecase struct {
	vehicleRepo    entities.VehicleRepository
	contextTimeout time.Duration
}

func NewVehicleUsecase(v entities.VehicleRepository, timeout time.Duration) entities.VehicleUsecase {
	return &vehicleUsecase{
		vehicleRepo:    v,
		contextTimeout: timeout,
	}
}

func (uc *vehicleUsecase) Add(ctx context.Context, userID uint, name, license, company, model string) (err error) {
	licenseExist := uc.vehicleRepo.IsLicenseExist(ctx, userID, license)
	if licenseExist {
		return errors.New(`license already exist`)
	}
	err = uc.vehicleRepo.Add(ctx, userID, name, license, company, model)

	return
}

func (uc *vehicleUsecase) GetList(ctx context.Context, userID uint) (list []entities.Vehicle, err error) {
	list, err = uc.vehicleRepo.GetList(ctx, userID)

	return
}

func (uc *vehicleUsecase) Get(ctx context.Context, id uint) (vehicle entities.VehicleDetail, err error) {
	vehicle, err = uc.vehicleRepo.Get(ctx, id)

	return
}

func (uc *vehicleUsecase) Edit(ctx context.Context, userID uint, vehicle entities.VehicleDetail) (err error) {
	err = uc.vehicleRepo.Edit(ctx, userID, vehicle)

	return
}

func (uc *vehicleUsecase) Delete(ctx context.Context, id uint) (err error) {
	err = uc.vehicleRepo.Delete(ctx, id)

	return
}
