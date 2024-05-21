package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"text/tabwriter"

	"teddy_bears_api_v2/config"

	"github.com/go-chi/chi/v5"
)

func printAllRoutes(app *chi.Mux, config config.Configuration) {
	slog.Info("Registered endpoints:\n")

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	defer writer.Flush()

	chi.Walk(
		app,
		func(
			method string,
			route string,
			handler http.Handler,
			middlewares ...func(http.Handler) http.Handler,
		) error {
			fmt.Fprintf(
				writer,
				"[%s]\thttp://%s%s%s\n",
				method,
				config.HTTP.Domain,
				config.HTTP.Port,
				route,
			)
			return nil
		},
	)
}
