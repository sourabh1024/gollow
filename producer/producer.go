package producer

import (
	"encoding/json"
	"errors"
	"fmt"
	"gollow/core"
	"gollow/core/snapshot"
	"gollow/logging"
	"gollow/sources"
	"gollow/storage"
	"gollow/util"
	"time"
)

const (
	separator = "-"
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
}

/**
Producer produces a given data source DataModel
@param model: DataModel to produce
@returns :
*/
func Producer(announcedVersionStorage storage.Storage, model sources.DataModel) {

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

	defer util.Duration(time.Now(), fmt.Sprintf("Producer for : %s", model.GetDataName()))

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

	currBytes, err := MarshalDataModels(currData)
	if err != nil {
		logging.GetLogger().Error("Error in marshalling current data bytes : ", err)
	}

	lastAnnouncedSnapshot, err := snapshot.GetLatestAnnouncedVersion(announcedVersionStorage, announcedVersionKeyName(model))
	if err != nil {
		logging.GetLogger().Error("Error in getting previous announced version : ", err)
		return
	}

	if lastAnnouncedSnapshot == "" {
		newSnapshotFileName := getAnnouncedVersionName(model, -1)
		snapshot.WriteNewSnapshot(newSnapshotFileName, currBytes)
		snapshot.UpdateLatestAnnouncedVersion(announcedVersionStorage, announcedVersionKeyName(model), newSnapshotFileName)
		return
	}

	prevSnapshotVersion := snapshot.GetVersionNumber(lastAnnouncedSnapshot)
	newSnapshotName := getAnnouncedVersionName(model, prevSnapshotVersion)
	snapshot.WriteNewSnapshot(newSnapshotName, currBytes)

	snapshotStorage := storage.NewStorage(lastAnnouncedSnapshot)
	prevBytes, err := snapshotStorage.Read()
	if err != nil {
		logging.GetLogger().Error("Error in reading previous data : ", err)
		return
	}

	prevData, err := UnMarshalDataModelsBytes(prevBytes, model)
	//prevData, err := util.UnmarshalDataModelBytesFast(prevBytes, model)

	logging.GetLogger().Info("Generating diff for : ", model.GetDataName())
	logging.GetLogger().Info("DiffObject prevVersion , currVersion : ", prevSnapshotVersion, prevSnapshotVersion+1)
	err = core.DiffObjectDao.CreateNewDiff(model, prevData, currData, prevSnapshotVersion, prevSnapshotVersion+1)

	if err != nil {
		logging.GetLogger().Error("Error in producing diff for : "+newSnapshotName, err)
	}
	newSnapshotFileName := getAnnouncedVersionName(model, prevSnapshotVersion)

	snapshot.UpdateLatestAnnouncedVersion(announcedVersionStorage, announcedVersionKeyName(model), newSnapshotFileName)

	return
}

func MarshalDataModels(data []sources.DataModel) ([]byte, error) {
	universalDto := &UniversalDTO{Data: data}
	return json.Marshal(universalDto)
}

func UnMarshalDataModelsBytes(data []byte, model sources.DataModel) ([]sources.DataModel, error) {

	defer util.Duration(time.Now(), fmt.Sprintf("UnmarshalDataModelBytes for : %s", model.GetDataName()))
	oldData := &UniversalDTO{}

	p := time.Now()
	err := json.Unmarshal(data, oldData)
	logging.GetLogger().Info("Unmarshalling time : ", time.Since(p))
	if err != nil {
		logging.GetLogger().Info("Error in unmarshalling old data bytes : ", err)
		return nil, err
	}

	dataInterface, ok := (oldData.Data).([]interface{})
	if !ok {
		logging.GetLogger().Error("Error in typecasting the oldData into interface array, Err :", err)
		return nil, errors.New("error in typecasting old data bytes")
	}

	return UnMarshalInterfaceToModel(dataInterface, model)
}

func UnMarshalInterfaceToModel(dataInterface []interface{}, model sources.DataModel) ([]sources.DataModel, error) {

	models := make([]sources.DataModel, 0)
	for i := 0; i < len(dataInterface); i++ {
		dataMap, ok := dataInterface[i].(map[string]interface{})

		if !ok {
			logging.GetLogger().Error("Error in typecasting datampa")
		}

		data, _ := json.Marshal(dataMap)
		var result interface{}
		result = model.NewDataRef()
		err := json.Unmarshal(data, &result)

		if err != nil {
			logging.GetLogger().Error("err in unmarhsl ", err)
		}
		models = append(models, result.(sources.DataModel))
	}

	return models, nil
}

func announcedVersionKeyName(model sources.DataModel) string {
	return model.GetNameSpace() + separator + model.GetDataName()
}

// prevVersion = -1 means its being produced for the first time
func getAnnouncedVersionName(model sources.DataModel, prevVersion int64) string {
	if prevVersion == -1 {
		return fmt.Sprintf("%s-%d", announcedVersionKeyName(model), snapshot.DefaultVersionNumber)
	}
	return fmt.Sprintf("%s-%d", announcedVersionKeyName(model), prevVersion+1)
}
