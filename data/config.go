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
