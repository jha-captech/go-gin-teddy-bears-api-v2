package main

import (
	"log/slog"
	"strings"

	"teddy_bears_api_v2/cmd/gin/routes"
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/database"
	"teddy_bears_api_v2/logic"

	"github.com/glebarez/sqlite"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	Execute()
}

func Execute() {
	slog.Info("running gin.main()")

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
	logic := logic.NewLogic(db)

	// router struct setup
	router := routes.NewHandler(logic, config)

	// setup app
	app := gin.Default()

	// set env mode
	switch strings.ToLower(config.Env) {
	case "dev":
		slog.Info("env set to dev")
		gin.SetMode(gin.DebugMode)
	case "qa":
		slog.Info("env set to qa")
		gin.SetMode(gin.TestMode)
	case "prod":
		slog.Info("env set to prod")
		gin.SetMode(gin.ReleaseMode)
	}

	// setup routes
	routes.RoutesInit(app, router)

	err := app.Run(config.HTTP.Port)
	slog.Error("app encountered an error", "err", err)
}
