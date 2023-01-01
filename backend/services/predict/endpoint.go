package predict

import (
	"context"

	"back/pkg/endpoint"
	"back/pkg/transport"
	"back/protocol"
)

type predictRequest struct {
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

func makePredictEndpoint(s protocol.PredictService) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		request := req.(predictRequest)

		cost, err := s.Predict(ctx, protocol.CarData(request))
		if err != nil {
			return nil, err
		}

		return transport.Response{Data: cost}, nil
	}
}
