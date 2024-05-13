package main

import (
	"fmt"
	"log/slog"
	"os"
	"text/tabwriter"

	"teddy_bears_api_v2/config"

	"github.com/gofiber/fiber/v2"
)

func printAllRoutes(app *fiber.App, config *config.Config) {
	slog.Info("Registered endpoints:\n")

	writer := tabwriter.NewWriter(
		os.Stdout,
		0,
		8,
		1,
		'\t',
		tabwriter.AlignRight,
	)

	defer writer.Flush()

	for _, r := range app.Stack() {
		for _, rr := range r {
			if rr.Path != "" && rr.Path != "/" && rr.Method != "HEAD" {
				fmt.Fprintf(
					writer,
					"%s\thttp://%s%s%s\n",
					rr.Method,
					config.HTTP.Domain,
					config.HTTP.Port,
					rr.Path,
				)
			}
		}
	}
}
