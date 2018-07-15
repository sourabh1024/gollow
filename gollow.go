package main

import (
	"gollow/logging"
	"gollow/producer"
	"gollow/sources/datamodel"
	"gollow/storage"
	"time"
)

func main() {

	announcedFileName := "announced.version"
	announcedVersionStorage := storage.NewStorage(announcedFileName)
	producer.Producer(announcedVersionStorage, &datamodel.HeatMapData{})

	logging.GetLogger().Info("complete")

	time.Sleep(5 * time.Minute)
}
