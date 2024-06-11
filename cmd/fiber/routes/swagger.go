package routes

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func (h Handler) swagger(r fiber.Router) {
	r.Get("/swagger/*", fiberSwagger.WrapHandler)
	slog.Info(
		fmt.Sprintf(
			"Swagger URL: http://%s%s/swagger/index.html",
			h.Config.HTTP.Domain,
			h.Config.HTTP.Port,
		),
	)
}
