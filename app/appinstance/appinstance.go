package appinstance

import (
	"api.default.indicoinnovation.pt/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Config *config.Config
	DB     *pgxpool.Pool
	Server *fiber.App
}

var Data *Application
