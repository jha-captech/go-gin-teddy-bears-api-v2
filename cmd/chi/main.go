package main

import (
	"log/slog"
	"net/http"
	"strings"

	"teddy_bears_api_v2/cmd/chi/routes"
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/database"
	"teddy_bears_api_v2/logic"

	"github.com/glebarez/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func main() {
	Execute()
}

func Execute() {
	config.LoggerInit()

	slog.Info("running chi.main()")

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
	app := chi.NewRouter()

	// set middleware
	app.Use(middleware.Logger)

	// setup routes
	routes.RoutesInit(app, handler)

	// log routes
	if strings.ToLower(config.Env) != "prod" {
		printAllRoutes(app, config)
	}

	http.ListenAndServe(config.HTTP.Port, app)
}
