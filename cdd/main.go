package main

import (
	"context"
	"gollow/cdd/config"
	"gollow/cdd/core/snapshot"
	"gollow/cdd/core/storage"
	"gollow/cdd/producer"
	"gollow/cdd/server"
	"gollow/cdd/sources/datamodel/dummy"
)

func main() {
	//Register Data Model to be produced
	RegisterDataModel()
	//initialise all you want
	Init(context.Background())
	//initialise server
	server.ServerInit()
}

func Init(ctx context.Context) {

	//initialise everything here
	snapshotStorage := storage.NewStorage(config.GlobalConfig.Storage.AnnouncedVersion)
	snapshot.Init(snapshotStorage)

	go producer.ScheduleProducers()
}

func RegisterDataModel() {
	producer.Register(&dummy.DummyData{}, struct{}{})
}
