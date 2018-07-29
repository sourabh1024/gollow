package producer

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"gollow/cdd/core"
	"gollow/cdd/core/snapshot"
	"gollow/cdd/core/storage"
	"gollow/cdd/logging"
	"gollow/cdd/sources"
	"gollow/cdd/util"
	"sync"
	"time"
)

const (
	currentData  = "current"
	previousData = "prevData"
)

type result struct {
	data  sources.Bag
	err   error
	state string
}

func ProduceModel(model sources.ProtoDataModel) {

	defer util.Duration(time.Now(), fmt.Sprintf("ProduceModel for : %s", model.GetDataName()))

	logging.GetLogger().Info("Starting data producing for : ", model.GetDataName())

	lastAnnouncedSnapshot, err := snapshot.SnapshotImpl.GetLatestAnnouncedVersion(model.GetDataName())
	if err != nil {
		logging.GetLogger().Error("error in loading snapshot , err : %+v", err)
		return
	}

	logging.GetLogger().Info("Last announced snapshot for %s, %s", model.GetDataName(), lastAnnouncedSnapshot)

	var wg sync.WaitGroup
	wg.Add(1)

	dataChan := make(chan result, 2)
	go loadCurrentData(model, &wg, dataChan)

	if lastAnnouncedSnapshot != "" {
		wg.Add(1)
		snapshotStorage := storage.NewStorage(lastAnnouncedSnapshot)
		go loadPreviousData(snapshotStorage, model, &wg, dataChan)
	}
	wg.Wait()
	close(dataChan)

	var currData, prevData sources.Bag
	var currDataErr, prevDataErr error

	for val := range dataChan {
		switch val.state {
		case currentData:
			currData = val.data
			currDataErr = val.err

		case previousData:
			prevData = val.data
			prevDataErr = val.err
		}
	}

	logging.GetLogger().Info("producer received all data")

	if currDataErr != nil {
		logging.GetLogger().Error("error in loading the current data :, %+v", currDataErr)
		return
	}

	//TODO : delegate it to the model interface
	currBytes, err := proto.Marshal(currData)
	// save current snapshot
	if err != nil {
		logging.GetLogger().Error("error in marshalling current data, err :+v", err)
		return
	}

	// default value
	prevVersion := int64(-1)
	if lastAnnouncedSnapshot == "" {
		err = storeCurrentSnapshot(model, prevVersion, currBytes)
		if err != nil {
			logging.GetLogger().Error("error in saving current snapshot : %+v", err)
		}
		logging.GetLogger().Info("diff produced for the first time for : ", model.GetDataName())
		return
	}

	prevVersion = snapshot.GetVersionNumber(lastAnnouncedSnapshot)
	if prevDataErr != nil {
		logging.GetLogger().Error("error in loading the previous data , could not proceed to produce diff, err :+v", err)
		return
	}

	// generate diff
	logging.GetLogger().Info("Generating diff for : ", model.GetDataName())
	ok, err := core.DiffObjectDao.CreateNewDiff(model, prevData, currData, prevVersion, prevVersion+1)
	if err != nil {
		logging.GetLogger().Error("error in producing diff , err : %+v", err)
		return
	}

	if !ok {
		logging.GetLogger().Info("no diff to produce for : %s", model.GetDataName())
		return
	}

	prevVersion = snapshot.GetVersionNumber(lastAnnouncedSnapshot)
	err = storeCurrentSnapshot(model, prevVersion, currBytes)
	if err != nil {
		logging.GetLogger().Error("error in saving current snapshot : %+v", err)
	}
	return
}

//func storeCurrentSnapshot(model sources.DataModel, prevVersion int64, data []byte) error {
func storeCurrentSnapshot(model sources.ProtoDataModel, prevVersion int64, data []byte) error {
	newSnapshotFileName := getAnnouncedVersionName(model, prevVersion)
	err := snapshot.WriteNewSnapshot(newSnapshotFileName, data)
	if err != nil {
		return err
	}
	return snapshot.SnapshotImpl.UpdateLatestAnnouncedVersion(model.GetDataName(), newSnapshotFileName)
}

func loadCurrentData(model sources.ProtoDataModel, wg *sync.WaitGroup, response chan result) {
	defer wg.Done()
	defer util.Duration(time.Now(), "loadCurrentData")

	data, err := model.LoadAll()
	if err != nil {
		response <- result{
			state: currentData,
			data:  nil,
			err:   err,
		}
	}

	response <- result{
		state: currentData,
		data:  data,
		err:   nil,
	}
}

func loadPreviousData(storage storage.Storage, model sources.ProtoDataModel, wg *sync.WaitGroup, response chan result) {
	defer wg.Done()
	defer util.Duration(time.Now(), "loadPreviousData")
	prevBytes, err := storage.Read()
	if err != nil {
		response <- result{
			state: previousData,
			data:  nil,
			err:   errors.New(fmt.Sprintf("Error in reading previous data : %+v", err)),
		}
	}

	prevData := model.NewBag()
	err = proto.Unmarshal(prevBytes, prevData)

	if err != nil {
		response <- result{
			state: previousData,
			data:  nil,
			err:   err,
		}
	}

	response <- result{
		state: previousData,
		data:  prevData,
		err:   err,
	}
}

func getAnnouncedVersionName(model sources.ProtoDataModel, prevVersion int64) string {
	if prevVersion == -1 {
		return fmt.Sprintf("%s-%d", model.GetDataName(), snapshot.DefaultVersionNumber)
	}
	return fmt.Sprintf("%s-%d", model.GetDataName(), prevVersion+1)
}
