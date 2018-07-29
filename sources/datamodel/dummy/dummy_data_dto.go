package dummy

import (
	"context"
	"gollow/config"
	"gollow/data"
	"gollow/logging"
	"gollow/sources"
	"gollow/util"
	"gollow/util/profile"
	"strconv"
	"time"
)

//DummyDataRef is the reference object for DummyData Data
var DummyDataRef = &DummyDataDTO{}

//DummyData is the DataModelName being loaded
type DummyDataDTO struct {
	ID        int64     `sql-col:"id" sql-key:"id" sql-insert:"false" primary-key:"true" json:"id"`
	PID       int64     `sql-col:"pid" json:"pid"`
	FirstName string    `sql-col:"first_name" json:"firstname"`
	LastName  string    `sql-col:"last_name" json:"lastname"`
	Balance   float64   `sql-col:"balance" json:"balance"`
	MaxCredit float64   `sql-col:"max_credit" json:"max_credit"`
	MaxDebit  float64   `sql-col:"max_debit" json:"max_debit"`
	Score     float64   `sql-col:"score" json:"score"`
	IsActive  bool      `sql-col:"is_shown" json:"is_shown"`
	CreatedAt time.Time `sql-col:"created_at" json:"created_at"`
}

func (d DummyDataDTO) ToPB() sources.Message {

	return &DummyData{
		ID:         d.ID,
		PID:        d.PID,
		FirstName:  d.FirstName,
		LastName:   d.LastName,
		Balance:    d.Balance,
		MaxCredit:  d.MaxCredit,
		MaxDebit:   d.MaxDebit,
		Score:      d.Score,
		IsActive:   d.IsActive,
		Created_At: int64(d.CreatedAt.UnixNano()),
	}
}

// NewEntity implements the data.Entity interface.
func (d *DummyDataDTO) NewDataRef() sources.Bag {
	return &DummyDataBag{}
}

// NewEntity implements the data.Entity interface.
func (d *DummyDataDTO) NewEntity() data.Entity {
	return &DummyDataDTO{}
}

//GetNameSpace implements the data.Entity interface
func (d *DummyDataDTO) GetNameSpace() string {
	return "test-consumer-group1"
}

//GetPrimaryKey implements the data.Entity interface
func (d *DummyDataDTO) GetPrimaryKey() string {
	return strconv.FormatInt(d.ID, 10)
}

//GetDataName implements the data.Entity interface
func (d *DummyDataDTO) GetDataName() string {
	return "dummy_data"
}

func (d *DummyDataDTO) CacheDuration() int64 {
	return int64(time.Duration(2 * time.Minute))
}

func (d *DummyDataDTO) LoadAll() (sources.Bag, error) {

	defer util.Duration(time.Now(), "DummyDataDataFetchFromSQL")
	logging.GetLogger().Info("Starting loading of data for %s", d.GetDataName())

	var result = &DummyDataBag{}

	query := "SELECT * FROM dummy_data_large_5"
	entities, err := data.NewMySQLConnectionRef().NativeQueryRows(context.Background(), config.GlobalConfig.MySQLConfig, query, &DummyDataDTO{})

	if err != nil {
		logging.GetLogger().Error("Error in fetching data from DB : ", err)
		return result, err
	}

	defer util.Duration(time.Now(), "DummyData converting")

	lenResult := len(entities)

	for i := 0; i < lenResult; i++ {
		entity, ok := entities[i].(*DummyDataDTO)

		if !ok {
			logging.GetLogger().Error("Error in typecasting the results , err: ", err)
			continue
		}

		result.AddEntry(entity.ToPB())
	}

	logging.GetLogger().Info("Length of result returned from DB : %d ", lenResult)

	profile.GetMemoryProfile()
	return result, nil
}
