package diff

import (
	"gollow/sources"
	"reflect"
)

type Diff struct {
	ChangedObjects interface{} `json:"changed_objects"`
	NewObjects     interface{} `json:"new_objects"`
	MissingKeys    []string    `json:"missing_keys"`
}

func GetNewDiffObj() *Diff {
	return &Diff{}
}

func (d *Diff) GetDiffBetweenModels(oldData []sources.DataModel, newData []sources.DataModel) {

	oldDataMap := GetMapFromDataModel(oldData)
	newDataMap := GetMapFromDataModel(newData)

	diff := GetNewDiffObj()

	newObjects := make([]sources.DataModel, 0)
	changedObjects := make([]sources.DataModel, 0)
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
			d.MissingKeys = append(diff.MissingKeys, key)
		}
	}

	d.NewObjects = newObjects
	d.ChangedObjects = changedObjects
}

//
//func (d *Diff) Marshal() (bytes []byte, err error) {
//
//	return json.Marshal(d)
//}
//
//func (d *Diff) Unmarshal(data []byte) error {
//	return json.Unmarshal(data, &d)
//}
//

func GetMapFromDataModel(data []sources.DataModel) map[string]sources.DataModel {

	var result = make(map[string]sources.DataModel)
	lenData := len(data)

	for i := 0; i < lenData; i++ {
		key := data[i].GetPrimaryKey()
		//TODO : Check for collisions and WARN users about it
		result[key] = data[i]
	}

	return result
}
