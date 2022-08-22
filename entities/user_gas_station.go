package entities

import "context"

type UserGasStation struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"userID"`
	Name   string `json:"name"`
}

type UserGasStationUsecase interface {
	Add(ctx context.Context, userID uint, name string) (err error)
	GetList(ctx context.Context, userID uint) (list []UserGasStation, err error)
	Delete(ctx context.Context, userID uint, id uint) (err error)
}

type UserGasStationRepository interface {
	Add(ctx context.Context, userID uint, name string) (err error)
	GetList(ctx context.Context, userID uint) (list []UserGasStation, err error)
	Delete(ctx context.Context, userID uint, id uint) (err error)
}
