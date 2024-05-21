package routes

import (
	"fmt"
	"log/slog"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (h Handler) swagger(r chi.Router) {
	r.Get("/*", httpSwagger.WrapHandler)
	slog.Info(
		fmt.Sprintf(
			"Swagger URL: http://%s%s/swagger/index.html",
			h.Config.HTTP.Domain,
			h.Config.HTTP.Port,
		),
	)
}
