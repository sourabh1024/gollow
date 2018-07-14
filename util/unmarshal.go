package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"gollow/logging"
	"gollow/sources"
	"reflect"
	"sync"
)

const tag_json_fields = "json"

var errNilPtr = errors.New("destination pointer is nil") // embedded in descriptive error

// columnsInfo defines a custom data type "list" of database columns
type columnsInfo struct {
	all      []string
	colMap   map[string]int
	fieldMap map[string]string
}

func getColsInfo(typ reflect.Type) *columnsInfo {
	defer columnsCache.lock.Unlock()
	columnsCache.lock.Lock()
	return columnsCache.data[typ]
}

func setColsInfo(typ reflect.Type, ci *columnsInfo) {
	defer columnsCache.lock.Unlock()
	columnsCache.lock.Lock()
	columnsCache.data[typ] = ci
}

var (
	// columnsInfo cache
	columnsCache = &columnsInfoMap{data: make(map[reflect.Type]*columnsInfo)}
)

type columnsInfoMap struct {
	data map[reflect.Type]*columnsInfo
	lock sync.Mutex
}

var buildCoulumnsList = func(typ reflect.Type) *columnsInfo {

	output := getColsInfo(typ)
	if output != nil {
		// found in cache return
		return output
	}

	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		panic(fmt.Errorf("dest must be pointer to struct; got %T", typ))
	}

	output = &columnsInfo{
		colMap:   make(map[string]int),
		fieldMap: make(map[string]string),
	}

	elem := typ.Elem()
	totalFields := elem.NumField()

	for index := 0; index < totalFields; index++ {
		field := elem.Field(index)
		//extract any with json

		tagMarshalledFields := field.Tag.Get(tag_json_fields)
		if tagMarshalledFields == "" {
			continue
		}

		tagMarshalledFields = "`" + tagMarshalledFields + "`"

		output.fieldMap[field.Name] = tag_json_fields

		output.all = append(output.all, tagMarshalledFields)

		output.colMap[tagMarshalledFields] = index
	}

	// cache and return
	setColsInfo(typ, output)
	return output
}

type UniversalDTO struct {
	Data interface{} `json:"data"`
}

func UnmarshalDataModelBytesFast(data []byte, reference interface{}) ([]sources.DataModel, error) {
	//defer Duration(time.Now(), fmt.Sprintf("UnmarshalDataModelBytesFast for : %s", model.GetDataName()))
	oldData := &UniversalDTO{}
	err := json.Unmarshal(data, oldData)
	if err != nil {
		logging.GetLogger().Info("Error in unmarshalling old data bytes : ", err)
		return nil, err
	}

	dataInterface, ok := (oldData.Data).([]interface{})
	if !ok {
		logging.GetLogger().Error("Error in typecasting the oldData into interface array, Err :", err)
		return nil, errors.New("error in typecasting old data bytes")
	}

	output := []interface{}{}

	fieldsInfo := buildCoulumnsList(reflect.TypeOf(reference))
	numberOfFields := len(fieldsInfo.all)
	oneRow := make([]interface{}, numberOfFields)

	for i := 0; i < len(dataInterface); i++ {

		var result interface{}

		switch v := reference.(type) {
		case sources.DataModel:
			result = v.NewDataRef()
		default:
			continue
		}

		outputStruct := reflect.ValueOf(result).Elem()

		dataMap, ok := dataInterface[i].(map[string]interface{})

		if !ok {
			logging.GetLogger().Error("Error in typecasting the value to map[string]interface{}")
		}

		for i := 0; i < numberOfFields; i++ {
			columnName := fieldsInfo.all[i]
			fieldID := fieldsInfo.colMap[columnName]
			oneRow[i] = outputStruct.Field(fieldID).Addr().Interface()
		}
		//
		//for i:=0; i< numberOfFields; i++ {
		//
		//	switch x:= oneRow[i].(type) {
		//	case blake2b.XOF =:
		//
		//	}
		//}
		id := dataMap["id"].(int)
		fmt.Println(id)

		if err != nil {
			logging.GetLogger().Warning("load relations failed to scan row", err)
			return nil, err
		}

		output = append(output, result)

	}

	return nil, nil
}
