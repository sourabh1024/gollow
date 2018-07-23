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
	"time"
)

const DiffSeparator = "-"

var (
	DiffObjectDao = &DiffObject{}
)

type DiffObject struct {
	Namespace      string      `json:"namespace"`
	EntityName     string      `json:"entity_name"`
	ChangedObjects interface{} `json:"changed_objects"`
	NewObjects     interface{} `json:"new_objects"`
	MissingKeys    []string    `json:"missing_keys"`
}

type Diff interface {
	GetDiffName(model sources.DataModel, prevVersion, currVersion int64) string
	CreateNewDiff(model sources.DataModel, prevData, currData []sources.DataModel, prevVersion, currVersion int64) error
	SaveDiff()
}

func (d *DiffObject) GetDiffName(model sources.DataModel, prevVersion, currVersion int64) string {
	return "diff" + DiffSeparator + model.GetNameSpace() + DiffSeparator + model.GetDataName() + DiffSeparator + strconv.FormatInt(prevVersion, 10) + DiffSeparator + strconv.FormatInt(currVersion, 10)
}

func (d *DiffObject) CreateNewDiff(model sources.DataModel, prevData, currData []sources.DataModel, prevVersion, currVersion int64) (bool, error) {
	diffName := d.GetDiffName(model, prevVersion, currVersion)
	logging.GetLogger().Info("DiffObject name produced : ", diffName)
	diffStorage := storage.NewStorage(diffName)
	return d.createDiff(model, prevData, currData, prevVersion, currVersion, diffStorage)
}

func (d *DiffObject) createDiff(model sources.DataModel, prevData, currData []sources.DataModel,
	prevVersion, currVersion int64, store storage.Storage) (bool, error) {

	defer util.Duration(time.Now(), "CreateNewDiff")
	delta := getDiffBetweenModels(prevData, currData)

	if !shouldDiffBeProduced(delta) {
		return false, nil
	}

	delta.Namespace = model.GetNameSpace()
	delta.EntityName = model.GetDataName()

	err := d.SaveDiff(store, delta, model, prevVersion, currVersion)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (d *DiffObject) SaveDiff(store storage.Storage, delta *DiffObject, model sources.DataModel, prevVersion, currVersion int64) error {
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

func getDiffBetweenModels(oldData []sources.DataModel, newData []sources.DataModel) *DiffObject {

	defer util.Duration(time.Now(), fmt.Sprintf("GetDiffBetweenModels for : %s", newData[0].GetDataName()))

	diff := &DiffObject{}
	oldDataMap := getMapFromDataModel(oldData)
	newDataMap := getMapFromDataModel(newData)

	newObjects := make([]sources.DataModel, 0)
	changedObjects := make([]sources.DataModel, 0)
	missingKeys := make([]string, 0)
	for key, newValue := range newDataMap {

		oldValue, ok := oldDataMap[key]

		// key missing from oldData  => its a new key
		if !ok {
			newObjects = append(newObjects, newValue)
			continue
		}

		// data has changed for given key
		if !reflect.DeepEqual(oldValue, newValue) {
			changedObjects = append(changedObjects, newValue)
		}
	}

	for key, _ := range oldDataMap {
		if _, ok := newDataMap[key]; !ok {
			missingKeys = append(missingKeys, key)
		}
	}

	diff.NewObjects = newObjects
	diff.ChangedObjects = changedObjects
	diff.MissingKeys = missingKeys
	return diff
}

func getMapFromDataModel(data []sources.DataModel) map[string]sources.DataModel {

	defer util.Duration(time.Now(), fmt.Sprintf("getMapFromDataModel for len : %d", len(data)))
	var result = make(map[string]sources.DataModel)
	lenData := len(data)

	for i := 0; i < lenData; i++ {
		key := data[i].GetPrimaryKey()
		//TODO : Check for collisions and WARN users about it
		result[key] = data[i]
	}

	return result
}

func shouldDiffBeProduced(diff *DiffObject) bool {
	if diff == nil ||
		(reflect.DeepEqual(diff, &DiffObject{
			NewObjects:     make([]sources.DataModel, 0),
			ChangedObjects: make([]sources.DataModel, 0),
			MissingKeys:    make([]string, 0),
		})) {
		return false
	}
	return true
}

func marshalDiff(delta *DiffObject) ([]byte, error) {
	return json.Marshal(delta)
}
