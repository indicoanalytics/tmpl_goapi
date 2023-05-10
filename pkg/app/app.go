package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"api.default.indicoinnovation.pt/adapters/logging"
	"api.default.indicoinnovation.pt/clients/iam"
	"api.default.indicoinnovation.pt/config"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"api.default.indicoinnovation.pt/pkg/postgres"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Application struct {
	DBInstance *gorm.DB
	Config     *config.Config
	Server     *fiber.App
}

var Inst *Application

func ApplicationInit() {
	configs := config.New()

	iam.New()

	databaseConnection, err := postgres.Connect(configs.DBString, logger.LogLevel(configs.DBLogMode), configs.Debug)
	if err != nil {
		log.Panicln(fmt.Sprintf("Failed to initialize %s DB Connection", configs.DBString), err)
	}

	log.Printf("Database is now connected")

	Inst = &Application{
		DBInstance: databaseConnection,
		Config:     configs,
		Server: fiber.New(fiber.Config{
			ServerHeader: "Death Star",
			ErrorHandler: customErrorHandler,
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
			Prefork:      true,
		}),
	}
}

func Setup() {
	var err error
	if Inst.Config.Environment != "production" {
		err = Inst.Server.ListenTLS(fmt.Sprintf(":%s", Inst.Config.Port), "cert.pem", "key.pem")
	} else {
		err = Inst.Server.Listen(fmt.Sprintf(":%s", Inst.Config.Port))
	}

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func customErrorHandler(context *fiber.Ctx, err error) error {
	var code int = fiber.StatusInternalServerError
	var capturedError *fiber.Error
	message := "unknown error"

	if errors.As(err, &capturedError) {
		code = capturedError.Code
		if code == fiber.StatusNotFound {
			message = "route not found"
		}
	}

	var errorResponse *entity.ErrorResponse

	erro := json.Unmarshal([]byte(err.Error()), &errorResponse)
	if erro != nil {
		errorResponse = &entity.ErrorResponse{
			Message:     message,
			StatusCode:  code,
			Description: err.Error(),
		}
	}

	go logging.Log(
		&entity.LogDetails{
			Message:     message,
			Reason:      err.Error(),
			StatusCode:  code,
			RequestData: string(context.Body()),
		},
		"critical",
		nil,
	)

	return helpers.CreateResponse(context, errorResponse, code) //nolint: wrapcheck
}
