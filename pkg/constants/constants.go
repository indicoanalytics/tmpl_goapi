package constants

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	Port             = "9090"
	MainLoggerName   = "health"
	MainServiceName  = MainLoggerName + "_api"
	MaxResquestLimit = 2
)

var (
	Debug, _     = strconv.ParseBool(os.Getenv("DEBUG"))
	GcpProjectID = os.Getenv("PROJECT")
	SecretPrefix = os.Getenv("SEC_PREFIX")
)

var (
	AllowedContentTypes = []string{fiber.MIMEApplicationJSON}
	AllowedOrigins      = "https://tbd, https://tbd"
	AllowedStageOrigins = "https://localhost:3000, http://localhost:3000"
)
