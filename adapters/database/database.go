package database

import (
	"api.default.indicoinnovation.pt/pkg/app"
)

func Query(query string, outputType interface{}, args ...interface{}) (interface{}, error) {
	err := app.Inst.DBInstance.Raw(query, args...).Scan(&outputType).Error

	return outputType, err
}

func QueryCount(query string, args ...interface{}) (int, error) {
	var count int

	err := app.Inst.DBInstance.Raw(query, args...).Scan(&count).Error

	return count, err
}
