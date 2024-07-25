package middleware

import (
	"strings"

	"api.default.indicoinnovation.pt/adapters/jwt"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

func Authorize() func(context *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		// TODO: Implement mechanism to authorize requests from given IP to measure endpoint status and metrics
		// TODO: Log intent to Authorize request

		authorization := context.GetReqHeaders()["Authorization"]
		if len(authorization) == 0 {
			helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: constants.HTTPStatusUnauthorized,
			}, constants.HTTPStatusUnauthorized)
		}

		authSpec := strings.Split(authorization[0], " ")
		if authSpec[0] != "Bearer" {
			helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: constants.HTTPStatusUnauthorized,
			}, constants.HTTPStatusUnauthorized)
		}

		if authSpec[1] == "" {
			helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: constants.HTTPStatusUnauthorized,
			}, constants.HTTPStatusUnauthorized)
		}

		if !jwt.New().Validate(authSpec[1]) {
			helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Unauthorized",
				StatusCode: constants.HTTPStatusUnauthorized,
			}, constants.HTTPStatusUnauthorized)
		}

		return context.Next()
	}
}
