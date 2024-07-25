package bigquery

import (
	"context"

	"api.default.indicoinnovation.pt/adapters/logging"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"cloud.google.com/go/bigquery"
)

func Insert(dataset, table string, data interface{}) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, constants.GcpProjectID)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message: "error to create new bigquery client",
			Reason:  err.Error(),
		}, "critical", nil)

		return err
	}

	inserter := client.Dataset(dataset).Table(table).Inserter()
	if err := inserter.Put(ctx, data); err != nil {
		logging.Log(&entity.LogDetails{
			Message: "error to send to bigquery",
			Reason:  err.Error(),
		}, "critical", nil)

		return err
	}

	return nil
}
