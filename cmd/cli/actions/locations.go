package actions

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func (a *Actions) listAllLocations(cCtx *cli.Context) error {
	// get data from db
	locations, err := a.Logic.ListLocations()
	if err != nil {
		return err
	}

	// Marshal the struct to JSON with 4-space indentation
	jsonData, err := json.MarshalIndent(locations, "", "    ")
	if err != nil {
		return err
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	return nil
}

func (a *Actions) fetchLocationById(cCtx *cli.Context) error {
	// get user input
	idStr := cCtx.Args().Get(0)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("not a valid id. err: %s", "Not a valid id")
	}

	// get data from db
	location, err := a.Logic.FetchLocationByID(id)
	if err != nil {
		return err
	}

	// Marshal the struct to JSON with 4-space indentation
	jsonData, err := json.MarshalIndent(location, "", "    ")
	if err != nil {
		return err
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	return nil
}
