package main

import (
	"golang.org/x/net/context"

	"github.com/sourabh1024/gollow/cache"
	"github.com/sourabh1024/gollow/core/snapshot"
	"github.com/sourabh1024/gollow/gollow-demo/cache/client_datamodel"
	"github.com/sourabh1024/gollow/gollow-demo/config"
	"github.com/sourabh1024/gollow/gollow-demo/server"
	"github.com/sourabh1024/gollow/sources/datamodel/dummy"
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
