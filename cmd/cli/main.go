package main

import (
	"log"
	"os"

	"teddy_bears_api_v2/cmd/cli/actions"
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
	config.DotEnvInit()
	config := config.HydrateConfigFromEnv()

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

	// a struct setup
	a := &actions.Actions{Logic: logic, Config: config}

	app := actions.InitActions(a)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
