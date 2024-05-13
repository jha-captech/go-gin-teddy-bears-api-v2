package main

import (
	"log/slog"

	"teddy_bears_api_v2/cmd/fiber/routes"
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// TODO add docker compatibility

func main() {
	Execute()
}

func Execute() {
	slog.Info("running http.main()")

	config.LoggerInit()
	config.DotEnvInit()
	config := config.HydrateConfigFromEnv()

	SwaggerInit(config)

	// logic setup
	dbSetup := logic.SqliteOpen(config)
	logic, err := logic.InitLogic(config, dbSetup)
	if err != nil {
		panic(err)
	}

	// router struct setup
	router := &routes.Router{
		Logic:  logic,
		Config: config,
	}

	// setup app
	app := fiber.New()

	// set middleware
	app.Use(logger.New())

	// setup routes
	router.InitRouter(app)

	// log routes
	if config.GoEnv.Env != "PROD" {
		printAllRoutes(app, config)
	}

	err = app.Listen(config.HTTP.Port)
	slog.Error("app encountered an error", "err", err)
}
