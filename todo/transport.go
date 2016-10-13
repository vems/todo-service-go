package todo

import (
	"errors"
	"golang.org/x/net/context"

	pb "github.com/vems/pb/todo"
	"github.com/vems/todo-service-go/model"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.TodoServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		all: grpctransport.NewServer(
			ctx,
			endpoints.AllEndpoint,
			DecodeGRPCAllRequest,
			EncodeGRPCTodosResponse,
			options...,
		),
		create: grpctransport.NewServer(
			ctx,
			endpoints.CreateEndpoint,
			DecodeGRPCCreateRequest,
			EncodeGRPCTodoResponse,
			options...,
		),
		find: grpctransport.NewServer(
			ctx,
			endpoints.FindEndpoint,
			DecodeGRPCFindRequest,
			EncodeGRPCTodoResponse,
			options...,
		),
		update: grpctransport.NewServer(
			ctx,
			endpoints.UpdateEndpoint,
			DecodeGRPCUpdateRequest,
			EncodeGRPCTodoResponse,
			options...,
		),
		delete: grpctransport.NewServer(
			ctx,
			endpoints.DeleteEndpoint,
			DecodeGRPCDeleteRequest,
			EncodeGRPCDeleteResponse,
			options...,
		),
		deleteAll: grpctransport.NewServer(
			ctx,
			endpoints.DeleteAllEndpoint,
			DecodeGRPCDeleteAllRequest,
			EncodeGRPCDeleteAllResponse,
			options...,
		),
	}
}

type grpcServer struct {
	all       grpctransport.Handler
	create    grpctransport.Handler
	find      grpctransport.Handler
	update    grpctransport.Handler
	delete    grpctransport.Handler
	deleteAll grpctransport.Handler
}

func (s *grpcServer) All(ctx context.Context, req *pb.AllRequest) (*pb.TodosResponse, error) {
	_, rep, err := s.all.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TodosResponse), nil
}

func DecodeGRPCTodosResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	var todos model.TodoList
	resp := grpcReply.(*pb.TodosResponse)

	for _, todo := range resp.Todos {
		todos = append(todos, &model.Todo{
			Id:        todo.Id,
			Title:     todo.Title,
			Completed: todo.Completed,
		})
	}

	return allResponse{Todos: todos}, nil
}

func EncodeGRPCTodosResponse(_ context.Context, response interface{}) (interface{}, error) {
	var todos []*pb.TodoResponse
	resp := response.(allResponse)

	for _, todo := range resp.Todos {
		todos = append(todos, &pb.TodoResponse{
			Id:        todo.Id,
			Title:     todo.Title,
			Completed: todo.Completed,
		})
	}
	return &pb.TodosResponse{todos}, nil
}

func DecodeGRPCAllRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return allRequest{}, nil
}

func EncodeGRPCAllRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &pb.AllRequest{}, nil
}

func (s *grpcServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.TodoResponse, error) {
	_, rep, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TodoResponse), nil
}

func DecodeGRPCTodoResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	response := grpcReply.(*pb.TodoResponse)

	todo := &model.Todo{
		Id:        response.Id,
		Title:     response.Title,
		Completed: response.Completed,
	}

	return todo, nil
}

func EncodeGRPCTodoResponse(_ context.Context, svcResp interface{}) (interface{}, error) {
	response := svcResp.(*model.Todo)
	return &pb.TodoResponse{
		Id:        response.Id,
		Title:     response.Title,
		Completed: response.Completed,
	}, nil
}

func DecodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.CreateRequest).Todo
	newTodo := model.Todo{
		Title:     request.Title,
		Completed: request.Completed,
	}
	return newTodo, nil
}

func EncodeGRPCCreateRequest(_ context.Context, svcReq interface{}) (interface{}, error) {
	request := svcReq.(model.Todo)
	newTodo := &pb.TodoRequest{
		Title:     request.Title,
		Completed: request.Completed,
	}
	return &pb.CreateRequest{
		Todo: newTodo,
	}, nil
}

func (s *grpcServer) Find(ctx context.Context, req *pb.FindRequest) (*pb.TodoResponse, error) {
	_, rep, err := s.find.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TodoResponse), nil
}

func DecodeGRPCFindRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.FindRequest)
	return request.Id, nil
}

func EncodeGRPCFindRequest(_ context.Context, svcReq interface{}) (interface{}, error) {
	request := svcReq.(int64)
	return &pb.FindRequest{
		Id: request,
	}, nil
}

func (s *grpcServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.TodoResponse, error) {
	_, rep, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TodoResponse), nil
}

func DecodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.UpdateRequest)
	updateTodo := model.Todo{
		Title:     request.Todo.Title,
		Completed: request.Todo.Completed,
	}
	return updateRequest{
		Id:   request.Id,
		Todo: updateTodo,
	}, nil
}

func EncodeGRPCUpdateRequest(_ context.Context, svcReq interface{}) (interface{}, error) {
	request := svcReq.(updateRequest)
	return &pb.UpdateRequest{
		Id: request.Id,
		Todo: &pb.TodoRequest{
			Title:     request.Todo.Title,
			Completed: request.Todo.Completed,
		},
	}, nil
}

func (s *grpcServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	_, rep, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteResponse), nil
}

func DecodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.DeleteRequest)
	return request.Id, nil
}

func EncodeGRPCDeleteRequest(_ context.Context, svcReq interface{}) (interface{}, error) {
	id := svcReq.(int64)
	return &pb.DeleteRequest{
		Id: id,
	}, nil
}

func DecodeGRPCDeleteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	response := grpcReply.(*pb.DeleteResponse)

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	todo := &model.Todo{
		Id:        response.Todo.Id,
		Title:     response.Todo.Title,
		Completed: response.Todo.Completed,
	}

	return todo, nil
}

func EncodeGRPCDeleteResponse(_ context.Context, svcResp interface{}) (interface{}, error) {
	response := svcResp.(*model.Todo)
	return &pb.DeleteResponse{
		Todo: &pb.TodoResponse{
			Id:        response.Id,
			Title:     response.Title,
			Completed: response.Completed,
		},
	}, nil
}

func (s *grpcServer) DeleteAll(ctx context.Context, req *pb.DeleteAllRequest) (*pb.DeleteAllResponse, error) {
	_, rep, err := s.deleteAll.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteAllResponse), nil
}

func DecodeGRPCDeleteAllRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return deleteAllRequest{}, nil
}

func EncodeGRPCDeleteAllRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return &pb.DeleteAllRequest{}, nil
}

func DecodeGRPCDeleteAllResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return deleteAllResponse{}, nil
}

func EncodeGRPCDeleteAllResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &pb.DeleteAllResponse{}, nil
}
