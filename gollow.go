package main

import (
	"gollow/core/snapshot"
	"gollow/logging"
	"gollow/producer"
	"gollow/sources/datamodel"
	"gollow/storage"
	"time"
)

func main() {

	announcedFileName := "announced.version"
	announcedVersionStorage := storage.NewStorage(announcedFileName)
	snapshot.Init(announcedVersionStorage)

	producer.Producer(announcedVersionStorage, &datamodel.DummyData{})

	logging.GetLogger().Info("complete")

	time.Sleep(5 * time.Minute)
}
