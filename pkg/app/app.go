package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"api.default.indicoinnovation.pt/adapters/database"
	"api.default.indicoinnovation.pt/adapters/logging"
	"api.default.indicoinnovation.pt/app/appinstance"
	"api.default.indicoinnovation.pt/clients/iam"
	"api.default.indicoinnovation.pt/config"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func ApplicationInit() {
	configs := config.New()
	ctx := context.Background()

	if constants.Environment != constants.Test {
		iam.New()
	}

	appinstance.Data = &appinstance.Application{
		Config: configs,
		Server: fiber.New(fiber.Config{
			ServerHeader: "Death Star",
			ErrorHandler: customErrorHandler,
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
			Prefork:      constants.Prefork,
		}),
	}

	appinstance.Data.DB = database.Connect(ctx)
}

func Setup() {
	var err error
	if constants.UseTLS {
		err = appinstance.Data.Server.ListenTLS(fmt.Sprintf(":%s", constants.Port), "cert.pem", "key.pem")
	} else {
		err = appinstance.Data.Server.Listen(fmt.Sprintf(":%s", constants.Port))
	}

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func customErrorHandler(ctx *fiber.Ctx, err error) error {
	var code int = fiber.StatusInternalServerError
	var capturedError *fiber.Error
	message := "unknown error"

	if errors.As(err, &capturedError) {
		code = capturedError.Code
		if code == fiber.StatusNotFound {
			message = "route not found"
		}
	}

	var errorResponse *entity.ErrorResponse

	erro := json.Unmarshal([]byte(err.Error()), &errorResponse)
	if erro != nil {
		errorResponse = &entity.ErrorResponse{
			Message:     message,
			StatusCode:  code,
			Description: err.Error(),
		}
	}

	go logging.Log(
		&entity.LogDetails{
			Message:  message,
			Method:   ctx.Method(),
			Reason:   err.Error(),
			RemoteIP: ctx.IP(),
			Request: map[string]interface{}{
				"body":       string(ctx.BodyRaw()),
				"query":      ctx.Queries(),
				"url_params": ctx.Locals("url_params"),
			},
			StatusCode: code,
			URLpath:    ctx.Path(),
		},
		"critical",
		nil,
	)

	helpers.CreateResponse(ctx, errorResponse, code) //nolint: wrapcheck

	return nil
}

func Log(ctx *fiber.Ctx) error {
	logSeverity := ctx.Locals(constants.LogSeverityKey)

	payload := new(entity.LogDetails)
	bytedata, _ := helpers.Marshal(ctx.Locals(constants.LogDataKey))
	helpers.Unmarshal(bytedata, &payload) //nolint: errcheck

	if logSeverity == nil {
		logSeverity = "debug"
	}

	body := map[string]interface{}{}
	helpers.Unmarshal(ctx.BodyRaw(), &body) //nolint: errcheck

	request := map[string]interface{}{
		"body":       body,
		"query":      ctx.Queries(),
		"url_params": ctx.Locals("url_params"),
	}

	severity, _ := logSeverity.(string)

	logging.Log(&entity.LogDetails{
		Message:    payload.Message,
		StatusCode: payload.StatusCode,
		Reason:     payload.Reason,
		Response:   payload.Response,
		Request:    request,
		Method:     ctx.Method(),
		RemoteIP:   ctx.IP(),
		URLpath:    ctx.Path(),
	}, severity, nil)

	return nil
}
