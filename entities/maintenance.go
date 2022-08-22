package entities

import (
	"context"

	"github.com/shopspring/decimal"
)

type MaintenanceRecord struct {
	ID        uint            `json:"id"`
	VehicleID uint            `json:"vehicleID"`
	Date      string          `json:"date"`
	Mileage   decimal.Decimal `json:"mileage"`
	Shop      string          `json:"shop"`
	Amount    int             `json:"amount"`
}
type MaintenanceRecordDetail struct {
	Name    string `json:"name"`
	Value   int    `json:"value"`
	Content string `json:"content"`
}

type MaintenanceUsecase interface {
	CreateRecord(ctx context.Context, record MaintenanceRecord) (err error)
	GetRecordList(ctx context.Context, vehicleID uint) (list []MaintenanceRecord, err error)
	UpdateRecord(ctx context.Context, record MaintenanceRecord) (err error)
	DeleteRecord(ctx context.Context, recordID uint) (err error)

	GetDetailList(ctx context.Context, recordID uint) (list []MaintenanceRecordDetail, err error)
	SetDetailList(ctx context.Context, recordID uint, detailList []MaintenanceRecordDetail) (err error)
}

type MaintenanceRepository interface {
	CreateRecord(ctx context.Context, record MaintenanceRecord) (err error)
	GetRecordList(ctx context.Context, vehicleID uint) (list []MaintenanceRecord, err error)
	UpdateRecord(ctx context.Context, record MaintenanceRecord) (err error)
	DeleteRecord(ctx context.Context, recordID uint) (err error)

	SetDetailList(ctx context.Context, recordID uint, detailList []MaintenanceRecordDetail) (err error)
	GetDetailList(ctx context.Context, recordID uint) (list []MaintenanceRecordDetail, err error)
}
