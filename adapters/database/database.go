package database

import (
	"database/sql"
	"errors"
	"log"

	"api.default.indicoinnovation.pt/clients/postgres"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/pkg/app"
	"gorm.io/gorm"
)

type Database struct{}

func New() *Database {
	return &Database{}
}

var errConnectDB = errors.New("error to connect to database")

// Using generics.
func (db *Database) Query(query string, outputType interface{}, args ...interface{}) (interface{}, error) {
	gormConn, conn := connect(app.Inst.Config.DBString, app.Inst.Config.DBLogMode, constants.Debug)
	defer conn.Close()

	err := gormConn.Raw(query, args...).Scan(&outputType).Error

	return outputType, err
}

func (db *Database) Exec(query string, args ...interface{}) error {
	_, conn := connect(app.Inst.Config.DBString, app.Inst.Config.DBLogMode, constants.Debug)
	defer conn.Close()

	err := conn.QueryRow(query, args...).Err()

	return err
}

func (db *Database) QueryCount(query string, args ...interface{}) (int, error) {
	var count int

	_, conn := connect(app.Inst.Config.DBString, app.Inst.Config.DBLogMode, constants.Debug)
	defer conn.Close()

	err := conn.QueryRow(query, args...).Scan(&count)

	return count, err
}

func connect(dbString string, logLevel int, debug bool) (*gorm.DB, *sql.DB) {
	gormDB, databaseConnection, err := postgres.Connect(dbString, logLevel, debug)
	if err != nil {
		log.Panicln(errConnectDB, err)
	}

	return gormDB, databaseConnection
}
