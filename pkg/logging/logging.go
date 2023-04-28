package logging

import (
	"api.default.indicoinnovation.pt/pkg/constants"
	gcpLogging "github.com/INDICO-INNOVATION/gcp_logging_easycall"
)

type LogDetails struct {
	User         string      `json:"token"`
	Message      string      `json:"message"`
	Reason       string      `json:"reason"`
	RemoteIP     string      `json:"ipaddress"`
	Method       string      `json:"method"`
	URLpath      string      `json:"route"`
	StatusCode   int         `json:"status_code"`
	RequestData  interface{} `json:"request_data"`
	ResponseData interface{} `json:"response_data"`
	SessionID    string      `json:"sessid"`
}

func Log(message *LogDetails, severity string, resourceLabels *map[string]string) {
	logMessage := &gcpLogging.Logger{
		User:         message.User,
		Message:      message.Message,
		Reason:       message.Reason,
		RemoteIp:     message.RemoteIP,
		Method:       message.Method,
		Urlpath:      message.URLpath,
		StatusCode:   message.StatusCode,
		RequestData:  message.RequestData,
		ResponseData: message.ResponseData,
		SessionId:    message.SessionID,
	}

	labels := map[string]string{"service": constants.MainServiceName}
	if resourceLabels != nil {
		for k, v := range *resourceLabels {
			labels[k] = v
		}
	}

	gcpLogging.Log(
		constants.GcpProjectID,
		constants.MainLoggerName,
		logMessage,
		severity,
		"api",
		labels,
	)
}
