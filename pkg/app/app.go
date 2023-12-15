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

func customErrorHandler(context *fiber.Ctx, err error) error {
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
			Message:    message,
			Reason:     err.Error(),
			StatusCode: code,
			Request:    helpers.FromHTTPRequest(context),
		},
		"critical",
		nil,
	)

	return helpers.CreateResponse(context, errorResponse, code) //nolint: wrapcheck
}

func Log(ctx *fiber.Ctx) error {
	logMessage := ctx.Locals("log_message")
	errorReason := ctx.Locals("error_reason")
	response := ctx.Locals("response")
	logSeverity := ctx.Locals("log_severity")
	statusCode := ctx.Locals("status_code")

	if statusCode == nil {
		statusCode = constants.HTTPStatusOK
	}

	if logSeverity == nil {
		logSeverity = "info"
	}

	severity, _ := logSeverity.(string)
	reason, _ := errorReason.(string)
	message, _ := logMessage.(string)
	status, _ := statusCode.(int)

	logging.Log(&entity.LogDetails{
		Message:    message,
		StatusCode: status,
		Reason:     reason,
		Response:   response,
		Request:    string(ctx.BodyRaw()),
		Method:     ctx.Method(),
	}, severity, nil)

	return nil
}
