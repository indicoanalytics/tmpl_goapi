package constants

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	Port                      = "9090"
	MainLoggerName            = "health"
	MainServiceName           = MainLoggerName + "_api"
	MaxResquestLimit          = 2
	AccessTokenExpirationTime = 15
	SignedURLExp              = 60
	Audience                  = "https://iam.services.indicoinnovation.pt"
)

var (
	Debug, _     = strconv.ParseBool(os.Getenv("DEBUG"))
	GcpProjectID = os.Getenv("PROJECT")
	SecretPrefix = os.Getenv("SEC_PREFIX")
)

var (
	AllowedContentTypes   = []string{fiber.MIMEApplicationJSON}
	AllowedHeaders        = "X-Session-Id, Authorization, Content-Type, Accept, Origin"
	AllowedMethods        = "GET,POST,OPTIONS"
	AllowedOrigins        = "https://tbd, https://tbd"
	AllowedStageOrigins   = "https://localhost:3000, http://localhost:3000"
	AllowedUnthrottledIPs = []string{"127.0.0.1"}
)
