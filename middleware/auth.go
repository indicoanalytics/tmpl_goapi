package middleware

import (
	"strings"

	"api.default.indicoinnovation.pt/adapters/jwt"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"github.com/gofiber/fiber/v2"
)

func Authorize() func(context *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		// TODO: Implement mechanism to authorize requests from given IP to measure endpoint status and metrics
		// TODO: Log intent to Authorize request

		authorization := context.GetReqHeaders()["Authorization"]
		if len(authorization) == 0 {
			return context.Status(constants.HTTPStatusUnauthorized).JSON(&entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: constants.HTTPStatusUnauthorized,
			})
		}

		authSpec := strings.Split(authorization[0], " ")
		if authSpec[0] != "Bearer" {
			return context.Status(constants.HTTPStatusUnauthorized).JSON(&entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: constants.HTTPStatusUnauthorized,
			})
		}

		if authSpec[1] == "" {
			return context.Status(constants.HTTPStatusUnauthorized).JSON(&entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: constants.HTTPStatusUnauthorized,
			})
		}

		if !jwt.New().Validate(authSpec[1]) {
			return context.Status(constants.HTTPStatusUnauthorized).JSON(&entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: constants.HTTPStatusUnauthorized,
			})
		}

		return context.Next()
	}
}
