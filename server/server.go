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

	// Setup the cache gRPC options
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

// Init start a gRPC server and waits for connection
func Init() {

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
