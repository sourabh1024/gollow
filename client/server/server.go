package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"gollow/client/client_api"
	"gollow/logging"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

//startGRPC server creates a grpcServer
// Parameters : address
func startGRPCServer(ctx context.Context, address string) error {

	logging.GetLogger().Info("Starting GRPC server at : %s", address)

	// create a listener on TCP
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("Failed to start GRPC server: %v", err)
	}

	logging.GetLogger().Info("Started GRPC server on : %s", address)

	// create a server instance
	s := client_api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	client_api.RegisterPingServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %s", err)
	}

	return nil
}

func startRestServer(ctx context.Context, address, grpcAddress string) error {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Setup the client gRPC options
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register ping

	err := client_api.RegisterPingHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("ould not register service Ping: %s", err)
	}
	logging.GetLogger().Info("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}

// main start a gRPC server and waits for connection
func ServerInit(ctx context.Context) {

	logging.GetLogger().Info("Starting server...")

	restAddress := fmt.Sprintf("%s:%d", "localhost", 8888)
	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 8889)

	go func() {
		err := startGRPCServer(ctx, grpcAddress)
		if err != nil {
			logging.GetLogger().Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	go func() {
		err := startRestServer(ctx, restAddress, grpcAddress)
		if err != nil {
			logging.GetLogger().Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	logging.GetLogger().Info("Entering infinite loop")

	// infinite loop
	select {}
}
