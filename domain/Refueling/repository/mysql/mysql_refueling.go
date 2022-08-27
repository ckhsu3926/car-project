package mysql

import (
	"context"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"car-record/entities"
)

func (refueling) TableName() string {
	return "refueling_record"
}

type refueling struct {
	ID                uint            `gorm:"primaryKey;column:id"`
	VehicleID         uint            `gorm:"column:vehicle_id"`
	Date              string          `gorm:"column:date"`
	Station           string          `gorm:"column:station"`
	OctaneNumber      int             `gorm:"column:octane_number"`
	UnitPrice         decimal.Decimal `gorm:"column:unit_price"`
	Count             decimal.Decimal `gorm:"column:count"`
	Value             int             `gorm:"column:value"`
	Mileage           decimal.Decimal `gorm:"column:mileage"`
	MonitorFuelRecord decimal.Decimal `gorm:"column:monitor_fuel_record"`
}

type mysqlRefuelingRepository struct {
	Conn *gorm.DB
}

func NewMysqlRefuelingRepository(Conn *gorm.DB) entities.RefuelingRepository {
	return &mysqlRefuelingRepository{Conn}
}

func (r *mysqlRefuelingRepository) Add(ctx context.Context, record entities.RefuelingRecord) (err error) {
	newRecord := refueling{
		VehicleID:         record.VehicleID,
		Date:              record.Date,
		Station:           record.Station,
		OctaneNumber:      record.OctaneNumber,
		UnitPrice:         record.UnitPrice,
		Count:             record.Count,
		Value:             record.Value,
		Mileage:           record.Mileage,
		MonitorFuelRecord: record.MonitorFuelRecord,
	}
	err = r.Conn.WithContext(ctx).Create(&newRecord).Error

	return
}
func (r *mysqlRefuelingRepository) GetList(ctx context.Context, vehicleID uint) (list []entities.RefuelingRecord, err error) {
	list = []entities.RefuelingRecord{}
	result := []refueling{}

	err = r.Conn.WithContext(ctx).Model(&refueling{}).
		Where(&refueling{VehicleID: vehicleID}).
		Find(&result).Error
	if err != nil {
		return
	}

	for _, v := range result {
		list = append(list, entities.RefuelingRecord{
			ID:                v.ID,
			VehicleID:         v.VehicleID,
			Date:              v.Date,
			Station:           v.Station,
			OctaneNumber:      v.OctaneNumber,
			UnitPrice:         v.UnitPrice,
			Count:             v.Count,
			Value:             v.Value,
			Mileage:           v.Mileage,
			MonitorFuelRecord: v.MonitorFuelRecord,
		})
	}

	return
}

func (r *mysqlRefuelingRepository) Update(ctx context.Context, record entities.RefuelingRecord) (err error) {
	newRecord := refueling{
		ID:                record.ID,
		VehicleID:         record.VehicleID,
		Date:              record.Date,
		Station:           record.Station,
		OctaneNumber:      record.OctaneNumber,
		UnitPrice:         record.UnitPrice,
		Count:             record.Count,
		Value:             record.Value,
		Mileage:           record.Mileage,
		MonitorFuelRecord: record.MonitorFuelRecord,
	}
	err = r.Conn.WithContext(ctx).Save(&newRecord).Error

	return
}
func (r *mysqlRefuelingRepository) Delete(ctx context.Context, recordID uint) (err error) {
	record := refueling{ID: recordID}

	err = r.Conn.WithContext(ctx).Delete(&record).Error

	return
}
