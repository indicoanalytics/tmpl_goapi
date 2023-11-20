package health

import (
	"api.default.indicoinnovation.pt/adapters/logging"
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

func (handler *Handler) Check(context *fiber.Ctx) error {
	check, err := handler.usecase.Check()
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message:    "error to health check",
			Context:    context,
			StatusCode: constants.HTTPStatusInternalServerError,
			Reason:     err.Error(),
		}, constants.SeverityError, nil)

		return helpers.CreateResponse(context, &entity.ErrorResponse{
			Message:     "error to check health",
			Description: err.Error(),
			StatusCode:  constants.HTTPStatusInternalServerError,
		}, constants.HTTPStatusInternalServerError)
	}

	logging.Log(&entity.LogDetails{
		Message:    "successfully health checked",
		Context:    context,
		StatusCode: constants.HTTPStatusOK,
	}, constants.SeverityInfo, nil)

	return helpers.CreateResponse(context, check, constants.HTTPStatusOK)
}
