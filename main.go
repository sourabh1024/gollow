package main

import (
	"context"
	"gollow/config"
	"gollow/core/snapshot"
	"gollow/core/storage"
	"gollow/producer"
	"gollow/server"
	"gollow/sources/datamodel"
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
	snapshotStorage := storage.NewStorage(config.GlobalConfig.AnnouncedVersion)
	snapshot.Init(snapshotStorage)

	go producer.ScheduleProducers()
}

func RegisterDataModel() {
	producer.Register(datamodel.DummyDataRef, struct{}{})
}
