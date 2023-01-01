package predict

import (
	"context"

	"back/pkg/log"
	"back/protocol"
)

type loggingMW struct {
	logger log.Logger
	next   protocol.PredictService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next protocol.PredictService) protocol.PredictService {
		return loggingMW{
			logger: logger,
			next:   next,
		}
	}
}

func (l loggingMW) Predict(ctx context.Context, c protocol.CarData) (resp int, err error) {
	defer func() {
		if err != nil {
			l.logger.Error(domain, log.ServiceLayer, "Predict", log.Args{log.LogErrKey: err})
			return
		}
		l.logger.Info(domain, log.ServiceLayer, "Predict", log.Args{log.LogRespKey: resp})
	}()
	defer l.logger.PanicHandler()

	resp, err = l.next.Predict(ctx, c)

	return
}
