package main

import (
	"fmt"
	"net/http"
	"time"

	"api.default.indicoinnovation.pt/app/appinstance"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/handler/health"
	"api.default.indicoinnovation.pt/middleware"
	"api.default.indicoinnovation.pt/pkg/app"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func route() *fiber.App {
	allowedOrigins := constants.AllowedOrigins
	if constants.Environment != constants.Production {
		allowedOrigins += fmt.Sprintf(", %s", constants.AllowedStageOrigins)
	}

	appinstance.Data.Server.Use(logger.New())
	appinstance.Data.Server.Use(recover.New())
	appinstance.Data.Server.Use(favicon.New())
	appinstance.Data.Server.Use(cors.New(cors.Config{
		AllowMethods: constants.AllowedMethods,
		AllowOrigins: allowedOrigins,
		AllowHeaders: constants.AllowedHeaders,
	}))
	appinstance.Data.Server.Use(middleware.ValidateContentType())
	appinstance.Data.Server.Use(middleware.SecurityHeaders())
	appinstance.Data.Server.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	root := appinstance.Data.Server.Group("/")
	root.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return helpers.Contains(constants.AllowedUnthrottledIPs, c.IP())
		},
		Max:        constants.MaxResquestLimit,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			helpers.CreateResponse(c, &entity.ErrorResponse{
				Message:     "Calls Limit Reached",
				Description: "Rate Limit reached",
				StatusCode:  http.StatusTooManyRequests,
			}, http.StatusTooManyRequests)

			return nil
		},
	}))

	root.Get("/health", health.Handle().Check, app.Log)

	apiGroup := root.Group("/api")
	apiRoutesV1(apiGroup)

	return appinstance.Data.Server
}
