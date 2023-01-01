package transport

import (
	"context"

	"back/pkg/log"

	"bitbucket.imenaria.org/tool/toolkit/transport"
)

func ErrorHandler(logger log.Logger, domain string) transport.ErrorHandlerFunc {
	return func(_ context.Context, err error) {
		logger.Error(domain, log.TransportLayer, "ErrorHandler", log.Args{log.LogErrKey: err})
	}
}
