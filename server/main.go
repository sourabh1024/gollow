package main

import (
	"net"
	"log"
	"fmt"
	"gollow/api"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
)



//startGRPC server creates a grpcServer
// Parameters : address
func startGRPCServer(address string) error{

	log.Printf("Starting GRPC server at : %s", address)

	// create a listener on TCP
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("Failed to start GRPC server: %v", err)
	}

	log.Printf("Started GRPC server on : %s", address)

	// create a server instance
	s := api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	api.RegisterPingServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %s", err)
	}

	return nil
}

func startRestServer(address , grpcAddress string) error{

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Setup the client gRPC options
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register ping

	err := api.RegisterPingHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("ould not register service Ping: %s", err)
	}
	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}


// main start a gRPC server and waits for connection
func main()  {

	log.Println("Starting server...")

	restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)
	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)

	go func() {
		err := startGRPCServer(grpcAddress)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	go func() {
		err := startRestServer(restAddress, grpcAddress)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// infinite loop
	log.Printf("Entering infinite loop")
	select {}
}

