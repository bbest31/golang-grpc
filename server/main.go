package main

import (
	"context"
	"fmt"
	"net"
)

// import "github.com/bbest31/golang-grpc/server/echo"

type EchoServer struct{}

// Create a function for each of the rpc calls a server has defined in the protobuf file.
func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {

	return &echo.EchoResponse{
		Response: "Echo: " + req.Message,
	}, nil
}

func main() {
	// Create a net listener for the grpc server.
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// init a gRPC server.
	s := grpc.NewServer()

	srv := &EchoServer{}

	echo.RegisterEchoServerServer(s, srv)

	// Server our server
	err = s.Serve(lst)
	if err != nil {
		panic(err)
	}

	fmt.Println("Now serving on port 8080...")
}
