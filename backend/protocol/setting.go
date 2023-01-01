package protocol

import "context"

type SiteSetting struct {
	Body []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"body"`
	FrontChassisStatus []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"front_chassis_status"`
	RearChassisStatus []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"rear_chassis_status"`
	Gearbox []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"gearbox"`
	MotorStatus []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"motor_status"`
	Color []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"color"`
}

type SiteSettingService interface {
	Get(context.Context) (SiteSetting, error)
}
