package main

import (
	"golang.org/x/net/context"
	"gollow/cdd/cache"
	"gollow/cdd/core/snapshot"
	"gollow/cdd/sources/datamodel/dummy"
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
	//init everything here...
	snapshot.InitVersionStorage(config.GlobalConfig.Storage.AnnouncedVersion)
	cache.RefreshCaches(ctx)
}

func RegisterDataModels() {
	cache.Register(&dummy.DummyData{}, client_datamodel.DummyDataCache)
}
