package main

import (
	"golang.org/x/net/context"

	"github.com/gollow/cache"
	"github.com/gollow/core/snapshot"
	"github.com/gollow/gollow-demo/cache/client_datamodel"
	"github.com/gollow/gollow-demo/config"
	"github.com/gollow/gollow-demo/server"
	"github.com/gollow/sources/datamodel/dummy"
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
