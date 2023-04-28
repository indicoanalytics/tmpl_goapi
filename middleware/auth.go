package middleware

import (
	"net/http"
	"strings"

	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"api.default.indicoinnovation.pt/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func Authorize() func(context *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		authBearer := context.GetReqHeaders()["Authorization"]

		if authBearer == "" {
			return helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, http.StatusUnauthorized)
		}

		if _, err := jwt.Verify(strings.Split(authBearer, " ")[1]); err != nil {
			return helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, http.StatusUnauthorized)
		}

		return context.Next()
	}
}
