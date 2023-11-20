package database

import (
	"context"
	"errors"
	"reflect"

	"api.default.indicoinnovation.pt/pkg/app"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type (
	Output             interface{}
	Database[T Output] struct{}
)

func New[T Output]() *Database[T] {
	return &Database[T]{}
}

func Query[T Output](query string, outputType T, args ...interface{}) (T, error) { //nolint: ireturn
	if reflect.TypeOf(outputType).Elem().Kind() == reflect.Slice {
		return New[T]().Query(query, outputType, args...)
	}
	return New[T]().QueryOne(query, outputType, args...)
}

func Exec(query string, args ...interface{}) error {
	return New[Output]().Exec(query, args...)
}

func QueryOne[T Output](query string, outputType T, args ...interface{}) (T, error) { //nolint: ireturn
	return New[T]().QueryOne(query, outputType, args...)
}

func QueryCount(query string, args ...interface{}) (int, error) {
	return New[int]().QueryCount(query, args...)
}

func (db *Database[T]) Query(query string, outputType T, args ...interface{}) (T, error) { //nolint: ireturn
	err := pgxscan.Select(context.Background(), app.Inst.DB, outputType, query, args...)

	return outputType, err
}

func (db *Database[T]) Exec(query string, args ...interface{}) error {
	_, err := app.Inst.DB.Exec(context.Background(), query, args...)

	return err
}

func (db *Database[T]) QueryOne(query string, outputType T, args ...interface{}) (T, error) { //nolint: ireturn
	err := pgxscan.Get(context.Background(), app.Inst.DB, outputType, query, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		err = nil
	}

	return outputType, err
}

func (db *Database[T]) QueryCount(query string, args ...interface{}) (int, error) {
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
