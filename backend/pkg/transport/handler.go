package transport

import (
	"context"

	"back/pkg/log"

	transport "back/pkg/kittransport"
)

func ErrorHandler(logger log.Logger, domain string) transport.ErrorHandlerFunc {
	return func(_ context.Context, err error) {
		logger.Error(domain, log.TransportLayer, "ErrorHandler", log.Args{log.LogErrKey: err})
	}
}
