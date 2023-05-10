package health

import (
	"net/http"

	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func Handle() *Handler {
	return &Handler{}
}

func (handler *Handler) Check(context *fiber.Ctx) error {
	return helpers.CreateResponse(context, &entity.SuccessfulResponse{
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}
