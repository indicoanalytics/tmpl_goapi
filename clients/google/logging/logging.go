package logging

import (
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	gcpLogging "github.com/indicoinnovation/gcp_logging_easycall"
)

func Log(message *entity.LogDetails, severity string, resourceLabels *map[string]string) {
	labels := map[string]string{"service": constants.MainServiceName}
	if resourceLabels != nil {
		for k, v := range *resourceLabels {
			labels[k] = v
		}
	}

	gcpLogging.Log(
		constants.GcpProjectID,
		constants.MainLoggerName,
		message,
		severity,
		"api",
		labels,
	)
}
