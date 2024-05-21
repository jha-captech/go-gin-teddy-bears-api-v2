package actions

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

func (a *Actions) listAllTeddyBears(cCtx *cli.Context) error {
	// get data from db
	bears, err := a.Logic.ListTeddyBears()
	if err != nil {
		return fmt.Errorf("listAllTeddyBears: %w", err)
	}

	// Marshal the struct to JSON with 4-space indentation
	jsonData, err := json.MarshalIndent(bears, "", "    ")
	if err != nil {
		return fmt.Errorf("listAllTeddyBears: %w", err)
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	return nil
}

func (a *Actions) fetchTeddyBearByName(cCtx *cli.Context) error {
	// get user input
	name := cCtx.Args().Get(0)

	// get data from db
	bear, err := a.Logic.FetchTeddyBearByName(name)
	if err != nil {
		return fmt.Errorf("fetchTeddyBearByName: %w", err)
	}

	// Marshal the struct to JSON with 4-space indentation
	jsonData, err := json.MarshalIndent(bear, "", "    ")
	if err != nil {
		return fmt.Errorf("fetchTeddyBearByName: %w", err)
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	return nil
}
