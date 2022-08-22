package mysql

import (
	"context"

	"gorm.io/gorm"

	"car-record/entities"
)

func (userGasStation) TableName() string {
	return "user_gas_station"
}

type userGasStation struct {
	ID     uint   `gorm:"primaryKey;column:id"`
	UserID uint   `grom:"column:user_id"`
	Name   string `gorm:"column:name"`
}

type mysqlUserGasStationRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserGasStationRepository(Conn *gorm.DB) entities.UserGasStationRepository {
	return &mysqlUserGasStationRepository{Conn}
}

func (r *mysqlUserGasStationRepository) Add(ctx context.Context, userID uint, name string) (err error) {
	newUserGasStation := userGasStation{
		UserID: userID,
		Name:   name,
	}
	err = r.Conn.WithContext(ctx).Create(&newUserGasStation).Error

	return
}
func (r *mysqlUserGasStationRepository) GetList(ctx context.Context, userID uint) (list []entities.UserGasStation, err error) {
	list = []entities.UserGasStation{}
	result := []userGasStation{}

	err = r.Conn.WithContext(ctx).Model(&userGasStation{}).
		Where(&userGasStation{UserID: userID}).
		Find(&result).Error
	if err != nil {
		return
	}

	for _, v := range result {
		list = append(list, entities.UserGasStation{
			ID:     v.ID,
			UserID: userID,
			Name:   v.Name,
		})
	}

	return
}
func (r *mysqlUserGasStationRepository) Delete(ctx context.Context, userID uint, id uint) (err error) {
	record := userGasStation{
		ID:     id,
		UserID: userID,
	}

	err = r.Conn.WithContext(ctx).Delete(&record).Error

	return
}
