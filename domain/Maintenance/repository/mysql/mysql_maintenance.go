package mysql

import (
	"context"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"car-record/entities"
)

func (maintenanceRecord) TableName() string {
	return "maintenance_record"
}

type maintenanceRecord struct {
	ID        uint            `gorm:"primaryKey;column:id"`
	VehicleID uint            `gorm:"column:vehicleID"`
	Date      string          `gorm:"column:date"`
	Mileage   decimal.Decimal `gorm:"column:mileage"`
	Shop      string          `gorm:"column:shop"`
	Amount    int             `gorm:"column:amount"`
}

type mysqlMaintenanceRepository struct {
	Conn *gorm.DB
}

func NewMysqlMaintenanceRepository(Conn *gorm.DB) entities.MaintenanceRepository {
	return &mysqlMaintenanceRepository{Conn}
}

func (r *mysqlMaintenanceRepository) CreateRecord(ctx context.Context, record entities.MaintenanceRecord) (err error) {
	newRecord := maintenanceRecord{
		VehicleID: record.VehicleID,
		Date:      record.Date,
		Mileage:   record.Mileage,
		Shop:      record.Shop,
		Amount:    record.Amount,
	}
	err = r.Conn.WithContext(ctx).Create(&newRecord).Error

	return
}

func (r *mysqlMaintenanceRepository) GetRecordList(ctx context.Context, vehicleID uint) (list []entities.MaintenanceRecord, err error) {
	list = []entities.MaintenanceRecord{}
	result := []maintenanceRecord{}

	err = r.Conn.WithContext(ctx).Model(&maintenanceRecord{}).
		Where(&maintenanceRecord{VehicleID: vehicleID}).
		Find(&result).Error
	if err != nil {
		return
	}

	for _, v := range result {
		list = append(list, entities.MaintenanceRecord{
			ID:        v.ID,
			VehicleID: v.VehicleID,
			Date:      v.Date,
			Mileage:   v.Mileage,
			Shop:      v.Shop,
			Amount:    v.Amount,
		})
	}

	return
}

func (r *mysqlMaintenanceRepository) UpdateRecord(ctx context.Context, record entities.MaintenanceRecord) (err error) {
	newRecord := maintenanceRecord{
		ID:        record.ID,
		VehicleID: record.VehicleID,
		Date:      record.Date,
		Mileage:   record.Mileage,
		Shop:      record.Shop,
		Amount:    record.Amount,
	}
	err = r.Conn.WithContext(ctx).Save(&newRecord).Error

	return
}

func (r *mysqlMaintenanceRepository) DeleteRecord(ctx context.Context, recordID uint) (err error) {
	record := maintenanceRecord{ID: recordID}

	err = r.Conn.WithContext(ctx).Delete(&record).Error

	return
}
