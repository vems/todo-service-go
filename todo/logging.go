package todo

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/vems/todo-service-go/model"
	"golang.org/x/net/context"
)

type Middleware func(Service) Service

func EndpointLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.Log("error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)

		}
	}
}

func ServiceLoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return serviceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type serviceLoggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw serviceLoggingMiddleware) All(ctx context.Context) (todos model.TodoList, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "All",
			"todo_count", len(todos),
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.All(ctx)
}
