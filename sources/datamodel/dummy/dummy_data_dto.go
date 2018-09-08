//Copyright 2018 Sourabh Suman ( https://github.com/sourabh1024 )
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package dummy

import (
	"context"
	"strconv"
	"time"
	"github.com/gollow/data"
	"github.com/gollow/sources"
	"github.com/gollow/util"
	"github.com/gollow/logging"
	"github.com/gollow/config"
	"github.com/gollow/util/profile"
)

//var _ sources.DataModel = &DataDTO{}
var _ data.Entity = &DataDTO{}

// DummyDataRef is the reference object for DummyData Data
var DummyDataRef = &DataDTO{}

// DataDTO is the struct required by DTO to load the message from data source
// it can vary depending on the data source , should be defined by user
type DataDTO struct {
	ID        int64     `sql-col:"id" primary-key:"true" json:"id"`
	PID       int64     `sql-col:"pid" json:"pid"`
	FirstName string    `sql-col:"first_name" json:"firstname"`
	LastName  string    `sql-col:"last_name" json:"lastname"`
	Balance   float64   `sql-col:"balance" json:"balance"`
	MaxCredit float64   `sql-col:"max_credit" json:"max_credit"`
	MaxDebit  float64   `sql-col:"max_debit" json:"max_debit"`
	Score     float64   `sql-col:"score" json:"score"`
	IsActive  bool      `sql-col:"is_active" json:"is_shown"`
	CreatedAt time.Time `sql-col:"created_at" json:"created_at"`
}

// ToPB implements the DTO interface
// ToPB converts the DTO object to protoBuf Message type
// required for data production
func (d DataDTO) ToPB() sources.Message {

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
func (d *DataDTO) NewEntity() data.Entity {
	return &DataDTO{}
}

//GetPrimaryKey implements the data.Entity interface
func (d *DataDTO) GetPrimaryKey() string {
	return strconv.FormatInt(d.ID, 10)
}

//GetDataName implements the DataModel interface
func (d *DataDTO) GetDataName() string {
	return "dummy_data"
}

//LoadAll implements the DataModel interface
func (d *DataDTO) LoadAll() (sources.Bag, error) {

	defer util.Duration(time.Now(), "DummyDataDataFetchFromSQL")
	logging.GetLogger().Info("Starting loading of data for %s", d.GetDataName())

	var result = &DummyDataBag{}

	query := "SELECT * FROM dummy_data"
	entities, err := data.NewMySQLConnectionRef().QueryRows(context.Background(), config.GlobalConfig.MySQLConfig, query, &DataDTO{})

	if err != nil {
		logging.GetLogger().Error("Error in fetching data from DB : ", err)
		return result, err
	}

	defer util.Duration(time.Now(), "Converting Data")

	lenResult := len(entities)

	for i := 0; i < lenResult; i++ {
		entity, ok := entities[i].(*DataDTO)

		if !ok {
			logging.GetLogger().Error("Error in typecasting the results , err: %+v", err)
			continue
		}

		result.AddEntry(entity.ToPB())
	}

	logging.GetLogger().Info("Length of result returned from DB : %d ", lenResult)

	profile.GetMemoryProfile()
	return result, nil
}
