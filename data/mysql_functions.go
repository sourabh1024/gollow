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
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mysql driver package
	"golang.org/x/net/context"
	"gollow/logging"
	"reflect"
	"sync"
	"sync/atomic"
)

const (
	tagSQLCol       = "sql-col"
	tagSQLPrecision = "sql-precision"
)

// columnsInfo
type columnsInfo struct {
	all          []string
	colMap       map[string]int
	fieldMap     map[string]string
	precisionMap map[string]string
}

var (
	// columnsInfoCache caches the columns info
	columnsInfoCache = &columnsInfoMap{data: make(map[reflect.Type]*columnsInfo)}
)

type columnsInfoMap struct {
	data map[reflect.Type]*columnsInfo
	lock sync.Mutex
}

func setDB(config *MysqlConfig, db *sql.DB) {
	config.DBCache.DB = db
}

func getDB(config *MysqlConfig) *sql.DB {
	return config.DBCache.DB
}

var getDatabase = func(config *MysqlConfig) (db *sql.DB, err error) {

	//TODO : defer panic recovery

	// connect to the db once
	connect := func() {
		db, err = createDBConn(config)
	}

	config.ConnectOnce.Do(connect)

	if err != nil {
		logging.GetLogger().Error("Create DB error : %+v", err)
		return nil, err
	}

	// check for existing connection
	db = getDB(config)
	if db == nil {
		return nil, errors.New("failed to connect to DB")
	}

	return db, nil
}

var createDBConn = func(config *MysqlConfig) (db *sql.DB, err error) {
	logging.GetLogger().Info("[DB] Opening DB connections ")
	db, err = sql.Open("mysql", config.Dsn)

	if err != nil {
		logging.GetLogger().Error("Failed to open database connection. Error : ", err)
		return
	}

	db.SetMaxIdleConns(config.MaxIdle)
	db.SetMaxOpenConns(config.MaxOpen)

	logging.GetLogger().Info("Opened db connection to %s", config.Dsn)
	// store connection for later reuse
	setDB(config, db)
	return
}

type mySQLConnection interface {

	// query rows
	QueryRows(ctx context.Context, config *MysqlConfig, query string, reference interface{}, field ...interface{}) ([]interface{}, error)

	InsertRows(ctx context.Context, config *MysqlConfig, query []string, field ...interface{}) error
}

type mySQLConnectionImpl struct {
}

// NewMySQLConnectionRef returns the mysql connection impl
func NewMySQLConnectionRef() *mySQLConnectionImpl {
	return &mySQLConnectionImpl{}
}

// QueryRows queries the row with passed params
func (mySqlConnection *mySQLConnectionImpl) QueryRows(ctx context.Context, config *MysqlConfig, query string, reference interface{}, args ...interface{}) ([]interface{}, error) {

	pendingRequests := atomic.AddInt64(&config.PendingCalls, 1)
	defer atomic.AddInt64(&config.PendingCalls, -1)
	logging.GetLogger().Info("Number of pending sql requests : %d", pendingRequests)

	db, err := getDatabase(config)

	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, query, args...)

	if err != nil {
		logging.GetLogger().Error("Error in executing query context ", err)
		return nil, err
	}

	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	var output []interface{}

	colsInfo := buildColumnsInfo(reflect.TypeOf(reference))
	numberOfFields := len(colsInfo.all)
	oneRow := make([]interface{}, numberOfFields)

	for rows.Next() {

		var result interface{}

		switch v := reference.(type) {
		case Entity:
			result = v.NewEntity()
		default:
			continue
		}

		outputStruct := reflect.ValueOf(result).Elem()

		for i := 0; i < numberOfFields; i++ {
			columnName := colsInfo.all[i]
			fieldID := colsInfo.colMap[columnName]
			oneRow[i] = outputStruct.Field(fieldID).Addr().Interface()
		}

		err := rows.Scan(oneRow...)

		if err != nil {
			logging.GetLogger().Warning("error in scanning rows :%v", err)
			return nil, err
		}

		output = append(output, result)
	}

	if len(output) == 0 {
		return nil, ErrNoData
	}
	return output, nil
}

func (mySQLConnection *mySQLConnectionImpl) InsertRow(ctx context.Context, config *MysqlConfig, query string, args []interface{}) error {

	pendingRequests := atomic.AddInt64(&config.PendingCalls, 1)
	defer atomic.AddInt64(&config.PendingCalls, -1)
	logging.GetLogger().Info("Number of pending sql requests : %d", pendingRequests)

	db, err := getDatabase(config)

	if err != nil {
		return err
	}

	logging.GetLogger().Info("query : %s", query)

	// prepare the query
	stmt, _ := db.Prepare(query)

	_, err = stmt.Exec(args...)

	return err
}

var buildColumnsInfo = func(typ reflect.Type) *columnsInfo {

	output := getColsInfo(typ)
	if output != nil {
		// found in cache return
		return output
	}

	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		panic(fmt.Errorf("dest must be pointer to struct; got %T", typ))
	}

	output = &columnsInfo{
		colMap:       make(map[string]int),
		fieldMap:     make(map[string]string),
		precisionMap: make(map[string]string),
	}

	elem := typ.Elem()
	totalFields := elem.NumField()

	for index := 0; index < totalFields; index++ {
		field := elem.Field(index)
		// extract any with sql-col
		tagDbColumn := field.Tag.Get(tagSQLCol)
		if tagDbColumn == "" {
			continue
		}

		tagDbColumn = "`" + tagDbColumn + "`"

		output.fieldMap[field.Name] = tagDbColumn

		tagPrecision := field.Tag.Get(tagSQLPrecision)
		if tagPrecision != "" && field.Type.Name() == "float64" {
			output.precisionMap[tagDbColumn] = tagPrecision
		}

		output.all = append(output.all, tagDbColumn)

		output.colMap[tagDbColumn] = index
	}

	// cache and return
	setColsInfo(typ, output)
	return output
}

func getColsInfo(typ reflect.Type) *columnsInfo {
	defer columnsInfoCache.lock.Unlock()
	columnsInfoCache.lock.Lock()
	return columnsInfoCache.data[typ]
}

func setColsInfo(typ reflect.Type, ci *columnsInfo) {
	defer columnsInfoCache.lock.Unlock()
	columnsInfoCache.lock.Lock()
	columnsInfoCache.data[typ] = ci
}
