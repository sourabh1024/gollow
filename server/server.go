package server

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"gollow/logging"
	"gollow/server/api"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

//startGRPC server creates a grpcServer
// Parameters : address
func startGRPCServer(address string) error {

	logging.GetLogger().Info("Starting GRPC server at : %s", address)

	// create a listener on TCP
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("Failed to start GRPC server: %v", err)
	}

	logging.GetLogger().Info("Started GRPC server on : %s", address)

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

func startRestServer(address, grpcAddress string) error {

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
	logging.GetLogger().Info("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}

// main start a gRPC server and waits for connection
func ServerInit() {

	logging.GetLogger().Info("Starting server...")

	restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)
	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)

	go func() {
		err := startGRPCServer(grpcAddress)
		if err != nil {
			logging.GetLogger().Error("failed to start gRPC server: %s", err)
		}
	}()

	go func() {
		err := startRestServer(restAddress, grpcAddress)
		if err != nil {
			logging.GetLogger().Error("failed to start gRPC server: %s", err)
		}
	}()
	// infinite loop
	logging.GetLogger().Info("Entering infinite loop")
	select {}
}
