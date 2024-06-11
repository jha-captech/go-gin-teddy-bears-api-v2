package main

import (
	"log"
	"log/slog"
	"net/http"

	"teddy_bears_api_v2/cmd/net_http/middleware"
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
	app := routes.NewRouter()

	// setup routes
	routes.RoutesInit(app, handler)

	// middleware
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.AddTrailingBackSlash,
	)

	server := http.Server{
		Addr:    config.HTTP.Port,
		Handler: stack(app.Mux),
	}
	log.Fatal(server.ListenAndServe())
}
