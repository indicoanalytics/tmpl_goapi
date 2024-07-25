package health

import (
	"errors"

	"api.default.indicoinnovation.pt/app/errs"
	"api.default.indicoinnovation.pt/app/usecases/health"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usecase *health.Usecase
}

func Handle() *Handler {
	return &Handler{
		usecase: health.New(),
	}
}

func (handler *Handler) Check(ctx *fiber.Ctx) error {
	check, err := handler.usecase.Check()
	if err != nil {
		ctx.Locals(constants.LogDataKey, &entity.LogDetails{
			Message:    "error to check health",
			Reason:     err.Error(),
			StatusCode: constants.HTTPStatusInternalServerError,
		})
		ctx.Locals(constants.LogSeverityKey, constants.SeverityError)

		var (
			description  string
			requestError *errs.RequestError
		)

		if errors.As(err, &requestError) {
			description = requestError.Unwrap().Error()
		}

		helpers.CreateResponse(ctx, &entity.ErrorResponse{
			Message:     "error to check health",
			Description: description,
			StatusCode:  constants.HTTPStatusInternalServerError,
		}, constants.HTTPStatusInternalServerError)

		return ctx.Next()
	}

	ctx.Locals(constants.LogDataKey, &entity.LogDetails{
		Message:    "successfully health checked",
		StatusCode: constants.HTTPStatusOK,
		Response:   check,
	})
	ctx.Locals(constants.LogSeverityKey, constants.SeverityInfo)

	helpers.CreateResponse(ctx, check, constants.HTTPStatusOK)

	return ctx.Next()
}
