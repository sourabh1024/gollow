package producer

import (
	"encoding/json"
	"gollow/diff"
	"gollow/logging"
	"gollow/snapshot"
	"gollow/sources"
	"gollow/util"
	"gollow/write"
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

/**
Producer produces a given data source DataModel
@param model: DataModel to produce
@returns :
*/
func Producer(io write.SnapshotReaderWriter, announcedVersionPath string, model sources.DataModel) {

	/*
	 1. Get the previous announced version of the data.
	 2. De-Serialize and create back the old object
	 3. Load the current data from the data source
	 4. Get the diff of the data.
	 5. Serialize and create the new snapshot
	 6. Create the diff
	 7. Update the announced version
	*/

	logging.GetLogger().Info("Starting data producing for : ", model.GetDataName())

	dbData, err := model.LoadAll()
	newData := dbData.([]sources.DataModel)

	if err != nil {
		logging.GetLogger().Error("Error in loading data from data source, Err: ", err)
		return
	}

	newDataBytes, err := json.Marshal(newData)

	if err != nil {
		logging.GetLogger().Error("Error in generating marshalled data from data source, Err: ", err)
		return
	}

	lastAnnouncedVersion, err := snapshot.GetLatestAnnouncedVersion(io, announcedVersionPath, model.GetNameSpace()+"-"+model.GetDataName())

	if err != nil {
		logging.GetLogger().Error("Error in getting previous announced version : ", err)
		return
	}

	if lastAnnouncedVersion == "" {
		// Generate new Snapshot and update and go away
		//TODO: improve this
		io.Write("/Users/sourabh.suman/gopath/src/gollow/snapshots"+"/"+model.GetNameSpace()+"-"+model.GetDataName()+"-"+util.GetCurrentTimeString(), newDataBytes)
		return
	}

	bytes, err := io.Read(lastAnnouncedVersion)

	if err != nil {
		logging.GetLogger().Error("Error in Reading last announced version snapshot: ", err)
		return
	}

	var oldData []sources.DataModel
	err = json.Unmarshal(bytes, oldData)

	if err != nil {
		logging.GetLogger().Error("Error in unmarshalling previous announced version : ", err)
		return
	}

	delta := diff.GetDelta(oldData, newData)

	if delta != nil {
		// produce delta
	}
}
