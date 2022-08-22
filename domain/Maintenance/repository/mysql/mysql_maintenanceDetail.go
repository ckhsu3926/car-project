package mysql

import (
	"context"

	"car-record/entities"

	"gorm.io/gorm"
)

func (maintenanceRecordDetail) TableName() string {
	return "maintenance_detail"
}

type maintenanceRecordDetail struct {
	ID                  uint   `gorm:"primaryKey;column:id"`
	MaintenanceRecordID uint   `gorm:"column:maintenance_record_id"`
	Name                string `gorm:"column:name"`
	Value               int    `gorm:"column:value"`
	Content             string `gorm:"column:content"`
}

func (r *mysqlMaintenanceRepository) GetDetailList(ctx context.Context, recordID uint) (list []entities.MaintenanceRecordDetail, err error) {
	list = []entities.MaintenanceRecordDetail{}
	result := []maintenanceRecordDetail{}

	err = r.Conn.WithContext(ctx).Model(&maintenanceRecordDetail{}).
		Where(&maintenanceRecordDetail{MaintenanceRecordID: recordID}).
		Find(&result).Error
	if err != nil {
		return
	}

	for _, v := range result {
		list = append(list, entities.MaintenanceRecordDetail{
			Name:    v.Name,
			Value:   v.Value,
			Content: v.Content,
		})
	}

	return
}

func (r *mysqlMaintenanceRepository) SetDetailList(ctx context.Context, recordID uint, detailList []entities.MaintenanceRecordDetail) (err error) {
	newList := []maintenanceRecordDetail{}
	for _, d := range detailList {
		newList = append(newList, maintenanceRecordDetail{
			MaintenanceRecordID: recordID,
			Name:                d.Name,
			Value:               d.Value,
			Content:             d.Content,
		})
	}

	err = r.Conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if deleteErr := tx.Where(&maintenanceRecordDetail{MaintenanceRecordID: recordID}).Delete(&maintenanceRecordDetail{}).Error; deleteErr != nil {
			return deleteErr
		}

		if createErr := tx.Create(&newList).Error; createErr != nil {
			return createErr
		}

		return nil
	})

	return
}
