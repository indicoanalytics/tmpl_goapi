package logging

import (
	"api.default.indicoinnovation.pt/clients/google/logging"
	"api.default.indicoinnovation.pt/entity"
)

func Log(details *entity.LogDetails, severity string, resourceLabels *map[string]string) {
	go logging.Log(details, severity, resourceLabels)
}
