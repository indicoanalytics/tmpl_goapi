package response

import (
	"api.default.indicoinnovation.pt/adapters/logging"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

func CreateError(context *fiber.Ctx, err error, message string, response interface{}, statusCode ...int) error {
	httpCode := constants.HTTPStatusBadRequest
	if len(statusCode) > 0 {
		httpCode = statusCode[0]
	}

	go logging.Log(&entity.LogDetails{
		Message:    message,
		Reason:     err.Error(),
		Request:    helpers.FromHTTPRequest(context),
		Response:   response,
		StatusCode: httpCode,
	}, "error", nil)

	return helpers.CreateResponse(context, &entity.ErrorResponse{
		Message:     message,
		Description: getDescription(err),
		StatusCode:  httpCode,
	}, httpCode)
}

func CreateSuccess(context *fiber.Ctx, message string, response interface{}, statusCode ...int) error {
	httpCode := constants.HTTPStatusOK
	if len(statusCode) > 0 {
		httpCode = statusCode[0]
	}

	go logging.Log(&entity.LogDetails{
		Message:    message,
		StatusCode: httpCode,
		Request:    helpers.FromHTTPRequest(context),
		Response:   response,
	}, "debug", nil)

	return helpers.CreateResponse(context, response, httpCode)
}

func getDescription(err error) string {
	if helpers.ContainsError(constants.MappedClientErrors, err) {
		return err.Error()
	}

	return ""
}
