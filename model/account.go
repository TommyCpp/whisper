package model

import (
	"database/sql"
	"github.com/tommycpp/Whisper/sqlconnection"
	"log"
)

type Account struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Account) StoreIntoDB(db *sqlconnection.SqlConnection) (sql.Result, error) {
	res, err := db.Exec("INSERT INTO user(`id`,username,`password`) VALUES (?,?,?)", a.Id, a.Username, a.Password)
	if err != nil {
		log.Print("Cannot create user")
		log.Print(err)
		return nil, err
	}
	return res, nil
}

func (a *Account) CheckIfValid(db *sqlconnection.SqlConnection) (bool, error) {
	res, err := db.Query("SELECT * FROM user WHERE username=? AND password=?", a.Username, a.Password)
	defer res.Close()
	if err != nil {
		log.Print("Error when connect to DB")
		return false, err
	}
	if res.Next() {
		return true, nil
	} else {
		return false, nil
	}
}
