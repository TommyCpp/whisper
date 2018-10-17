package sqlconnection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

func NewConnection(db *sql.DB) *SqlConnection {
	return &SqlConnection{
		db: db,
	}
}

func GetSqlConnection(configuration *config.Configuration) *SqlConnection {
	if sqlConnection != nil {
		return sqlConnection // fixme: not thread-safe
	} else {
		db, err := sql.Open(configuration.DatabaseDriveName, configuration.DatabaseURLName)
		if err != nil || db == nil {
			log.Fatal("Error when open DB")
			return nil
		} else {
			if err := db.Ping(); err != nil {
				log.Fatal(err)
				log.Fatal("Error when connecting to DB")
				return nil
			}
			sqlConnection = &SqlConnection{
				db: db,
			}
			return sqlConnection
		}
	}
}

func (s *SqlConnection) Exec(query string, paras ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, paras...)
}
func (s *SqlConnection) Query(query string, para ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, para...)
}
