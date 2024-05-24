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
	config := config.MustNewConfig()

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
	actionsSession := actions.NewActions(logicSession, config)

	app := actions.InitActions(actionsSession)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
