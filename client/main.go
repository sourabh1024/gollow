package main

import (
	"encoding/gob"
	"golang.org/x/net/context"
	"gollow/api"
	"gollow/client/cache"
	"gollow/logging"
	"gollow/sources"
	"gollow/sources/datamodel"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)

	go UpdateSnapshots(c, datamodel.DummyDataRef)
	go ReadValue(c, datamodel.DummyDataRef)
	select {}
}

func UpdateSnapshots(c api.PingClient, model sources.DataModel) {
	ticker := time.NewTicker(time.Duration(model.CacheDuration()))
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				logging.GetLogger().Info("Update Snapshot")
				cache.FetchSnapshot(c, model, cache.DummyDataCache)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	gob.Register()
}

func ReadValue(c api.PingClient, model sources.DataModel) {
	ticker := time.NewTicker(time.Duration(30 * time.Second))
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				logging.GetLogger().Info("Read data")
				val, err := cache.DummyDataCache.Get("1")
				if err != nil {
					logging.GetLogger().Error("Error in reading value : ", err)
					continue
				}
				dummyData, ok := val.(*datamodel.DummyData)
				if !ok {
					logging.GetLogger().Error("Error in parsing value : ")
					continue
				}
				logging.GetLogger().Info("created at : ", dummyData.FirstName)
				logging.GetLogger().Info("Value for id 1 : ", val)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
