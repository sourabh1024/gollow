package main

import (
	"net"
	"log"
	"fmt"
	"gollow/api"
	"google.golang.org/grpc"
)

// main start a gRPC server and waits for connection
func main()  {

	fmt.Println("Starting server...")
	// create a listener on TCP port 2424
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6574))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		fmt.Println("Started server on port 6574")
	}

	// create a server instance
	s := api.Server{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	api.RegisterPingServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

