package actions

import (
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"

	"github.com/urfave/cli/v2"
)

type Actions struct {
	Logic  logic.Logic
	Config config.Configuration
}

func NewActions(logic logic.Logic, config config.Configuration) Actions {
	return Actions{
		Logic:  logic,
		Config: config,
	}
}

// TODO: Add struct/slice to JSON generic converter

func InitActions(a Actions) *cli.App {
	return &cli.App{
		Name:  "Teddy Bears",
		Usage: "Teddy Bear Tech Challenge Console APP",
		Commands: []*cli.Command{
			{
				Name:    "location",
				Aliases: []string{"loc"},
				Usage:   "location group actions",
				Subcommands: []*cli.Command{
					{
						Name:   "list",
						Usage:  "list all picnic locations",
						Action: a.listAllLocations,
					},
					{
						Name:   "fetch_by_id",
						Usage:  "fetch a picnic location by id",
						Action: a.fetchLocationById,
					},
				},
			},
			{
				Name:    "teddy_bear",
				Aliases: []string{"bear"},
				Usage:   "teddy bear group actions",
				Subcommands: []*cli.Command{
					{
						Name:   "list",
						Usage:  "list all teddy bears",
						Action: a.listAllTeddyBears,
					},
					{
						Name:   "fetch_by_id",
						Usage:  "fetch a teddy bear by name",
						Action: a.fetchTeddyBearByName,
					},
				},
			},
		},
	}
}
