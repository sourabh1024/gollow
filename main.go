package main

import (
	"context"
	"gollow/config"
	"gollow/core/snapshot"
	"gollow/core/storage"
	"gollow/logging"
	"gollow/producer"
	"gollow/server"
	"gollow/sources/datamodel/dummy"
)

func main() {
	//Register Data Model to be produced
	RegisterDataModel()
	//initialise all you want
	Init(context.Background())
	//initialise server
	server.Init()
}

// Init initialises everything here
// should initialise which storage to use
// schedule the producers
func Init(ctx context.Context) {

	//initialise everything here
	snapshotStorage, err := storage.NewStorage(config.GlobalConfig.Storage.AnnouncedVersion)

	if err != nil {
		logging.GetLogger().Error("error in getting new snapshot storage with err : %v", err)
		panic(err)
	}
	snapshot.Init(snapshotStorage)

	go producer.ScheduleProducers()
}

// RegisterDataModel registers the data model for production
func RegisterDataModel() {
	producer.Register(&dummy.DummyData{}, struct{}{})
}
