package usecase

import (
	"context"
	"errors"
	"strconv"
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

func (uc *vehicleUsecase) Get(ctx context.Context, userID uint, id string) (vehicle entities.VehicleDetail, err error) {
	uintID, parseUintErr := strconv.ParseUint(id, 10, 32)
	if parseUintErr != nil {
		err = parseUintErr
		return
	}

	vehicle, err = uc.vehicleRepo.Get(ctx, userID, uint(uintID))

	return
}

func (uc *vehicleUsecase) Edit(ctx context.Context, userID uint, vehicle entities.VehicleDetail) (err error) {
	err = uc.vehicleRepo.Edit(ctx, userID, vehicle)

	return
}
