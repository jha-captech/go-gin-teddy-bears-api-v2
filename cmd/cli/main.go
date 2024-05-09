package main

import (
	"log"
	"os"

	"teddy_bears_api_v2/cmd/cli/actions"
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"
)

func main() {
	Execute()
}

func Execute() {
	config.LoggerInit()
	config.DotEnvInit()
	config := config.HydrateConfigFromEnv()

	// logic setup
	logic, err := logic.InitLogic(config)
	if err != nil {
		panic(err)
	}

	// a struct setup
	a := &actions.Actions{
		Logic:  logic,
		Config: config,
	}

	app := actions.InitActions(a)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
