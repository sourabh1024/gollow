package data

import (
	"database/sql"
	"encoding/json"
	"gollow/cdd/logging"
	"io/ioutil"
	"sync"
)

// MysqlConfig struct for initialising mysql based on client specific config.
type MysqlConfig struct {
	Dsn          string    `json:"dsn"`
	MaxIdle      int       `json:"maxIdle"`
	MaxOpen      int       `json:"maxOpen"`
	hostnameOnce sync.Once // we use this field to generate config.Name once
	stringOnce   sync.Once // we use this field to stringify once
	PendingCalls int64     `json:"-"`
	ConnectOnce  sync.Once `json:"-"` // we use this field to guarantee that we only connect to DB once
	stringValue  string
	DBCache      struct {
		*sql.DB
	} `json:"-"`
}

func (config *MysqlConfig) Init() error {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		logging.GetLogger().Error("Error in reading config file :+v", err)
		return err
	}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		logging.GetLogger().Error("Error in reading config file :+v", err)
		return err
	}
	return nil
}
