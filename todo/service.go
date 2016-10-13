package todo

import (
	"errors"
	"github.com/vems/todo-service-go/model"
	"golang.org/x/net/context"
)

type Service interface {
	All(ctx context.Context) (model.TodoList, error)
	Create(ctx context.Context, req model.Todo) (*model.Todo, error)
	Find(ctx context.Context, id int64) (*model.Todo, error)
	Update(ctx context.Context, id int64, req model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, id int64) (*model.Todo, error)
	DeleteAll(ctx context.Context) error
}

func NewBasicService() Service {
	return &basicService{}
}

type basicService struct {
	id    int64
	Todos model.TodoList
}

func (s basicService) All(_ context.Context) (model.TodoList, error) {
	return s.Todos, nil
}

func (s *basicService) Create(_ context.Context, req model.Todo) (*model.Todo, error) {
	s.id += 1
	req.Id = s.id
	s.Todos = append(s.Todos, &req)
	return &req, nil
}

func (s basicService) Find(_ context.Context, id int64) (*model.Todo, error) {
	var match int
	hasMatch := false
	for i, todo := range s.Todos {
		if todo.Id == id {
			match = i
			hasMatch = true
			break
		}
	}

	if hasMatch {
		return s.Todos[match], nil
	}

	return nil, errors.New("Failed to find match")
}

func (s *basicService) Update(_ context.Context, id int64, request model.Todo) (*model.Todo, error) {
	var match int
	hasMatch := false
	for i, todo := range s.Todos {
		if todo.Id == id {
			match = i
			hasMatch = true
			break
		}
	}

	if hasMatch {
		updatedTodo := &request
		updatedTodo.Id = s.Todos[match].Id // Don't change Id
		s.Todos[match] = updatedTodo
		return updatedTodo, nil
	}

	return nil, errors.New("Failed to find match")
}

func (s *basicService) Delete(_ context.Context, id int64) (*model.Todo, error) {
	var match int
	hasMatch := false
	for i, todo := range s.Todos {
		if todo.Id == id {
			match = i
			hasMatch = true
			break
		}
	}

	if hasMatch {
		removedTodo := s.Todos[match]
		s.Todos = append(s.Todos[:match], s.Todos[match+1:]...)
		return removedTodo, nil
	}

	return nil, errors.New("Failed to find match")
}

func (s *basicService) DeleteAll(_ context.Context) error {
	s.Todos = model.TodoList{}
	return nil
}
