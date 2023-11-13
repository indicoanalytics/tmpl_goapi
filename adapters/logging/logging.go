package logging

import (
	"strings"

	"api.default.indicoinnovation.pt/clients/google/logging"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
)

func Log(payload *entity.LogDetails, severity constants.LoggingSeverity, resourceLabels *map[string]string) {
	payload.Message = strings.ToLower(payload.Message)
	payload.Message = strings.TrimSpace(payload.Message)

	if payload.Context != nil {
		payload.Request = helpers.FromHTTPRequest(payload.Context)
	}

	logging.Log(payload, string(severity), resourceLabels)
}
