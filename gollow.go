package main

import (
	"encoding/gob"
	"github.com/stackimpact/stackimpact-go"
	"gollow/logging"
	"gollow/producer"
	"gollow/sources/datamodel"
	"gollow/write"
	"time"
)

func main() {

	agent := stackimpact.Start(stackimpact.Options{
		AgentKey: "f2a18802c1126a98e6175bd99e481ba72b33a2dc",
		AppName:  "MyGoApp",
	})

	span := agent.Profile()
	defer span.Stop()

	for i := 0; i < 1; i++ {

		// Start CPU profiler.
		agent.StartCPUProfiler()
		agent.StartBlockProfiler()

		/*
			1. Create announced.version file if not exists
		*/

		//dbData, err := datamodel.HeatMapDataRef.LoadAll()
		//
		//data := dbData.([]sources.DataModel)
		//
		//logging.GetLogger().Info("length of data returned  :", len(data))
		//logging.GetLogger().Info("Size of data returned  :", reflect.TypeOf(data).Size())
		//logging.GetLogger().Info("type of data returned  :", reflect.TypeOf(data))
		//profile.GetMemoryProfile()
		//
		//for i := 0; i < 5; i++ {
		//	logging.GetLogger().Info(data[i].GetPrimaryKey())
		//
		//}
		//
		//logging.GetLogger().Info(data[0].GetNameSpace() + "-" + data[0].GetDataName())
		//file := &write.FileWriter{
		//	FilePath: "/Users/sourabh.suman/gopath/src/gollow/snapshots/",
		//	FileName: data[0].GetNameSpace() + "-" + data[0].GetDataName() + "-" + time.Now().String(),
		//}
		//
		//marshalData, err := json.Marshal(data)
		//
		//if err != nil {
		//	logging.GetLogger().Error("Marshalling error : ", err)
		//}
		//_, err = write.WriteSnapshot(file, file.GetFullPath(), marshalData)
		//
		//if err != nil {
		//	logging.GetLogger().Error("Err in writing : ", err)
		//}
		//
		//var response []datamodel.HeatMapData
		//marshalData, err = write.ReadSnapshot(file, file.GetFullPath())
		//
		//if err != nil {
		//	logging.GetLogger().Error("Error in fetching the data from snapshot : ", err)
		//}
		//
		//err = json.Unmarshal(marshalData, &response)
		//
		//logging.GetLogger().Info("Length of unmarshalled data : ", len(response))

		file := &write.FileWriter{
			FilePath: "/Users/sourabh.suman/gopath/src/gollow/snapshots/",
		}
		//
		//snapshot, _ := json.Marshal(&snapshot2.Snapshot{
		//	AnnouncedSnapshot: map[string]string{
		//		"version": "1.0.0",
		//	},
		//})
		//
		//file.Write("/Users/sourabh.suman/gopath/src/gollow/snapshots/announced.version", snapshot)

		gob.Register(&datamodel.HeatMapData{})
		producer.Producer(file, "/Users/sourabh.suman/gopath/src/gollow/snapshots/announced.version", &datamodel.HeatMapData{})

		logging.GetLogger().Info("complete")
		agent.StopCPUProfiler()
		agent.StopBlockProfiler()
	}

	time.Sleep(5 * time.Minute)
}
