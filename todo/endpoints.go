package todo

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/vems/todo-service-go/model"
	"golang.org/x/net/context"
)

type Endpoints struct {
	AllEndpoint       endpoint.Endpoint
	CreateEndpoint    endpoint.Endpoint
	FindEndpoint      endpoint.Endpoint
	UpdateEndpoint    endpoint.Endpoint
	DeleteEndpoint    endpoint.Endpoint
	DeleteAllEndpoint endpoint.Endpoint
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

func MakeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.Todo)
		todo, err := s.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return todo, nil
	}
}

func (e Endpoints) Create(ctx context.Context, request model.Todo) (*model.Todo, error) {
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*model.Todo), nil
}

func MakeFindEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(int64)
		todo, err := s.Find(ctx, id)
		if err != nil {
			return nil, err
		}
		return todo, nil
	}
}

func (e Endpoints) Find(ctx context.Context, id int64) (*model.Todo, error) {
	response, err := e.FindEndpoint(ctx, id)
	if err != nil {
		return nil, err
	}
	return response.(*model.Todo), nil
}

// Update Todos endpoint
type updateRequest struct {
	Id   int64
	Todo model.Todo
}

func MakeUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateRequest)
		todo, err := s.Update(ctx, req.Id, req.Todo)
		if err != nil {
			return nil, err
		}
		return todo, nil
	}
}

func (e Endpoints) Update(ctx context.Context, id int64, request model.Todo) (*model.Todo, error) {
	response, err := e.UpdateEndpoint(ctx, updateRequest{id, request})
	if err != nil {
		return nil, err
	}
	return response.(*model.Todo), nil
}

// Delete Todo endpoint
func MakeDeleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(int64)
		todo, err := s.Delete(ctx, id)
		if err != nil {
			return nil, err
		}
		return todo, nil
	}
}

func (e Endpoints) Delete(ctx context.Context, id int64) (*model.Todo, error) {
	response, err := e.DeleteEndpoint(ctx, id)
	if err != nil {
		return nil, err
	}
	return response.(*model.Todo), nil
}

type deleteAllRequest struct{}
type deleteAllResponse struct{}

func MakeDeleteAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		err := s.DeleteAll(ctx)
		if err != nil {
			return nil, err
		}
		return deleteAllResponse{}, nil
	}
}

func (e Endpoints) DeleteAll(ctx context.Context) error {
	_, err := e.DeleteAllEndpoint(ctx, deleteAllRequest{})
	if err != nil {
		return err
	}
	return nil
}
