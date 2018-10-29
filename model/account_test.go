package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/tommycpp/Whisper/sqlconnection"
	"testing"
)

func openConnection() *sqlconnection.SqlConnection {
	db, err := sql.Open("mysql", "whisperAdmin:123456@tcp(127.0.0.1:3306)/whisper")
	if err != nil || db == nil {
		fmt.Println(err)
	}
	return sqlconnection.NewConnection(db)
}

func TestAccount_StoreIntoDB(t *testing.T) {
	//MUST open validate mysql connection
	connection := openConnection()
	//clear all data
	connection.Exec("DELETE FROM user WHERE 1=1")
	var account Account
	account.Username = "test"
	account.Id = 1
	account.Password = "1234"
	res, _ := account.StoreIntoDB(connection)
	rowAffected, _ := res.RowsAffected()
	assert.EqualValues(t, 1, rowAffected)
}

func TestAccount_CheckIfValid(t *testing.T) {
	connection := openConnection()
	connection.Exec("DELETE FROM user WHERE 1=1")
	// Add an account
	var account Account
	account.Username = "test"
	account.Id = 1
	account.Password = "1234"
	account.StoreIntoDB(connection)

	res, _ := account.CheckIfValid(connection)
	assert.True(t, res)
}
