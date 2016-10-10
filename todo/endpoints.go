package todo

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/vems/todo-service-go/model"
	"golang.org/x/net/context"
)

type Endpoints struct {
	AllEndpoint endpoint.Endpoint
}

// All Todos endpoint
type allRequest struct{}

type allResponse struct {
	Todos model.TodoList
}

func MakeAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		todos, err := s.All(ctx)
		if err != nil {
			return nil, err
		}
		return allResponse{todos}, nil
	}
}

func (e Endpoints) All(ctx context.Context) (model.TodoList, error) {
	request := allRequest{}
	response, err := e.AllEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(allResponse).Todos, nil
}
