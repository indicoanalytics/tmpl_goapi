package health

import (
	"api.default.indicoinnovation.pt/adapters/response"
	"api.default.indicoinnovation.pt/app/usecases/health"
	"api.default.indicoinnovation.pt/config/constants"
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
		return response.CreateError(context, err, "error to check health", check, constants.HTTPStatusInternalServerError)
	}

	return response.CreateSuccess(context, "success to check health", check, constants.HTTPStatusOK)
}
