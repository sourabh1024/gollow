package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"gollow/logging"
	"gollow/sources"
	"gollow/storage"
	"gollow/util"
	"reflect"
	"strconv"
	"time"
)

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
	return "diff-" + model.GetNameSpace() + "-" + model.GetDataName() + "-" + strconv.FormatInt(prevVersion, 10) + "-" + strconv.FormatInt(currVersion, 10)
}

func (d *DiffObject) CreateNewDiff(model sources.DataModel, prevData, currData []sources.DataModel, prevVersion, currVersion int64) error {

	defer util.Duration(time.Now(), "CreateNewDiff")
	delta := getDiffBetweenModels(prevData, currData)

	if !shouldDiffBeProduced(delta) {
		return errors.New("no new data for diff")
	}

	delta.Namespace = model.GetNameSpace()
	delta.EntityName = model.GetDataName()

	return d.SaveDiff(delta, model, prevVersion, currVersion)
}

func (d *DiffObject) SaveDiff(delta *DiffObject, model sources.DataModel, prevVersion, currVersion int64) error {
	diffBytes, err := marshalDiff(delta)
	if err != nil {
		logging.GetLogger().Error("Error in marshalling diff : ", diffBytes)
		return err
	}

	diffName := d.GetDiffName(model, prevVersion, currVersion)
	logging.GetLogger().Info("DiffObject name produced : ", diffName)
	diffManager := storage.NewStorage(diffName)

	_, err = diffManager.Write(diffBytes)
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
