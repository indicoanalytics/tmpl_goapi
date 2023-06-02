package database

import (
	"database/sql"
	"errors"
	"log"

	"api.default.indicoinnovation.pt/clients/postgres"
	"api.default.indicoinnovation.pt/pkg/app"
)

var errConnectDB = errors.New("error to connect to database")

func Query(query string, outputType interface{}, args ...interface{}) (interface{}, error) {
	conn := Connect(app.Inst.Config.DBString, app.Inst.Config.DBLogMode, app.Inst.Config.Debug)
	err := conn.QueryRow(query, args...).Scan(&outputType)
	conn.Close()

	return outputType, err
}

func Exec(query string, args ...interface{}) error {
	conn := Connect(app.Inst.Config.DBString, app.Inst.Config.DBLogMode, app.Inst.Config.Debug)
	err := conn.QueryRow(query, args...).Err()
	conn.Close()

	return err
}

func QueryCount(query string, args ...interface{}) (int, error) {
	var count int

	conn := Connect(app.Inst.Config.DBString, app.Inst.Config.DBLogMode, app.Inst.Config.Debug)
	err := conn.QueryRow(query, args...).Scan(&count)
	conn.Close()

	return count, err
}

func Connect(dbString string, logLevel int, debug bool) *sql.DB {
	databaseConnection, err := postgres.Connect(dbString, logLevel, debug)
	if err != nil {
		log.Panicln(errConnectDB, err)
	}

	log.Printf("Database is now connected")

	return databaseConnection
}
