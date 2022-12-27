package protocol

import (
	"context"

	"github.com/mehdieidi/carify/data/pkg/divar"
)

type CarID uint

type Car struct {
	ID                     CarID               `json:"id"`
	Token                  string              `json:"token"`
	Year                   int                 `json:"year"`
	Color                  divar.Color         `json:"color"`
	UsageKM                int                 `json:"usage_km"`
	BodyStatus             divar.BodyStatus    `json:"body_status"`
	CashCost               int                 `json:"cash_cost"`
	MotorStatus            divar.MotorStatus   `json:"motor_status"`
	FrontChassisStatus     divar.ChassisStatus `json:"front_chassis_status"`
	RearChassisStatus      divar.ChassisStatus `json:"rear_chassis_status"`
	ThirdPartyInsuranceDue int                 `json:"third_party_insurance_due"`
	Gearbox                divar.Gearbox       `json:"gearbox"`
}

type CarToken string

type CarStorage interface {
	Store(context.Context, Car) (CarID, error)
	FindByToken(context.Context, CarToken) (Car, error)
	FindByID(context.Context, CarID) (Car, error)
	Update(context.Context, CarID, Car) error
	Delete(context.Context, CarID) error
}

type CarService interface {
	Search(ctx context.Context, cities []string, brandModels []string) ([]CarToken, error)
	Get(context.Context, CarToken) (Car, error)
	GetByID(context.Context, CarID) (Car, error)
	Update(context.Context, CarID, Car) error
	Delete(context.Context, CarID) error
}
