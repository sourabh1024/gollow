package main

import (
	"encoding/json"
	"gollow/logging"
	"gollow/sources"
	"gollow/sources/datamodel"
	"gollow/util/profile"
	"gollow/write"
	"reflect"
	"time"
)

func main() {

	//dataModels := &sources.DataModels{
	//	//	Models: make(map[string][]sources.DataModel),
	//	//}
	//	//
	//	//dataModels.Add(datamodel.HeatMapDataRef)
	//	//
	//	//models, ok := dataModels.GetDMInNameSpace(datamodel.HeatMapDataRef.GetNameSpace())
	//	//
	//	//if !ok {
	//	//	logging.GetLogger().Error("No model found in namespace ")
	//	//	return
	//	//}
	//	//
	//	//for i := 0; i < len(models); i++ {
	//	//	data := models[i].LoadAll()
	//	//	logging.GetLogger().Info("length of data returned  :", len(data))
	//	//	logging.GetLogger().Info("type of data returned  :", reflect.TypeOf(data))
	//	//}

	data := datamodel.HeatMapDataRef.LoadAll().([]sources.DataModel)

	logging.GetLogger().Info("length of data returned  :", len(data))
	logging.GetLogger().Info("Size of data returned  :", reflect.TypeOf(data).Size())
	logging.GetLogger().Info("type of data returned  :", reflect.TypeOf(data))
	profile.GetMemoryProfile()

	for i := 0; i < 5; i++ {
		logging.GetLogger().Info(data[i].GetPrimaryKey())

	}

	logging.GetLogger().Info(data[0].GetNameSpace() + "-" + data[0].GetDataName())
	file := &write.FileWriter{
		FilePath: "/Users/sourabh.suman/gopath/src/gollow/snapshots/",
		FileName: data[0].GetNameSpace() + "-" + data[0].GetDataName() + "-" + time.Now().String(),
	}

	marshalData, err := json.Marshal(data)

	if err != nil {
		logging.GetLogger().Error("Marshalling error : ", err)
	}
	_, err = write.WriteSnapshot(file, marshalData)

	if err != nil {
		logging.GetLogger().Error("Err in writing : ", err)
	}

	var response []datamodel.HeatMapData
	marshalData, err = write.ReadSnapshot(file)

	if err != nil {
		logging.GetLogger().Error("Error in fetching the data from snapshot : ", err)
	}

	err = json.Unmarshal(marshalData, &response)

	logging.GetLogger().Info("Length of unmarshalled data : ", len(response))

}
