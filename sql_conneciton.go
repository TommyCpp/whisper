package main

import (
	"database/sql"
	"github.com/tommycpp/Whisper/config"
	"log"
)

var sqlConnection *SqlConnection

type SqlConnection struct {
	db *sql.DB
}

func (sql *SqlConnection) Close() {
	defer sql.db.Close()
}

func GetSqlConnection() *SqlConnection {
	if sqlConnection != nil {
		return sqlConnection // fixme: not thread-safe
	} else {
		db, err := sql.Open(config.Config.DatabaseDriveName, config.Config.DatabaseDriveName)
		if err != nil {
			log.Fatal("Error when open DB")
		} else {
			sqlConnection = &SqlConnection{
				db: db,
			}
			return sqlConnection
		}
	}

}
