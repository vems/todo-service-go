package todo

import (
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
	}
}

type grpcServer struct {
	all grpctransport.Handler
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

func (s *grpcServer) Create(ctx context.Context, _ *pb.CreateRequest) (*pb.TodoResponse, error) {
	return &pb.TodoResponse{}, nil
}

func (s *grpcServer) Find(ctx context.Context, _ *pb.FindRequest) (*pb.TodoResponse, error) {
	return &pb.TodoResponse{}, nil
}

func (s *grpcServer) Update(ctx context.Context, _ *pb.UpdateRequest) (*pb.TodoResponse, error) {
	return &pb.TodoResponse{}, nil
}

func (s *grpcServer) DeleteAll(ctx context.Context, _ *pb.DeleteAllRequest) (*pb.DeleteAllResponse, error) {
	return &pb.DeleteAllResponse{}, nil
}
func (s *grpcServer) Delete(ctx context.Context, _ *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}
