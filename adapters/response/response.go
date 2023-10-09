package response

import (
	"fmt"

	"api.default.indicoinnovation.pt/adapters/logging"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

func CreateError(context *fiber.Ctx, err error, action string, response interface{}, statusCode ...int) error {
	httpCode := constants.HTTPStatusBadRequest
	if len(statusCode) > 0 {
		httpCode = statusCode[0]
	}

	go logging.Log(&entity.LogDetails{
		Message:    fmt.Sprintf("error to %s at %s", action, context.Path()),
		Reason:     err.Error(),
		Request:    helpers.FromHTTPRequest(context),
		Response:   response,
		StatusCode: httpCode,
	}, "error", nil)

	return helpers.CreateResponse(context, &entity.ErrorResponse{
		Message:     fmt.Sprintf("error in %s", action),
		Description: getDescription(err),
		StatusCode:  httpCode,
	}, httpCode)
}

func CreateSuccess(context *fiber.Ctx, action string, response interface{}, statusCode ...int) error {
	httpCode := constants.HTTPStatusOK
	if len(statusCode) > 0 {
		httpCode = statusCode[0]
	}

	go logging.Log(&entity.LogDetails{
		Message:    fmt.Sprintf("success to %s at %s", action, context.Path()),
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
