package datamodel

import (
	"context"
	"gollow/config"
	"gollow/data"
	"gollow/logging"
	"gollow/sources"
	"gollow/util"
	"gollow/util/profile"
	"reflect"
	"strconv"
	"time"
)

//HeatMapDataRef is the reference object for HeatMap Data
var HeatMapDataRef = &HeatMapData{}

//HeatMapData is the DataModelName being loaded
type HeatMapData struct {
	ID             int64     `sql-col:"id" sql-key:"id" sql-insert:"false" primary-key:"true" json:"id"`
	VehicleTypeID  int64     `sql-col:"vehicle_type_id" json:"vehicle_type_id"`
	Start          time.Time `sql-col:"start_time" json:"start"`
	End            time.Time `sql-col:"end_time" json:"end"`
	Geohash        string    `sql-col:"geohash" json:"geohash"`
	UnmetDemand    int       `sql-col:"unmet_demand" json:"unmet_demand"`
	TotalDemand    int       `sql-col:"total_demand" json:"total_demand"`
	Imbalance      int       `sql-col:"imbalance" json:"imbalance"`
	Surge          float64   `sql-col:"surge" json:"surge"`
	Score          float64   `sql-col:"score" json:"score"`
	Sources        string    `sql-col:"sources" json:"sources"`
	IsOverSupplied bool      `sql-col:"is_over_supplied" json:"is_over_supplied"`
	IsShown        bool      `sql-col:"is_shown" json:"is_shown"`
	Version        string    `sql-col:"version" json:"version"`
}

//////////////////////////////////////////////////////////////////
///////////// Implement data.Entity interface ////////////////////
//////////////////////////////////////////////////////////////////

// NewEntity implements the data.Entity interface.
func (hd HeatMapData) NewEntity() data.Entity {
	return &HeatMapData{}
}

//GetNameSpace implements the data.Entity interface
func (hd HeatMapData) GetNameSpace() string {
	return "test-consumer-group1"
}

//GetPrimaryKey implements the data.Entity interface
func (hd HeatMapData) GetPrimaryKey() string {
	return strconv.FormatInt(hd.ID, 10)
}

//GetDataName implements the data.Entity interface
func (hd HeatMapData) GetDataName() string {
	return "heatmap_data"
}

//////////////////////////////////////////////////////////////////
///////////// End of Implement data.Entity interface /////////////
//////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////
///////////// Implement producer.DataProducer interface /////////////
//////////////////////////////////////////////////////////////////

func (hd HeatMapData) NewDataRef() sources.DataModel {
	return &HeatMapData{}
}

func (hd *HeatMapData) CacheDuration() int64 {
	return int64(time.Duration(5 * time.Second))
}

func (hd *HeatMapData) LoadAll() (interface{}, error) {

	defer util.Duration(time.Now(), "HeatMapDataFetchFromSQL")
	logging.GetLogger().Info("Starting load of data")

	//TODO : Remove this from here
	config.Init()

	var result []sources.DataModel

	query := "SELECT * FROM heatmap_data"
	entities, err := data.NewMySQLConnectionRef().NativeQueryRows(context.Background(), config.MySQLConfig, query, &HeatMapData{})

	if err != nil {
		logging.GetLogger().Error("Error in fetching data from DB : ", err)
		return result, err
	}

	defer util.Duration(time.Now(), "HeatMapDataConverting")
	logging.GetLogger().Info("Starting converting of data")

	lenResult := len(entities)
	logging.GetLogger().Info("Length of data returned from DB : %d ", lenResult)
	logging.GetLogger().Info("Size of data returned  from DB : %d ", reflect.TypeOf(entities).Size())

	result = make([]sources.DataModel, lenResult)

	for i := 0; i < lenResult; i++ {
		entity, ok := entities[i].(*HeatMapData)

		if !ok {
			logging.GetLogger().Error("Error in typecasting the results , err: ", err)
			continue
		}

		result[i] = entity
	}

	logging.GetLogger().Info("Length of result returned from DB : %d ", lenResult)
	logging.GetLogger().Info("Size of result returned from DB : %d ", reflect.TypeOf(result).Size())
	logging.GetLogger().Info("Type of result returned from DB : %d ", reflect.TypeOf(result))

	profile.GetMemoryProfile()
	return result, nil
}

//func (hd *HeatMapData) MarshalJSON() ([]byte, error) {
//	if hd.ID%10000 == 0 {
//		logging.GetLogger().Info("Marshalling : ", hd.GetPrimaryKey())
//	}
//	return json.Marshal(hd)
//}
//
//func (hd *HeatMapData) UnmarshalJSON(data []byte) error {
//	err := json.Unmarshal(data, &hd)
//	if err != nil {
//		logging.GetLogger().Error(" Error in unmarshalling  : ", err)
//	}
//	if hd.ID%10000 == 0 {
//		logging.GetLogger().Info("UnMarshalling : ", hd.GetPrimaryKey())
//	}
//	return nil
//}

//////////////////////////////////////////////////////////////////
///////////// End of Implement data.Producer interface /////////////
//////////////////////////////////////////////////////////////////
