package logging

import (
	"context"
	"time"

	"api.default.indicoinnovation.pt/clients/google/logging"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
)

func Log(details *entity.LogDetails, severity string, resourceLabels *map[string]string) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.DefaultContextTimeout))
	defer cancel()

	go logging.Log(ctx, details, severity, resourceLabels)
}
