package iam

import (
	"context"
	"errors"

	"api.default.indicoinnovation.pt/adapters/logging"
	"api.default.indicoinnovation.pt/entity"
	indicoserviceauth "github.com/INDICO-INNOVATION/indico_service_auth"
)

var errIAMConnection = errors.New("error to connect to iam client, this could happen due to system outages, or unsupported errors")

func New() (*indicoserviceauth.Client, context.Context) {
	client, context, err := indicoserviceauth.NewClient()
	if err != nil {
		// TODO: specify error details to logging library
		logging.Log(&entity.LogDetails{}, "critical", nil)

		panic(errIAMConnection)
	}

	if context.Err() != nil {
		// TODO: specify error details to logging library
		logging.Log(&entity.LogDetails{}, "critical", nil)

		panic(errIAMConnection)
	}

	return client, context
}
