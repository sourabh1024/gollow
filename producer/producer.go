//Copyright 2018 Sourabh Suman ( https://github.com/sourabh1024 )
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Package producer provides all methods for producing a data model
package producer

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/gollow/core"
	"github.com/gollow/core/snapshot"
	"github.com/gollow/core/storage"
	"github.com/gollow/data"
	"github.com/gollow/logging"
	"github.com/gollow/sources"
	"github.com/gollow/util"

	"sync"
	"time"
)

const (
	currentData  = "current"
	previousData = "prevData"
)

//dataLoadResult is used to store the results of the data load
//data is the data being loaded , in case of error or no data it is nil
//err stores the error if any during loading the data
//state denotes the state for which data is being loaded , currently it is currentData and prevData
type dataLoadResult struct {
	data  sources.Bag
	err   error
	state string
}

//ProduceModel produces the data for given DataModel
//1. Load the LastAnnouncedSnapshotName
//2. Load the Current and PrevData in parallel
//3. If data is produced for the first time produce data and return
//4. Generate Diff from curr and prev data if prev data was present
//5. Store the current snapshot
func ProduceModel(model sources.DataModel) {

	defer util.Duration(time.Now(), fmt.Sprintf("ProduceModel for : %s", model.GetDataName()))

	logging.GetLogger().Info("Starting data producing for : %s ", model.GetDataName())

	lastAnnouncedSnapshot, err := snapshot.VersionImpl.GetVersion(model.GetDataName())
	if err != nil && err != data.ErrNoData {
		logging.GetLogger().Error("error in loading snapshot , err : %+v", err)
		return
	}

	logging.GetLogger().Info("Last announced snapshot for %s, is : %s", model.GetDataName(),
		lastAnnouncedSnapshot)

	currData, prevData, ok := loadPrevAndCurrData(model, lastAnnouncedSnapshot)

	if !ok {
		// error in loading data , abort
		return
	}

	currBytes, err := proto.Marshal(currData)
	// Error in marshalling current data snapshot cannot be produced
	if err != nil {
		logging.GetLogger().Error("Error in marshalling current data, err :+v", err)
		return
	}

	// default value
	prevVersion := int64(-1)
	if lastAnnouncedSnapshot == "" {
		err := storeCurrentSnapshot(model, prevVersion, currBytes)
		if err != nil {
			logging.GetLogger().Error("Error in saving current snapshot : %+v", err)
		}
		logging.GetLogger().Info("Snapshot produced for the first time for : %s", model.GetDataName())

		return
	}

	prevVersion, err = snapshot.VersionImpl.ParseVersionNumber(lastAnnouncedSnapshot)
	if err != nil {
		logging.GetLogger().Error("Error in parsing the version number from : %s, err : %+v", lastAnnouncedSnapshot, err)
	}

	// generate diff
	ok = generateDiffFromData(model, prevData, currData, prevVersion, err)
	if !ok {
		//diff was not produced
		return
	}

	//store the current data snapshot
	err = storeCurrentSnapshot(model, prevVersion, currBytes)
	if err != nil {
		logging.GetLogger().Error("Error in saving current snapshot : %+v", err)
	}
	return
}

// generateDiffFromData generates the diff from given data model , current and previous data bag abd version information
// returns a bool whether the diff was generated or not
// in case of errors or no change in data diff is not generated
func generateDiffFromData(model sources.DataModel, prevData sources.Bag, currData sources.Bag, prevVersion int64, err error) bool {
	logging.GetLogger().Info("Generating diff for : ", model.GetDataName())
	diffParams := &core.DiffParams{
		Model:       model,
		OldData:     prevData,
		NewData:     currData,
		PrevVersion: prevVersion,
		CurrVersion: prevVersion + 1,
	}
	ok, err := diffParams.GenerateNewDiff()
	if err != nil {
		logging.GetLogger().Error("error in producing diff , err : %+v", err)
		return false

	}
	if !ok {
		logging.GetLogger().Info("no diff to produce for : %s", model.GetDataName())
		return false
	}

	return true
}

// loadPrevAndCurrData loads the previous and current data
func loadPrevAndCurrData(model sources.DataModel, lastAnnouncedSnapshot string) (
	sources.Bag, sources.Bag, bool) {
	var wg sync.WaitGroup
	wg.Add(1)
	dataChan := make(chan dataLoadResult, 2)
	go loadCurrentData(model, &wg, dataChan)
	if lastAnnouncedSnapshot != "" {
		wg.Add(1)
		snapshotStorage, err := storage.NewStorage(lastAnnouncedSnapshot)
		if err != nil {
			// error loading previous data
			dataChan <- dataLoadResult{
				state: previousData,
				data:  nil,
				err:   fmt.Errorf("error in reading previous data : %+v", err),
			}
		} else {
			go loadPreviousData(snapshotStorage, model, &wg, dataChan)
		}
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

	if currDataErr != nil {
		logging.GetLogger().Error("error in loading the current data , err : %+v", currDataErr)
		return nil, nil, false
	}

	if prevDataErr != nil {
		logging.GetLogger().Error("error in loading the previous data , err : +v", prevDataErr)
		return currData, nil, false
	}
	return currData, prevData, true
}

// storeCurrentSnapshot saves the current data bytes
// and updates the announced version for the data
func storeCurrentSnapshot(model sources.DataModel, prevVersion int64, data []byte) error {
	newSnapshotFileName := getNewVersionName(model, prevVersion)
	store, err := storage.NewStorage(newSnapshotFileName)
	if err != nil {
		return err
	}
	_, err = store.Write(data)
	if err != nil {
		return err
	}
	return snapshot.VersionImpl.UpdateVersion(model.GetDataName(), newSnapshotFileName)
}

// loadCurrentData loads the current data using the model.LoadAll
// returns the result on the dataLoadResult channel
func loadCurrentData(model sources.DataModel, wg *sync.WaitGroup,
	response chan dataLoadResult) {
	defer wg.Done()
	defer util.Duration(time.Now(), "loadCurrentData")

	bytes, err := model.LoadAll()
	if err != nil {
		response <- dataLoadResult{
			state: currentData,
			data:  nil,
			err:   err,
		}
	}

	response <- dataLoadResult{
		state: currentData,
		data:  bytes,
		err:   nil,
	}
}

// loadPreviousData loads the previous data from the storage
// returns the result on the dataLoadResult channel
func loadPreviousData(storage storage.Storage, model sources.DataModel,
	wg *sync.WaitGroup, response chan dataLoadResult) {
	defer wg.Done()
	defer util.Duration(time.Now(), "loadPreviousData")
	prevBytes, err := storage.Read()
	if err != nil {
		response <- dataLoadResult{
			state: previousData,
			data:  nil,
			err:   fmt.Errorf("error in reading previous data : %+v", err),
		}
	}

	prevData := model.NewBag()
	err = proto.Unmarshal(prevBytes, prevData)

	if err != nil {
		response <- dataLoadResult{
			state: previousData,
			data:  nil,
			err:   err,
		}
	}

	response <- dataLoadResult{
		state: previousData,
		data:  prevData,
		err:   err,
	}
}

// getNewVersionName returns the newVersion number given the model and prevVersion
func getNewVersionName(model sources.DataModel, prevVersion int64) string {
	if prevVersion == -1 {
		return fmt.Sprintf("%s-%d", model.GetDataName(), core.DefaultVersionNumber)
	}
	return fmt.Sprintf("%s-%d", model.GetDataName(), prevVersion+1)
}
