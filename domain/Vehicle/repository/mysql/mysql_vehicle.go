package mysql

import (
	"context"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"car-record/entities"
)

func (vehicle) TableName() string {
	return "vehicle"
}

type vehicle struct {
	ID                  uint            `gorm:"primaryKey;column:id"`
	UserID              uint            `gorm:"not null;column:user_id"`
	Name                string          `gorm:"not null;column:name"`
	License             string          `gorm:"not null;column:license"`
	Company             string          `gorm:"column:company"`
	Model               string          `gorm:"column:model"`
	EngineDisplacement  decimal.Decimal `gorm:"column:engine_displacement"`
	EngineNumber        string          `gorm:"column:engine_number"`
	DefaultOctaneNumber int             `gorm:"column:default_octane_number"`
	Purchase            int             `gorm:"column:purchase"`
	PurchaseDate        string          `gorm:"column:purchase_date"`
	PurchaseLocation    string          `gorm:"column:purchase_location"`
	PurchaseMileage     decimal.Decimal `gorm:"column:purchase_mileage"`
	Sold                int             `gorm:"column:sold"`
	SoldDate            string          `gorm:"column:sold_date"`
	SoldMileage         decimal.Decimal `gorm:"column:sold_mileage"`
	MileageReset        decimal.Decimal `gorm:"column:mileage_reset"`
}

type mysqlVehicleRepository struct {
	Conn *gorm.DB
}

func NewMysqlVehicleRepository(Conn *gorm.DB) entities.VehicleRepository {
	return &mysqlVehicleRepository{Conn}
}

func (r *mysqlVehicleRepository) Add(ctx context.Context, userID uint, name, license, company, model string) (err error) {
	newVehicle := vehicle{UserID: userID, Name: name, License: license, Company: company, Model: model}
	err = r.Conn.WithContext(ctx).Create(&newVehicle).Error

	return
}

func (r *mysqlVehicleRepository) GetList(ctx context.Context, userID uint) (list []entities.Vehicle, err error) {
	list = []entities.Vehicle{}
	result := []vehicle{}

	err = r.Conn.WithContext(ctx).Model(&vehicle{}).
		Where(&vehicle{UserID: userID}).
		Select(`id, user_id, name, license, company, model`).
		Find(&result).Error
	if err != nil {
		return
	}

	for _, v := range result {
		list = append(list, entities.Vehicle{
			ID:      v.ID,
			Name:    v.Name,
			License: v.License,
			Company: v.Company,
			Model:   v.Model,
		})
	}

	return
}

func (r *mysqlVehicleRepository) Get(ctx context.Context, id uint) (v entities.VehicleDetail, err error) {
	result := vehicle{}
	err = r.Conn.WithContext(ctx).Model(&vehicle{}).
		Where(&vehicle{ID: id}).
		First(&result).Error
	if err != nil {
		return
	}

	v = entities.VehicleDetail{
		ID:                  result.ID,
		Name:                result.Name,
		License:             result.License,
		Company:             result.Company,
		Model:               result.Model,
		EngineDisplacement:  result.EngineDisplacement,
		EngineNumber:        result.EngineNumber,
		DefaultOctaneNumber: result.DefaultOctaneNumber,
		Purchase:            result.Purchase,
		PurchaseDate:        result.PurchaseDate,
		PurchaseLocation:    result.PurchaseLocation,
		PurchaseMileage:     result.PurchaseMileage,
		Sold:                result.Sold,
		SoldDate:            result.SoldDate,
		SoldMileage:         result.SoldMileage,
		MileageReset:        result.MileageReset,
	}

	return
}

func (r *mysqlVehicleRepository) Edit(ctx context.Context, userID uint, v entities.VehicleDetail) (err error) {
	record := vehicle{
		ID:                  v.ID,
		UserID:              userID,
		Name:                v.Name,
		License:             v.License,
		Company:             v.Company,
		Model:               v.Model,
		EngineDisplacement:  v.EngineDisplacement,
		EngineNumber:        v.EngineNumber,
		DefaultOctaneNumber: v.DefaultOctaneNumber,
		Purchase:            v.Purchase,
		PurchaseDate:        v.PurchaseDate,
		PurchaseLocation:    v.PurchaseLocation,
		PurchaseMileage:     v.PurchaseMileage,
		Sold:                v.Sold,
		SoldDate:            v.SoldDate,
		SoldMileage:         v.SoldMileage,
		MileageReset:        v.MileageReset,
	}

	err = r.Conn.WithContext(ctx).Save(&record).Error

	return
}

func (r *mysqlVehicleRepository) Delete(ctx context.Context, id uint) (err error) {
	record := vehicle{ID: id}
	err = r.Conn.WithContext(ctx).Delete(&record).Error

	return
}

func (r *mysqlVehicleRepository) IsLicenseExist(ctx context.Context, userID uint, license string) bool {
	var count int64
	r.Conn.WithContext(ctx).Model(&vehicle{}).Where(&vehicle{UserID: userID, License: license}).Count(&count)

	return count > 0
}
