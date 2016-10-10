package main

import (
	"log"
	"net"
	"os"

	"github.com/vems/todo-service-go/todo"
	"google.golang.org/grpc"

	pb "github.com/vems/pb/todo"
)

const (
	port = "PORT"
)

func main() {
	port := os.Getenv(port)
	// default for port
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("[Error][Server] Could not listen on port %v. %v", port, err)
	}
	defer lis.Close()

	s := grpc.NewServer()

	todoSrv, err := todo.NewTodo()
	if err != nil {
		log.Fatalf("[Error][Server] Could not start todo server: %v.", err)
	}

	pb.RegisterTodoServer(s, todoSrv)

	log.Printf("[Info][Server] Starting server on port %v", port)
	log.Printf("[Info][Server] The server has been stopped: %v", s.Serve(lis))
}
