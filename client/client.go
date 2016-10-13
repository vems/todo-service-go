package client

import (
	"time"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"

	pb "github.com/vems/pb/todo"
	"github.com/vems/todo-service-go/todo"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

func New(conn *grpc.ClientConn, logger log.Logger) todo.Service {

	var allEndpoint endpoint.Endpoint
	{
		allEndpoint = grpctransport.NewClient(
			conn,
			"Todo",
			"All",
			todo.EncodeGRPCAllRequest,
			todo.DecodeGRPCTodosResponse,
			pb.TodosResponse{},
		).Endpoint()
		allEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "All",
			Timeout: 30 * time.Second,
		}))(allEndpoint)
	}

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = grpctransport.NewClient(
			conn,
			"Todo",
			"Create",
			todo.EncodeGRPCCreateRequest,
			todo.DecodeGRPCTodoResponse,
			pb.TodoResponse{},
		).Endpoint()
		createEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Create",
			Timeout: 30 * time.Second,
		}))(createEndpoint)
	}

	var findEndpoint endpoint.Endpoint
	{
		findEndpoint = grpctransport.NewClient(
			conn,
			"Todo",
			"Find",
			todo.EncodeGRPCFindRequest,
			todo.DecodeGRPCTodoResponse,
			pb.TodoResponse{},
		).Endpoint()
		findEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Find",
			Timeout: 30 * time.Second,
		}))(findEndpoint)
	}

	var updateEndpoint endpoint.Endpoint
	{
		updateEndpoint = grpctransport.NewClient(
			conn,
			"Todo",
			"Update",
			todo.EncodeGRPCUpdateRequest,
			todo.DecodeGRPCTodoResponse,
			pb.TodoResponse{},
		).Endpoint()
		updateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Update",
			Timeout: 30 * time.Second,
		}))(updateEndpoint)
	}

	var deleteEndpoint endpoint.Endpoint
	{
		deleteEndpoint = grpctransport.NewClient(
			conn,
			"Todo",
			"Delete",
			todo.EncodeGRPCDeleteRequest,
			todo.DecodeGRPCDeleteResponse,
			pb.DeleteResponse{},
		).Endpoint()
		deleteEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Delete",
			Timeout: 30 * time.Second,
		}))(deleteEndpoint)
	}

	var deleteAllEndpoint endpoint.Endpoint
	{
		deleteAllEndpoint = grpctransport.NewClient(
			conn,
			"Todo",
			"DeleteAll",
			todo.EncodeGRPCDeleteAllRequest,
			todo.DecodeGRPCDeleteAllResponse,
			pb.DeleteAllResponse{},
		).Endpoint()
		deleteAllEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "DeleteAll",
			Timeout: 30 * time.Second,
		}))(deleteAllEndpoint)
	}

	return todo.Endpoints{
		AllEndpoint:       allEndpoint,
		CreateEndpoint:    createEndpoint,
		FindEndpoint:      findEndpoint,
		UpdateEndpoint:    updateEndpoint,
		DeleteEndpoint:    deleteEndpoint,
		DeleteAllEndpoint: deleteAllEndpoint,
	}
}
