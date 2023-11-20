package database

import (
	"context"
	"errors"

	"api.default.indicoinnovation.pt/pkg/app"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type Database struct{}

func New() *Database {
	return &Database{}
}

func (db *Database) Query(query string, outputType interface{}, args ...interface{}) (interface{}, error) {
	err := pgxscan.Select(context.Background(), app.Inst.DB, outputType, query, args...)

	return outputType, err
}

func (db *Database) Exec(query string, args ...interface{}) error {
	_, err := app.Inst.DB.Exec(context.Background(), query, args...)

	return err
}

func (db *Database) QueryOne(query string, outputType interface{}, args ...interface{}) (interface{}, error) {
	err := pgxscan.Get(context.Background(), app.Inst.DB, outputType, query, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		err = nil
	}

	return outputType, err
}

func (db *Database) QueryCount(query string, args ...interface{}) (int, error) {
	type Count struct {
		Count int `json:"count"`
	}

	rows := &Count{}

	err := pgxscan.Get(context.Background(), app.Inst.DB, rows, query, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		err = nil
	}

	return rows.Count, err
}
