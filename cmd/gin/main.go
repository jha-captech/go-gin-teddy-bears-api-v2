package main

import (
	"log/slog"

	"teddy_bears_api_v2/cmd/gin/routes"
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"

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
	config := config.HydrateConfigFromEnv()

	SwaggerInit(config)

	// logic setup
	logic, err := logic.InitLogic(config, logic.SqliteOpen, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// router struct setup
	router := &routes.Router{Logic: logic, Config: config}

	// setup app
	app := gin.Default()

	// set env mode
	switch config.GoEnv.Env {
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
	routes.InitRouter(app, router)

	err = app.Run(config.HTTP.Port)
	slog.Error("app encountered an error", "err", err)
}
