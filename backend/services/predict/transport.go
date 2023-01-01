package predict

import (
	"net/http"

	kithttp "back/pkg/kittransport/http"
	"back/pkg/log"
	"back/pkg/transport"
	"back/protocol"

	"github.com/go-chi/chi/v5"
)

func MakeHTTPHandler(s protocol.PredictService, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.ErrorHandler(logger, domain)),
	}

	predictHandler := kithttp.NewServer(
		makePredictEndpoint(s),
		transport.DecodeJSONRequest[predictRequest],
		transport.EncodeResponse,
		transport.EncodeError,
		opts...,
	)

	mux := chi.NewRouter()

	mux.Method("POST", "/v1/costs/predict", predictHandler)

	return mux
}
