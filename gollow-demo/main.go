package main

import (
	"golang.org/x/net/context"

	"gollow/cache"
	"gollow/core/snapshot"
	"gollow/gollow-demo/cache/client_datamodel"
	"gollow/gollow-demo/config"
	"gollow/gollow-demo/server"
	"gollow/sources/datamodel/dummy"
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
