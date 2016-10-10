package todo

import (
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	pb "github.com/vems/pb/todo"
	"golang.org/x/net/context"
)

func NewTodo() (pb.TodoServer, error) {
	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}
	logger.Log("msg", "Service Started")

	// Business domain.
	var service Service
	{
		service = NewBasicService()
		service = ServiceLoggingMiddleware(logger)(service)
	}

	// Endpoint domain.
	var allEndpoint endpoint.Endpoint
	{
		allLogger := log.NewContext(logger).With("method", "All")
		allEndpoint = MakeAllEndpoint(service)
		allEndpoint = EndpointLoggingMiddleware(allLogger)(allEndpoint)
	}

	endpoints := Endpoints{
		AllEndpoint: allEndpoint,
	}

	// Mechanical domain.
	ctx := context.Background()
	logger = log.NewContext(logger).With("transport", "gRPC")
	srv := MakeGRPCServer(ctx, endpoints, logger)
	return srv, nil
}
