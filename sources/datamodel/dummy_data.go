package datamodel

//
//import (
//	"gollow/util/profile"
//	"reflect"
//	"gollow/logging"
//)
//
////DummyDataRef is the reference object for DummyData Data
//var DummyDataRef = &DummyData{}
//
////DummyData is the DataModelName being loaded
//type DummyData struct {
//	ID        int64     `sql-col:"id" sql-key:"id" sql-insert:"false" primary-key:"true" json:"id"`
//	PID       int64     `sql-col:"pid" json:"pid"`
//	FirstName string    `sql-col:"first_name" json:"firstname"`
//	LastName  string    `sql-col:"last_name" json:"lastname"`
//	Balance   float64   `sql-col:"balance" json:"balance"`
//	MaxCredit float64   `sql-col:"max_credit" json:"max_credit"`
//	MaxDebit  float64   `sql-col:"max_debit" json:"max_debit"`
//	Score     float64   `sql-col:"score" json:"score"`
//	IsActive  bool      `sql-col:"is_shown" json:"is_shown"`
//	CreatedAt time.Time `sql-col:"created_at" json:"created_at"`
//}
//
////////////////////////////////////////////////////////////////////
/////////////// Implement data.Entity interface ////////////////////
////////////////////////////////////////////////////////////////////
//
//// NewEntity implements the data.Entity interface.
//func (d *DummyData) NewEntity() data.Entity {
//	return &DummyData{}
//}
//
////GetNameSpace implements the data.Entity interface
//func (d *DummyData) GetNameSpace() string {
//	return "test-consumer-group1"
//}
//
////GetPrimaryKey implements the data.Entity interface
//func (d *DummyData) GetPrimaryKey() string {
//	return strconv.FormatInt(d.ID, 10)
//}
//
////GetDataName implements the data.Entity interface
//func (d *DummyData) GetDataName() string {
//	return "dummy_data"
//}
//
////////////////////////////////////////////////////////////////////
/////////////// End of Implement data.Entity interface /////////////
////////////////////////////////////////////////////////////////////
//
////////////////////////////////////////////////////////////////////
/////////////// Implement producer.DataProducer interface /////////////
////////////////////////////////////////////////////////////////////
//
//func (d *DummyData) NewDataRef() sources.DataModel {
//	return &DummyData{}
//}
//
//func (d *DummyData) CacheDuration() int64 {
//	return int64(time.Duration(2 * time.Minute))
//}
//
//func (d *DummyData) LoadAll() ([]sources.DataModel, error) {
//
//	defer util.Duration(time.Now(), "DummyDataDataFetchFromSQL")
//	logging.GetLogger().Info("Starting loading of data for %s", d.GetDataName())
//
//	var result []sources.DataModel
//
//	query := "SELECT * FROM dummy_data_large_2"
//	entities, err := data.NewMySQLConnectionRef().NativeQueryRows(context.Background(), config.GlobalConfig.MySQLConfig, query, &DummyData{})
//
//	if err != nil {
//		logging.GetLogger().Error("Error in fetching data from DB : ", err)
//		return result, err
//	}
//
//	defer util.Duration(time.Now(), "DummyData converting")
//
//	lenResult := len(entities)
//
//	result = make([]sources.DataModel, lenResult)
//
//	for i := 0; i < lenResult; i++ {
//		entity, ok := entities[i].(*DummyData)
//
//		if !ok {
//			logging.GetLogger().Error("Error in typecasting the results , err: ", err)
//			continue
//		}
//
//		result[i] = entity
//	}
//
//	logging.GetLogger().Info("Length of result returned from DB : %d ", lenResult)
//	logging.GetLogger().Info("Size of result returned from DB : %d ", reflect.TypeOf(result).Size())
//	logging.GetLogger().Info("Type of result returned from DB : %d ", reflect.TypeOf(result))
//
//	profile.GetMemoryProfile()
//	return result, nil
//}
//
//////////////////////////////////////////////////////////////////////
/////////////// End of Implement data.Producer interface /////////////
////////////////////////////////////////////////////////////////////
