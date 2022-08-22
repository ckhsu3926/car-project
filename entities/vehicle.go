package entities

import (
	"context"

	"github.com/shopspring/decimal"
)

type Vehicle struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	License string `json:"license"`
	Company string `json:"company"`
	Model   string `json:"model"`
}
type VehicleDetail struct {
	ID                  uint            `json:"id"`
	Name                string          `json:"name"`
	License             string          `json:"license"`
	Company             string          `json:"company"`
	Model               string          `json:"model"`
	EngineDisplacement  decimal.Decimal `json:"engineDisplacement"`
	EngineNumber        string          `json:"engineNumber"`
	DefaultOctaneNumber int             `json:"defaultOctaneNumber"`
	Purchase            int             `json:"purchase"`
	PurchaseDate        string          `json:"purchaseDate"`
	PurchaseLocation    string          `json:"purchaseLocation"`
	PurchaseMileage     decimal.Decimal `json:"purchaseMileage"`
	Sold                int             `json:"sold"`
	SoldDate            string          `json:"soldDate"`
	SoldMileage         decimal.Decimal `json:"soldMileage"`
	MileageReset        decimal.Decimal `json:"mileageReset"`
}

type VehicleUsecase interface {
	Add(ctx context.Context, userID uint, name, license, company, model string) (err error)
	GetList(ctx context.Context, userID uint) (list []Vehicle, err error)
	Get(ctx context.Context, id uint) (vehicle VehicleDetail, err error)
	Edit(ctx context.Context, userID uint, vehicle VehicleDetail) (err error)
	Delete(ctx context.Context, id uint) (err error)
}

type VehicleRepository interface {
	Add(ctx context.Context, userID uint, name, license, company, model string) (err error)
	GetList(ctx context.Context, userID uint) (list []Vehicle, err error)
	Get(ctx context.Context, id uint) (vehicle VehicleDetail, err error)
	Edit(ctx context.Context, userID uint, vehicle VehicleDetail) (err error)
	Delete(ctx context.Context, id uint) (err error)
	IsLicenseExist(ctx context.Context, userID uint, license string) bool
}
