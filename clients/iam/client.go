package iam

import (
	"context"
	"errors"

	indicoserviceauth "github.com/INDICO-INNOVATION/indico_service_auth"
)

var ErrIAMConnection = errors.New("error to connect to iam client, this could happen due to system outages, or unsupported errors")

func New() (*indicoserviceauth.Client, context.Context) {
	client, context, err := indicoserviceauth.NewClient()
	if err != nil {
		// TODO: Log client error
		panic(ErrIAMConnection)
	}

	if context.Err() != nil {
		panic(ErrIAMConnection)
	}

	return client, context
}
