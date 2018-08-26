package core

import (
	"encoding/json"
	"fmt"
	"reflect"

	"strconv"
	"sync"
	"time"

	"gollow/core/storage"
	"gollow/sources"
	"gollow/util"
)

//GenerateDiff interface to generate and save the diff
type GenerateDiff interface {

	//GenerateNewDiff generates the diff and returns whether diff was produced and error if any
	//requires all parameters of DiffParams
	GenerateNewDiff() (bool, error)

	//GetDiffName returns the name of the diff produced based on dataname and prevVersion and currVersion
	GetDiffName() string

	//LoadDiff loads the given diff name and returns the diffObject using the NewStorage
	LoadDiff(diffName string) (*DiffObject, error)
}

//DiffParams represents the diff params required for generating the diff
type DiffParams struct {
	Model       sources.DataModel
	OldData     sources.Bag
	NewData     sources.Bag
	PrevVersion int64
	CurrVersion int64
}

//DiffObject represents the DiffObject its the final diff which is produced
//DataName represents the DataName for which data was produced
//ChangedObjects represents the Bag of Changed object
//NewObjects represents the Bag of New Objects
//MissingKeys represents the keys being deleted
type DiffObject struct {
	DataName       string      `json:"dataName"`
	ChangedObjects sources.Bag `json:"changedObjects"`
	NewObjects     sources.Bag `json:"newObjects"`
	MissingKeys    []string    `json:"missingKeys"`
}

//GetDiffName returns the name of the diff produced based on dataname and prevVersion and currVersion
func (params *DiffParams) GetDiffName() string {
	return DiffPrefix + Separator +
		params.Model.GetDataName() + Separator +
		strconv.FormatInt(params.PrevVersion, 10) + Separator +
		strconv.FormatInt(params.CurrVersion, 10)
}

// GenerateNewDiff implements the GenerateDiff interface
// returns whether the diff was produced or not and error
func (params *DiffParams) GenerateNewDiff() (bool, error) {
	defer util.Duration(time.Now(), "CreateNewDiff")
	diff := getDiffBetweenModels(params.OldData, params.NewData, params.Model)

	if !shouldDiffBeProduced(diff) {
		return false, nil
	}
	err := saveDiff(params, diff)
	if err != nil {
		return false, err
	}
	return true, nil
}

// LoadDiff loads the given diff name and returns the diffObject using the NewStorage
func (params *DiffParams) LoadDiff(diffName string) (*DiffObject, error) {
	store, err := storage.NewStorage(diffName)
	if err != nil {
		return nil, err //couldn't get storage
	}
	data, err := store.Read()
	if err != nil {
		return nil, err
	}
	d := &DiffObject{}
	d.ChangedObjects = params.Model.NewBag()
	d.NewObjects = params.Model.NewBag()
	err = json.Unmarshal(data, &d)
	return d, err
}

//saveDiff saves the diff object
func saveDiff(params *DiffParams, diff *DiffObject) error {
	store, err := storage.NewStorage(params.GetDiffName())
	if err != nil {
		return err // couldn't get storage
	}
	diffBytes, err := marshalDiff(diff)
	if err != nil {
		return err
	}
	_, err = store.Write(diffBytes)
	if err != nil {
		return err
	}
	return nil
}

// getDiffBetweenModels finds the diff between old and new data
func getDiffBetweenModels(oldData sources.Bag, newData sources.Bag, model sources.DataModel) *DiffObject {

	defer util.Duration(time.Now(), fmt.Sprintf("GetDiffBetweenModels for : %s", model.GetDataName()))

	diff := &DiffObject{}

	//fetch the map of primaryKey -> Data to generate the diff
	oldDataMap, newDataMap := getDataMaps(oldData, newData)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		diff.findNewOrChangedKeys(oldDataMap, newDataMap, model)
		wg.Done()
	}()

	go func() {
		diff.findMissingKeys(oldDataMap, newDataMap)
		wg.Done()
	}()

	wg.Wait()

	return diff
}

// findNewOrChangedKeys finds the object which are new or whose any value have been changed
func (d *DiffObject) findNewOrChangedKeys(oldDataMap, newDataMap map[string]sources.Message, model sources.DataModel) {

	d.NewObjects = model.NewBag()
	d.ChangedObjects = model.NewBag()

	for key, newValue := range newDataMap {

		oldValue, ok := oldDataMap[key]

		// key missing from oldData  => its a new key
		if !ok {
			d.NewObjects.AddEntry(newValue)
			continue
		}

		// data has changed for given key
		if !reflect.DeepEqual(oldValue, newValue) {
			d.ChangedObjects.AddEntry(newValue)
		}
	}
}

// findMissingKeys finds the missing keys from the newData map compared to oldData map
func (d *DiffObject) findMissingKeys(oldDataMap, newDataMap map[string]sources.Message) {
	var missingKeys []string
	for key := range oldDataMap {
		//key is missing from newDataMap means key has been deleted
		if _, ok := newDataMap[key]; !ok {
			missingKeys = append(missingKeys, key)
		}
	}
	d.MissingKeys = missingKeys
}

// getDataMaps returns the map for old and new Data
func getDataMaps(oldSource, newSource sources.Bag) (oldMap, newMap map[string]sources.Message) {
	oldMap = make(map[string]sources.Message, len(oldSource.GetEntries()))
	newMap = make(map[string]sources.Message, len(newSource.GetEntries()))

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		getMapFromDataModel(oldSource, oldMap)
		wg.Done()
	}()
	go func() {
		getMapFromDataModel(newSource, newMap)
		wg.Done()
	}()

	wg.Wait()
	return oldMap, newMap
}

// getMapFromDataModel converts bag into map of primaryId -> object
func getMapFromDataModel(d sources.Bag, result map[string]sources.Message) {

	defer util.Duration(time.Now(), fmt.Sprintf("getMapFromDataModel for len : %d", len(d.GetEntries())))
	data := d.GetEntries()
	lenData := len(data)

	for i := 0; i < lenData; i++ {
		key := data[i].GetUniqueKey()
		result[key] = data[i]
	}
}

// shouldDiffBeProduced checks whether the diff should be produced or not
// in future should it be exported to model level decision ?
// checks if the diff object is nil or,
// checks if there is no data to be produced in new, changed or missing objects
func shouldDiffBeProduced(diff *DiffObject) bool {
	if diff == nil || (len(diff.ChangedObjects.GetEntries()) == 0 &&
		len(diff.NewObjects.GetEntries()) == 0 &&
		len(diff.MissingKeys) == 0) {
		return false
	}
	return true
}

// marshalDiff marshals the diff
func marshalDiff(diff *DiffObject) ([]byte, error) {
	return json.Marshal(diff)
}
