package postgres

import (
	"database/sql"
	"log"

	// Import postgres.
	_ "github.com/lib/pq"
)

func Connect(dbString string) *sql.DB {
	const maxIdleConnection = 5
	const maxOpenConnection = 30
	const connectionMaxLifetime = 60

	pool, err := sql.Open("postgres", dbString)
	if err != nil {
		panic(err)
	}

	pool.SetMaxIdleConns(maxIdleConnection)
	pool.SetMaxOpenConns(maxOpenConnection)
	pool.SetConnMaxLifetime(connectionMaxLifetime)

	if err := pool.Ping(); err != nil {
		panic(err)
	}

	log.Println("postgres database connected successfully")

	return pool
}
