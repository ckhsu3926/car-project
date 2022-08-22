package usecase

import (
	"context"
	"time"

	"car-record/entities"
)

type maintenanceUsecase struct {
	repo           entities.MaintenanceRepository
	contextTimeout time.Duration
}

func NewMaintenanceUsecase(r entities.MaintenanceRepository, timeout time.Duration) entities.MaintenanceUsecase {
	return &maintenanceUsecase{
		repo:           r,
		contextTimeout: timeout,
	}
}

func (uc *maintenanceUsecase) CreateRecord(ctx context.Context, record entities.MaintenanceRecord) (err error) {
	err = uc.repo.CreateRecord(ctx, record)

	return
}

func (uc *maintenanceUsecase) GetRecordList(ctx context.Context, vehicleID uint) (list []entities.MaintenanceRecord, err error) {
	list, err = uc.repo.GetRecordList(ctx, vehicleID)

	return
}
func (uc *maintenanceUsecase) UpdateRecord(ctx context.Context, record entities.MaintenanceRecord) (err error) {
	err = uc.repo.UpdateRecord(ctx, record)

	return
}
func (uc *maintenanceUsecase) DeleteRecord(ctx context.Context, recordID uint) (err error) {
	err = uc.repo.DeleteRecord(ctx, recordID)

	return
}
func (uc *maintenanceUsecase) GetDetailList(ctx context.Context, recordID uint) (list []entities.MaintenanceRecordDetail, err error) {
	list, err = uc.repo.GetDetailList(ctx, recordID)

	return
}
func (uc *maintenanceUsecase) SetDetailList(ctx context.Context, recordID uint, detailList []entities.MaintenanceRecordDetail) (err error) {
	err = uc.repo.SetDetailList(ctx, recordID, detailList)

	return
}
