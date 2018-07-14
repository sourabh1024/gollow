package diff

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

type Diff struct {
	Namespace      string      `json:"namespace"`
	EntityName     string      `json:"entity_name"`
	ChangedObjects interface{} `json:"changed_objects"`
	NewObjects     interface{} `json:"new_objects"`
	MissingKeys    []string    `json:"missing_keys"`
}

func GetNewDiffObj(nameSpace, entityName string) *Diff {
	return &Diff{
		Namespace:  nameSpace,
		EntityName: entityName,
	}
}

func (d *Diff) GetDiffBetweenModels(oldData []sources.DataModel, newData []sources.DataModel) {

	defer util.Duration(time.Now(), fmt.Sprintf("GetDiffBetweenModels for : %s", newData[0].GetDataName()))

	oldDataMap := GetMapFromDataModel(oldData)
	newDataMap := GetMapFromDataModel(newData)

	//diff := GetNewDiffObj(newData[0].GetNameSpace(), newData[0].GetDataName())

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

	d.NewObjects = newObjects
	d.ChangedObjects = changedObjects
	d.MissingKeys = missingKeys
}

func GetMapFromDataModel(data []sources.DataModel) map[string]sources.DataModel {

	defer util.Duration(time.Now(), fmt.Sprintf("GetMapFromDataModel for len : %d", len(data)))
	var result = make(map[string]sources.DataModel)
	lenData := len(data)

	for i := 0; i < lenData; i++ {
		key := data[i].GetPrimaryKey()
		//TODO : Check for collisions and WARN users about it
		result[key] = data[i]
	}

	return result
}

func CreateNewDiff(model sources.DataModel, prevData, currData []sources.DataModel, prevVersion, currVersion int64) error {

	defer util.Duration(time.Now(), "CreateNewDiff")
	delta := GetNewDiffObj(model.GetNameSpace(), model.GetDataName())
	delta.GetDiffBetweenModels(prevData, currData)

	if !shouldDiffBeProduced(delta) {
		return errors.New("no new data for diff")
	}

	diffBytes, err := MarshalDiff(delta)
	if err != nil {
		logging.GetLogger().Error("Error in marshalling diff : ", diffBytes)
		return err
	}

	diffName := GetDiffName(model, prevVersion, currVersion)
	logging.GetLogger().Info("Diff name produced : ", diffName)
	logging.GetLogger().Info(model.GetDataName())
	logging.GetLogger().Info("version : ", prevVersion)
	diffManager := storage.NewFileStorage(diffName)

	_, err = diffManager.Write(diffBytes)

	return err
}

func shouldDiffBeProduced(diff *Diff) bool {
	if diff == nil ||
		(reflect.TypeOf(diff.NewObjects).Size() == 0 &&
			reflect.TypeOf(diff.ChangedObjects).Size() == 0 &&
			len(diff.MissingKeys) == 0) {
		return false
	}
	return true
}

func GetDiffName(model sources.DataModel, prevVersion, currVersion int64) string {
	return "diff-" + model.GetNameSpace() + "-" + model.GetDataName() + "-" + strconv.FormatInt(prevVersion, 10) + "-" + strconv.FormatInt(currVersion, 10)
}

func MarshalDiff(delta *Diff) ([]byte, error) {
	return json.Marshal(delta)
}

//func UnMarshalDiffBytes(data []byte) (*Diff, error) {
//	d := GetNewDiffObj()
//	err := json.Unmarshal(data, &d)
//	return d, err
//}
