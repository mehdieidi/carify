package setting

import (
	"context"

	"back/pkg/endpoint"
	"back/pkg/transport"
	"back/protocol"
)

type getSettingsRequest struct{}

func makeGetSettingsEndpoint(s protocol.SiteSettingService) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		settings, err := s.Get(ctx)
		if err != nil {
			return nil, err
		}

		return transport.Response{Data: settings}, nil
	}
}
