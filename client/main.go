package main

import (
	"golang.org/x/net/context"
	"gollow/client/cache"
	"gollow/client/cache/client_datamodel"
	"gollow/client/server"
	"gollow/config"
	"gollow/core/snapshot"
	"gollow/core/storage"
	"gollow/logging"
	"gollow/server/api"
	"gollow/sources/datamodel/dummy"
	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()
	RegisterDataModels()

	Init(ctx)

	server.ServerInit(ctx)
}

func Init(ctx context.Context) {

	var err error
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())

	if err != nil {
		logging.GetLogger().Fatal("did not connect: %s", err)
	}

	logging.GetLogger().Info("Client server started")
	client := api.NewPingClient(conn)

	response, err := client.SayHello(ctx, &api.PingMessage{Greeting: "foo"})

	if err != nil {
		logging.GetLogger().Fatal("Error when calling SayHello: %s", err)
	}
	logging.GetLogger().Info("Response from server: %s", response.Greeting)

	//initialise everything here
	snapshotStorage := storage.NewStorage(config.GlobalConfig.AnnouncedVersion)
	snapshot.Init(snapshotStorage)

	go cache.ReadValue()
	cache.UpdateSnapshots(ctx)

}

func RegisterDataModels() {
	cache.Register(&dummy.DummyData{}, client_datamodel.DummyDataCache)
}
