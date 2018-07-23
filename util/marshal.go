package util

import (
	"encoding/json"
	"gollow/sources"
)

func MarshalDataModels(data []sources.DataModel) ([]byte, error) {
	universalDto := &UniversalDTO{Data: data}
	return json.Marshal(universalDto)
}
