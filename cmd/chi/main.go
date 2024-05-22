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

	slog.Info("running http.main()")

	config.DotEnvInit()
	config, err := config.NewConfig()
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
	router := &routes.Handler{Logic: logic, Config: config}

	// setup app
	app := chi.NewRouter()

	// set middleware
	app.Use(middleware.Logger)

	// setup routes
	router.InitRouter(app)

	// log routes
	if strings.ToLower(config.Env) != "prod" {
		printAllRoutes(app, config)
	}

	http.ListenAndServe(config.HTTP.Port, app)
}
