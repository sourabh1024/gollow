package data

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mysql driver package
	"golang.org/x/net/context"
	"gollow/cdd/logging"
	"reflect"
	"sync"
	"sync/atomic"
)

const (
	tagSQLCol       = "sql-col"
	tagSQLPrecision = "sql-precision"
	tagSQLKey       = "sql-key"

	// pre-set ID column names
	tagID          = "id"
	tagOriginID    = "originId"
	tagDestID      = "destId"
	tagIDAndDestID = "id_destId"
)

// columnsInfo defines a custom data type "list" of database columns
type columnsInfo struct {
	insertable   []string
	updatable    []string
	all          []string
	id           string
	idIsInt      bool
	originID     string
	destID       string
	colMap       map[string]int
	fieldMap     map[string]string
	precisionMap map[string]string
}

var (
	// columnsInfo cache
	columnsCache = &columnsInfoMap{data: make(map[reflect.Type]*columnsInfo)}
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
	logging.GetLogger().Info("Number of pending requests : %d", pendingRequests)

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

	output := []interface{}{}

	colsInfo := buildCoulumnsList(reflect.TypeOf(reference))
	numberOfFields := len(colsInfo.all)
	oneRow := make([]interface{}, numberOfFields)

	//data, err := gosqljson.QueryDbToMap(db, theCase, query)
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
			logging.GetLogger().Warning("load relations failed to scan row", err)
			return nil, err
		}

		output = append(output, result)
	}

	if len(output) == 0 {
		return nil, ErrNoData
	}
	return output, nil
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

		tagDbKey := field.Tag.Get(tagSQLKey)
		switch tagDbKey {
		case tagID:
			output.id = tagDbColumn
			output.idIsInt = field.Type.Kind() >= reflect.Int && field.Type.Kind() <= reflect.Uint64
		case tagOriginID:
			output.originID = tagDbColumn
		case tagDestID:
			output.destID = tagDbColumn
		case tagIDAndDestID:
			output.destID = tagDbColumn
			output.id = tagDbColumn
		}
		output.colMap[tagDbColumn] = index
	}

	// cache and return
	setColsInfo(typ, output)
	return output
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
