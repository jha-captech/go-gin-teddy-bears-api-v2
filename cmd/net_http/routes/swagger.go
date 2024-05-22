package routes

import (
	"fmt"
	"log/slog"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (h Handler) swagger(r Router) {
	r.get("/*", httpSwagger.WrapHandler)
	slog.Info(
		fmt.Sprintf(
			"Swagger URL: http://%s%s/swagger/index.html",
			h.Config.HTTP.Domain,
			h.Config.HTTP.Port,
		),
	)
}
