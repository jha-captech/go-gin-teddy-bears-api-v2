package main

import (
	"log/slog"
	"strings"

	"teddy_bears_api_v2/cmd/fiber/routes"
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/database"
	"teddy_bears_api_v2/logic"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func main() {
	Execute()
}

func Execute() {
	slog.Info("running fiber.main()")

	config.LoggerInit()
	config.DotEnvInit()
	config := config.MustNewConfig()

	SwaggerInit(config)

	// database connect
	db := database.MustNewDatabase(
		config,
		sqlite.Open(config.Database.Name),
		gorm.Config{},
		config.Database.ConnectionRetry,
	)

	// logic setup
	logicSession := logic.NewLogic(db)

	// router struct setup
	handler := routes.NewHandler(logicSession, config)

	// setup app
	app := fiber.New()

	// set middleware
	app.Use(logger.New())

	// setup routes
	routes.RoutesInit(app, handler)

	// log routes
	if strings.ToLower(config.Env) != "prod" {
		printAllRoutes(app, config)
	}

	err := app.Listen(config.HTTP.Port)
	slog.Error("app encountered an error", "err", err)
}
