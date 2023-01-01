package setting

import (
	"context"

	"back/pkg/log"
	"back/protocol"
)

type loggingMW struct {
	logger log.Logger
	next   protocol.SiteSettingService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next protocol.SiteSettingService) protocol.SiteSettingService {
		return loggingMW{
			logger: logger,
			next:   next,
		}
	}
}

func (l loggingMW) Get(ctx context.Context) (resp protocol.SiteSetting, err error) {
	defer func() {
		if err != nil {
			l.logger.Error(domain, log.ServiceLayer, "Get", log.Args{log.LogErrKey: err})
			return
		}
		l.logger.Info(domain, log.ServiceLayer, "Get", log.Args{log.LogRespKey: resp})
	}()
	defer l.logger.PanicHandler()

	resp, err = l.next.Get(ctx)

	return
}
