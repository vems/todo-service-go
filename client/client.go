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

	return todo.Endpoints{
		AllEndpoint: allEndpoint,
	}
}
