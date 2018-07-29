package core

import (
	"encoding/json"
	"fmt"
	"gollow/core/storage"
	"gollow/logging"
	"gollow/sources"
	"gollow/util"
	"reflect"
	"strconv"
	"sync"
	"time"
)

const DiffSeparator = "-"

var (
	DiffObjectDao = &DiffObject{}
)

type DiffObject struct {
	Namespace      string      `json:"namespace"`
	EntityName     string      `json:"entity_name"`
	ChangedObjects sources.Bag `json:"changed_objects"`
	NewObjects     sources.Bag `json:"new_objects"`
	MissingKeys    []string    `json:"missing_keys"`
}

type Diff interface {
	GetDiffName(model sources.ProtoDataModel, prevVersion, currVersion int64) string
	CreateNewDiff(model sources.ProtoDataModel, prevData, currData sources.Bag, prevVersion, currVersion int64) error
}

func (d *DiffObject) GetDiffName(model sources.ProtoDataModel, prevVersion, currVersion int64) string {
	return "diff" + DiffSeparator + model.GetDataName() + DiffSeparator + strconv.FormatInt(prevVersion, 10) + DiffSeparator + strconv.FormatInt(currVersion, 10)
}

func (d *DiffObject) CreateNewDiff(model sources.ProtoDataModel, prevData, currData sources.Bag, prevVersion, currVersion int64) (bool, error) {
	diffName := d.GetDiffName(model, prevVersion, currVersion)
	logging.GetLogger().Info("DiffObject name produced : ", diffName)
	diffStorage := storage.NewStorage(diffName)
	return d.createDiff(model, prevData, currData, prevVersion, currVersion, diffStorage)
}

func (d *DiffObject) createDiff(model sources.ProtoDataModel, prevData, currData sources.Bag,
	prevVersion, currVersion int64, store storage.Storage) (bool, error) {

	defer util.Duration(time.Now(), "CreateNewDiff")
	delta := getDiffBetweenModels(prevData, currData, model)

	if !shouldDiffBeProduced(delta) {
		return false, nil
	}

	delta.EntityName = model.GetDataName()

	err := d.SaveDiff(store, delta, model, prevVersion, currVersion)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (d *DiffObject) SaveDiff(store storage.Storage, delta *DiffObject, model sources.ProtoDataModel, prevVersion, currVersion int64) error {
	diffBytes, err := marshalDiff(delta)
	if err != nil {
		logging.GetLogger().Error("Error in marshalling diff : ", diffBytes)
		return err
	}

	_, err = store.Write(diffBytes)
	if err != nil {
		logging.GetLogger().Error("Error in writing diff : ", err)
	}
	return err
}

func getDiffBetweenModels(oldData sources.Bag, newData sources.Bag, model sources.ProtoDataModel) *DiffObject {

	defer util.Duration(time.Now(), fmt.Sprintf("GetDiffBetweenModels for : %s", model.GetDataName()))

	diff := &DiffObject{}

	//parllelize it
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

func (d *DiffObject) findNewOrChangedKeys(oldDataMap, newDataMap map[string]sources.Message, model sources.ProtoDataModel) {
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

func (d *DiffObject) findMissingKeys(oldDataMap, newDataMap map[string]sources.Message) {
	missingKeys := make([]string, 0)
	for key, _ := range oldDataMap {
		if _, ok := newDataMap[key]; !ok {
			missingKeys = append(missingKeys, key)
		}
	}
	d.MissingKeys = missingKeys
}

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
func getMapFromDataModel(d sources.Bag, result map[string]sources.Message) {

	defer util.Duration(time.Now(), fmt.Sprintf("getMapFromDataModel for len : %d"))
	data := d.GetEntries()
	lenData := len(data)

	for i := 0; i < lenData; i++ {
		key := data[i].GetPrimaryID()
		result[key] = data[i]
	}
}

func shouldDiffBeProduced(diff *DiffObject) bool {
	if diff == nil || (len(diff.ChangedObjects.GetEntries()) == 0 &&
		len(diff.NewObjects.GetEntries()) == 0 &&
		len(diff.MissingKeys) == 0) {
		return false
	}
	return true
}

func marshalDiff(delta *DiffObject) ([]byte, error) {
	return json.Marshal(delta)
}
