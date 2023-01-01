package setting

import (
	"context"
	"net/http"

	kithttp "back/pkg/kittransport/http"
	"back/pkg/log"
	"back/pkg/transport"
	"back/protocol"

	"github.com/go-chi/chi/v5"
)

func decodeGetSettingsRequest(_ context.Context, r *http.Request) (any, error) {
	var req getSettingsRequest
	return req, nil
}

func MakeHTTPHandler(s protocol.SiteSettingService, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.ErrorHandler(logger, domain)),
	}

	getSettingsHandler := kithttp.NewServer(
		makeGetSettingsEndpoint(s),
		decodeGetSettingsRequest,
		transport.EncodeResponse,
		transport.EncodeError,
		opts...,
	)

	mux := chi.NewRouter()

	mux.Method("GET", "/v1/site/settings/get", getSettingsHandler)

	return mux
}
