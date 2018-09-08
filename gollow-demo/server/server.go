//Copyright 2018 Sourabh Suman ( https://github.com/sourabh1024 )
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/sourabh1024/gollow/gollow-demo/client_api"
	"github.com/sourabh1024/gollow/gollow-demo/config"
	"github.com/sourabh1024/gollow/logging"
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

	// Setup the cache gRPC options
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

	restAddress := fmt.Sprintf("%s:%d", "localhost", config.GlobalConfig.ServerRPCPort)
	grpcAddress := fmt.Sprintf("%s:%d", "localhost", config.GlobalConfig.ServerHttpPort)

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
