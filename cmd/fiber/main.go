package main

import (
	"log/slog"

	"teddy_bears_api_v2/cmd/fiber/routes"
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/database"
	"teddy_bears_api_v2/logic"

	"teddy_bears_api_v2/database"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	Execute()
}

func Execute() {
	slog.Info("running http.main()")

	config.LoggerInit()
	config.DotEnvInit()
	config, err := config.HydrateConfigFromEnv()
	if err != nil {
		panic(err)
	}

	SwaggerInit(config)

	// database connect
	db, err := database.Connect(
		config,
		sqlite.Open(config.Database.Name),
		gorm.Config{},
		config.Database.ConnectionRetry,
	)
	if err != nil {
		panic(err)
	}

	// logic setup
	logic, err := logic.InitLogic(db)
	if err != nil {
		panic(err)
	}

	// router struct setup
	router := &routes.Router{Logic: logic, Config: config}

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
