package producer

import (
	"fmt"
	"gollow/core"
	"gollow/core/snapshot"
	"gollow/core/storage"
	"gollow/logging"
	"gollow/sources"
	"gollow/util"
	"time"
)

const (
	separator = "-"
)

type UniversalDTO struct {
	Data interface{} `json:"data"`
}

/**
ProduceModel produces a given data source DataModel
@param model: DataModel to produce
@returns :
*/
func ProduceModel(model sources.DataModel) {

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

	defer util.Duration(time.Now(), fmt.Sprintf("ProduceModel for : %s", model.GetDataName()))

	logging.GetLogger().Info("Starting data producing for : ", model.GetDataName())

	data, err := model.LoadAll()
	if err != nil {
		logging.GetLogger().Error("Error in fetching current data  : ", err)
	}

	currData := data

	currBytes, err := util.MarshalDataModels(currData)
	if err != nil {
		logging.GetLogger().Error("Error in marshalling current data bytes : ", err)
	}

	lastAnnouncedSnapshot, err := snapshot.SnapshotImpl.GetLatestAnnouncedVersion(snapshot.AnnouncedVersionKeyName(model.GetNameSpace(),
		model.GetDataName()))
	if err != nil {
		logging.GetLogger().Error("Error in getting previous announced version : ", err)
		return
	}

	if lastAnnouncedSnapshot == "" {
		newSnapshotFileName := getAnnouncedVersionName(model, -1)
		snapshot.WriteNewSnapshot(newSnapshotFileName, currBytes)
		snapshot.SnapshotImpl.UpdateLatestAnnouncedVersion(snapshot.AnnouncedVersionKeyName(model.GetNameSpace(),
			model.GetDataName()), newSnapshotFileName)
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

	prevData, err := util.UnMarshalDataModelsBytes(prevBytes, model)

	logging.GetLogger().Info("Generating diff for : ", model.GetDataName())
	logging.GetLogger().Info("DiffObject prevVersion , currVersion : ", prevSnapshotVersion, prevSnapshotVersion+1)
	ok, err := core.DiffObjectDao.CreateNewDiff(model, prevData, currData, prevSnapshotVersion, prevSnapshotVersion+1)

	if err != nil || !ok {
		logging.GetLogger().Error("Error in producing diff for : "+newSnapshotName, err)
		return
	}

	newSnapshotFileName := getAnnouncedVersionName(model, prevSnapshotVersion)

	snapshot.SnapshotImpl.UpdateLatestAnnouncedVersion(snapshot.AnnouncedVersionKeyName(model.GetNameSpace(),
		model.GetDataName()), newSnapshotFileName)

	return
}

// prevVersion = -1 means its being produced for the first time
func getAnnouncedVersionName(model sources.DataModel, prevVersion int64) string {
	if prevVersion == -1 {
		return fmt.Sprintf("%s-%d", snapshot.AnnouncedVersionKeyName(model.GetNameSpace(),
			model.GetDataName()), snapshot.DefaultVersionNumber)
	}
	return fmt.Sprintf("%s-%d", snapshot.AnnouncedVersionKeyName(model.GetNameSpace(),
		model.GetDataName()), prevVersion+1)
}
