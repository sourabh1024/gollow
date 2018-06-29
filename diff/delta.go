package diff

import "gollow/sources"

type Delta struct {
	ChangedObjects map[string]interface{}
	NewObjects     map[string]interface{}
	MissingKeys    []string
}

func GetDelta(oldData []sources.DataModel, newData []sources.DataModel) interface{} {

	//newDataCount := len(newData)
	//oldDataCount := len(oldData)
	//
	//for i := 0; i < newDataCount; i++ {
	//	newKey := newData[i]
	//}
	return nil
}
