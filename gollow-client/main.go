package main

import (
	"golang.org/x/net/context"
	"gollow/cdd/core/snapshot"
	"gollow/cdd/core/storage"
	"gollow/cdd/sources/datamodel/dummy"
	"gollow/gollow-client/cache"
	"gollow/gollow-client/cache/client_datamodel"
	"gollow/gollow-client/config"
	"gollow/gollow-client/server"
)

func main() {

	ctx := context.Background()
	RegisterDataModels()

	Init(ctx)

	server.ServerInit(ctx)
}

func Init(ctx context.Context) {

	//initialise everything here
	snapshotStorage := storage.NewStorage(config.GlobalConfig.AnnouncedVersion)
	snapshot.Init(snapshotStorage)

	go cache.ReadValue()
	cache.UpdateSnapshots(ctx)

}

func RegisterDataModels() {
	cache.Register(&dummy.DummyData{}, client_datamodel.DummyDataCache)
}
