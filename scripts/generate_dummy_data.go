package main

import (
	"context"
	"fmt"
	"gollow/config"
	"gollow/data"
	"gollow/logging"
	"gollow/sources/datamodel/dummy"
	"strings"
)

func insertIntoDummyData(numOfRows int) {
	datas := make([]*dummy.DataDTO, 0, numOfRows)
	for i := 0; i < 2; i++ {
		dummyData := &dummy.DataDTO{ID: int64(i), FirstName: "suman", LastName: "sourabh", Balance: 10.0, MaxCredit: 5.0, MaxDebit: 3.0, Score: 1.0, IsActive: true}
		datas = append(datas, dummyData)
	}
	err := BulkInsert(datas)
	if err != nil {
		logging.GetLogger().Error("error in inserting into mysql with err : %v", err)
	}
}
func main() {
	insertIntoDummyData(2)
}

func BulkInsert(unsavedRows []*dummy.DataDTO) error {
	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0)
	i := 0

	for _, post := range unsavedRows {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, interface{}(post.PID))
		valueArgs = append(valueArgs, interface{}(post.FirstName))
		valueArgs = append(valueArgs, interface{}(post.LastName))
		valueArgs = append(valueArgs, interface{}(post.Balance))
		valueArgs = append(valueArgs, interface{}(post.MaxCredit))
		valueArgs = append(valueArgs, interface{}(post.MaxDebit))
		valueArgs = append(valueArgs, interface{}(post.Score))
		valueArgs = append(valueArgs, interface{}(post.IsActive))
		valueStrings = append(valueStrings, fmt.Sprintf("(%d , '%s', '%s', %f, %f, %f, %f, %t)", post.PID, post.FirstName, post.LastName, post.Balance, post.MaxCredit, post.MaxDebit, post.Score, post.IsActive))
		i++
	}
	stmt := "INSERT INTO dummy_data (pid, first_name, last_name, balance, max_credit, max_debit, score, is_active) VALUES"
	stmt = stmt + strings.Join(valueStrings, ",")
	err := data.NewMySQLConnectionRef().InsertRow(context.Background(), config.GlobalConfig.MySQLConfig, stmt, valueArgs)

	return err
}
