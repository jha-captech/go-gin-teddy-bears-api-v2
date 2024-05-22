package main

import (
	"log"
	"log/slog"
	"net/http"

	"teddy_bears_api_v2/cmd/net_http/routes"
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/database"
	"teddy_bears_api_v2/logic"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

func main() {
	Execute()
}

func Execute() {
	config.LoggerInit()

	slog.Info("running net_http.main()")

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
	handler := routes.NewHandler(logic, config)

	// setup app
	app := routes.NewRouter()

	// set middleware
	// app.Use(middleware.Logger)

	// setup routes
	handler.InitRouter(app)

	server := http.Server{
		Addr:    config.HTTP.Port,
		Handler: app.Mux,
	}
	log.Fatal(server.ListenAndServe())
}
