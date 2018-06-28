package diff

type Delta struct {
	ChangedObjects map[string]interface{}
	NewObjects     map[string]interface{}
	MissingKeys    []string
}

func GetDelta(oldData map[string]interface{}, newData map[string]interface{}) {

	//newDataCount := len(newData)
	//oldDataCount := len(oldData)
	//
	//for i := 0; i < newDataCount; i++ {
	//	newKey := newData[i]
	//}
}
