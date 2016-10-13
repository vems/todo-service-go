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

	var createEndpoint endpoint.Endpoint
	{
		createLogger := log.NewContext(logger).With("method", "Create")
		createEndpoint = MakeCreateEndpoint(service)
		createEndpoint = EndpointLoggingMiddleware(createLogger)(createEndpoint)
	}

	var findEndpoint endpoint.Endpoint
	{
		findLogger := log.NewContext(logger).With("method", "Find")
		findEndpoint = MakeFindEndpoint(service)
		findEndpoint = EndpointLoggingMiddleware(findLogger)(findEndpoint)
	}

	var updateEndpoint endpoint.Endpoint
	{
		updateLogger := log.NewContext(logger).With("method", "Update")
		updateEndpoint = MakeUpdateEndpoint(service)
		updateEndpoint = EndpointLoggingMiddleware(updateLogger)(updateEndpoint)
	}

	var deleteEndpoint endpoint.Endpoint
	{
		deleteLogger := log.NewContext(logger).With("method", "Delete")
		deleteEndpoint = MakeDeleteEndpoint(service)
		deleteEndpoint = EndpointLoggingMiddleware(deleteLogger)(deleteEndpoint)
	}

	var deleteAllEndpoint endpoint.Endpoint
	{
		deleteAllLogger := log.NewContext(logger).With("method", "DeleteAll")
		deleteAllEndpoint = MakeDeleteAllEndpoint(service)
		deleteAllEndpoint = EndpointLoggingMiddleware(deleteAllLogger)(deleteAllEndpoint)
	}

	endpoints := Endpoints{
		AllEndpoint:       allEndpoint,
		CreateEndpoint:    createEndpoint,
		FindEndpoint:      findEndpoint,
		UpdateEndpoint:    updateEndpoint,
		DeleteEndpoint:    deleteEndpoint,
		DeleteAllEndpoint: deleteAllEndpoint,
	}

	// Mechanical domain.
	ctx := context.Background()
	logger = log.NewContext(logger).With("transport", "gRPC")
	srv := MakeGRPCServer(ctx, endpoints, logger)
	return srv, nil
}
