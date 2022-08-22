package entities

import (
	"context"

	"github.com/shopspring/decimal"
)

type RefuelingRecord struct {
	ID                uint            `json:"id"`
	VehicleID         uint            `json:"vehicleID"`
	Date              string          `json:"date"`
	Station           string          `json:"station"`
	OctaneNumber      int             `json:"octaneNumber"`
	UnitPrice         decimal.Decimal `json:"unitPrice"`
	Count             decimal.Decimal `json:"count"`
	Value             int             `json:"value"`
	Mileage           decimal.Decimal `json:"mileage"`
	MointorFuelRecord decimal.Decimal `json:"mointorFuelRecord"`
}

type RefuelingUsecase interface {
	Add(ctx context.Context, record RefuelingRecord) (err error)
	GetList(ctx context.Context, vehicleID uint) (list []RefuelingRecord, err error)
	Update(ctx context.Context, record RefuelingRecord) (err error)
	Delete(ctx context.Context, recordID uint) (err error)
}

type RefuelingRepository interface {
	Add(ctx context.Context, record RefuelingRecord) (err error)
	GetList(ctx context.Context, vehicleID uint) (list []RefuelingRecord, err error)
	Update(ctx context.Context, record RefuelingRecord) (err error)
	Delete(ctx context.Context, recordID uint) (err error)
}
