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
		// TODO: Log intent to Authorize request

		authBearer := context.GetReqHeaders()["Authorization"]
		authSpec := strings.Split(authBearer, " ")

		if authSpec[0] != "Bearer" {
			return helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, http.StatusUnauthorized)
		}

		if authSpec[1] == "" {
			return helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, http.StatusUnauthorized)
		}

		if _, err := jwt.Verify(authSpec[1]); err != nil {
			return helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, http.StatusUnauthorized)
		}

		return context.Next()
	}
}
