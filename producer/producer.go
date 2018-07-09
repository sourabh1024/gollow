package producer

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"gollow/diff"
	"gollow/logging"
	"gollow/snapshot"
	"gollow/sources"
	"gollow/util"
	"gollow/write"
	"reflect"
)

/**
ProducerWorker is Worker method , it consumes the jobs from the channel
@param workerID : workerID of the worker
@param jobs : jobs channel in which all DataModel to be produced are pushed
@param results : result channel in which the results are pushed back
*/
func ProducerWorker(workerID int, jobs <-chan sources.DataModel, results chan<- interface{}) {

	//for j := range jobs {
	//
	//}
}

func StartProducer([]sources.DataModel) {

	//ticker := time.NewTicker(1 * time.Minute)

	//go func() {
	//
	//	select {
	//	<-ticker.C:
	//
	//	}
	//}()
}

type UniversalDTO struct {
	Data interface{} `json:"data"`
	// more fields with important meta-data about the message...
}

/**
Producer produces a given data source DataModel
@param model: DataModel to produce
@returns :
*/
func Producer(writer write.SnapshotReaderWriter, announcedVersionPath string, model sources.DataModel) {

	/*
	 0. Load the current Data
	 1. Get the previous announced version of the data.
	 2. De-Serialize and create back the old object
	 3. Load the current data from the data source
	 4. Get the diff of the data.
	 5. Serialize and create the new snapshot
	 6. Create the diff
	 7. Update the announced version
	*/

	logging.GetLogger().Info("Starting data producing for : ", model.GetDataName())

	data, err := model.LoadAll()
	if err != nil {
		logging.GetLogger().Error("Error in fetching current data  : ", err)
	}

	currData, ok := data.([]sources.DataModel)
	if !ok {
		logging.GetLogger().Error("Error in typecasting interface{} to []source.DatModel ")
		return
	}

	rawBytes, err := MarshalDataModels(currData)
	if err != nil {
		logging.GetLogger().Error("Error in marshalling current data bytes : ", err)
	}

	lastAnnouncedVersion, err := snapshot.GetLatestAnnouncedVersion(writer, announcedVersionPath, model.GetNameSpace()+"-"+model.GetDataName())
	if err != nil {
		logging.GetLogger().Error("Error in getting previous announced version : ", err)
		return
	}

	if lastAnnouncedVersion == "" {
		newAnnouncedVersion := "/Users/sourabh.suman/gopath/src/gollow/snapshots" + "/" + model.GetNameSpace() + "-" + model.GetDataName() + "-" + util.GetCurrentTimeString()
		writer.Write(newAnnouncedVersion, rawBytes)
		snapshot.UpdateLatestAnnouncedVersion(writer, announcedVersionPath, model.GetNameSpace()+"-"+model.GetDataName(), newAnnouncedVersion)
		return
	}

	prevBytes, err := writer.Read(lastAnnouncedVersion)
	if err != nil {
		logging.GetLogger().Error("Error in reading previous data : ", err)
		return
	}

	prevData, err := UnMarshalDataModelsBytes(prevBytes, model)

	delta := diff.GetNewDiffObj()
	delta.GetDiffBetweenModels(prevData, currData)
	diffPath := "/Users/sourabh.suman/gopath/src/gollow/snapshots" + "/" + "diff1"
	if shouldDiffBeProduced(delta) {
		deltaBytes, err := MarshalDiff(delta)
		if err != nil {
			logging.GetLogger().Error("Error in Marshalling Diff , err : ", err)
		}
		writer.Write(diffPath, deltaBytes)
	}

	createdDiff := readDiff(writer, diffPath)

	p, ok := createdDiff.NewObjects.([]interface{})

	if p != nil {
		logging.GetLogger().Info("Length of new object : ", len(p))
	}

	return
}

func MarshalDiff(delta *diff.Diff) ([]byte, error) {
	return json.Marshal(delta)
}

func UnMarshalDiffBytes(data []byte) (*diff.Diff, error) {
	d := diff.GetNewDiffObj()
	err := json.Unmarshal(data, &d)
	return d, err
}

func MarshalDataModels(data []sources.DataModel) ([]byte, error) {
	universalDto := &UniversalDTO{Data: data}
	return json.Marshal(universalDto)
}

func UnMarshalDataModelsBytes(data []byte, model sources.DataModel) ([]sources.DataModel, error) {

	oldData := &UniversalDTO{}
	err := json.Unmarshal(data, oldData)
	if err != nil {
		logging.GetLogger().Info("Error in unmarshalling old data bytes : ", err)
		return nil, err
	}

	dataInterface, ok := (oldData.Data).([]interface{})
	if !ok {
		logging.GetLogger().Error("Error in typecasting the oldData into interface array, Err :", err)
		return nil, errors.New("error in typecasting old data bytes")
	}

	models := make([]sources.DataModel, 0)
	for i := 0; i < len(dataInterface); i++ {
		dataRef := model.NewDataRef()
		mapstructure.Decode(dataInterface[i], dataRef)
		models = append(models, dataRef)
	}

	return models, nil
}

func shouldDiffBeProduced(diff *diff.Diff) bool {
	if diff == nil ||
		(reflect.TypeOf(diff.NewObjects).Size() == 0 &&
			reflect.TypeOf(diff.ChangedObjects).Size() == 0 &&
			len(diff.MissingKeys) == 0) {
		return false
	}
	return true
}

//
//func generateDiff(diff *diff.Diff, writer write.SnapshotWriter, diffPath string) {
//	diffBytes, err := diff.Marshal()
//	if err != nil {
//		logging.GetLogger().Error("Error in marshalling diff , err: ", err)
//	}
//
//	writer.Write(diffPath, diffBytes)
//}
//
func readDiff(reader write.SnapshotReader, diffPath string) *diff.Diff {
	data, err := reader.Read(diffPath)
	if err != nil {
		logging.GetLogger().Error("Error in reading diff : ", err)
	}

	d := diff.GetNewDiffObj()
	err = json.Unmarshal(data, &d)
	if err != nil {
		logging.GetLogger().Error("Error in Unmarshalling : ", err)
	}
	return d
}
