package main

import (
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

	//announcedVersion, err := c.GetAnnouncedVersion(context.Background(), &api.AnnouncedVersionRequest{Namespace: "test-consumer-group1", Entity: "heatmap_data"})

	//cache.BuildSnapshot(announcedVersion.Currentversion)

	//if err != nil {
	//	log.Fatalf("Error when calling SayHello: %s", err)
	//}
	//log.Printf("Response from server: %s", announcedVersion.Currentversion)

	//h := cache.GetHeatMapDataInstance()
	//val, err := h.GetValue("1")
	//
	//if err != nil {
	//	logging.GetLogger().Error("Error in getting value : ", err)
	//	return
	//}

	//data, ok := val.(*datamodel.HeatMapData)
	//
	//if !ok {
	//	logging.GetLogger().Error("Error in typecasting value : ", err)
	//	return
	//}

	//logging.GetLogger().Info(fmt.Sprintf("ID %d , geoHash %s , vehicle : %d",
	//	data.ID, data.Geohash, data.VehicleTypeID))

	go UpdateSnapshots(c, datamodel.HeatMapDataRef)
	go ReadValue(c, datamodel.HeatMapDataRef)
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
				cache.FetchSnapshot(c, model)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
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
				val, err := cache.GetHeatMapDataInstance().GetValue("1")
				logging.GetLogger().Info("Size of cache : ", cache.GetHeatMapDataInstance().Size())
				if err != nil {
					logging.GetLogger().Error("Error in reading value : ", err)
					continue
				}
				logging.GetLogger().Info("Value for id 1 : ", val)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
