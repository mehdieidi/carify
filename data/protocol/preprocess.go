package protocol

import "context"

type PreProcessStorage interface {
	FindAll(context.Context) ([]Car, error)
}

type PreProcessService interface {
	List(context.Context) ([]Car, error)
	Year(context.Context) error
	Color(context.Context) error
	UsageKM(context.Context) error
	BodyStatus(context.Context) error
	CashCost(context.Context) error
	MotorStatus(context.Context) error
	FrontChassisStatus(context.Context) error
	RearChassisStatus(context.Context) error
	InsuranceDue(context.Context) error
	GearBox(context.Context) error
}
