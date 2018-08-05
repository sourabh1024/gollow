package data

import (
	"database/sql"
	"sync"
)

// MysqlConfig struct for initialising mysql settings
type MysqlConfig struct {
	Dsn          string    `json:"dsn"`
	MaxIdle      int       `json:"maxIdle"`
	MaxOpen      int       `json:"maxOpen"`
	hostnameOnce sync.Once // generate config.Name once
	PendingCalls int64     `json:"-"`
	ConnectOnce  sync.Once `json:"-"` //  connect to DB once
	stringValue  string
	DBCache      struct {
		*sql.DB
	} `json:"-"`
}
