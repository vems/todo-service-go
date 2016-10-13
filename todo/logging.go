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

func (mw serviceLoggingMiddleware) Create(ctx context.Context, newTodo model.Todo) (todo *model.Todo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Create",
			"todo_id", todo.Id,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Create(ctx, newTodo)
}

func (mw serviceLoggingMiddleware) Find(ctx context.Context, id int64) (todo *model.Todo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Find",
			"find_id", id,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Find(ctx, id)
}

func (mw serviceLoggingMiddleware) Update(ctx context.Context, id int64, todo model.Todo) (todoRes *model.Todo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Update",
			"update_id", id,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Update(ctx, id, todo)
}

func (mw serviceLoggingMiddleware) DeleteAll(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "DeleteAll",
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.DeleteAll(ctx)
}

func (mw serviceLoggingMiddleware) Delete(ctx context.Context, id int64) (todoRes *model.Todo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Delete",
			"delete_id", id,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Delete(ctx, id)
}
