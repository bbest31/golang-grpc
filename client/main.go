package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// import "github.com/bbest31/golang-grpc/client/echo"

func main() {

	ctx := context.Background()
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := echo.NewEchoServerClient(conn)
	response, err := client.Echo(ctx, &echo.EchoRequest{
		Message: "Hello World!",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Got from server: ", response.Response)

}
