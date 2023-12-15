package health

import (
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
		ctx.Locals("log_message", "error to check health")
		ctx.Locals("error_reason", err.Error())
		ctx.Locals("log_severity", constants.SeverityError)
		ctx.Locals("status_code", constants.HTTPStatusInternalServerError)
		ctx.Locals("response", &entity.ErrorResponse{
			Message:     "error to check health",
			Description: err.Error(),
			StatusCode:  constants.HTTPStatusInternalServerError,
		})

		helpers.CreateResponse(ctx, ctx.Locals("response"), constants.HTTPStatusInternalServerError)

		return ctx.Next()
	}

	ctx.Locals("log_message", "successfully health checked")
	ctx.Locals("response", check)

	helpers.CreateResponse(ctx, check, constants.HTTPStatusOK)

	return ctx.Next()
}
