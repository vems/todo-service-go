package todo

import (
	"github.com/vems/todo-service-go/model"
	"golang.org/x/net/context"
)

type Service interface {
	All(ctx context.Context) (model.TodoList, error)
}

func NewBasicService() Service {
	return basicService{}
}

type basicService struct{}

func (s basicService) All(_ context.Context) (model.TodoList, error) {
	examples := model.TodoList{
		&model.Todo{
			Id:        1,
			Title:     "Hello, World",
			Completed: false,
			Order:     0,
			Url:       "",
		},
	}
	return examples, nil
}
