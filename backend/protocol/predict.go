package protocol

import "context"

type CarData struct {
	Year                    int `json:"year"`
	ColorValue              int `json:"color_value"`
	UsageKM                 int `json:"usage_km"`
	BodyStatusValue         int `json:"body_status_value"`
	MotorStatusValue        int `json:"motor_status_value"`
	RearChassisStatusValue  int `json:"rear_chassis_status_value"`
	FrontChassisStatusValue int `json:"front_chassis_status_value"`
	Insurance               int `json:"insurance"`
	Gearbox                 int `json:"gearbox"`
}

type PredictService interface {
	Predict(context.Context, CarData) (int, error)
}
