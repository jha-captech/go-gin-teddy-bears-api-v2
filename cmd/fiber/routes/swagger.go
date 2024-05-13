package routes

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func (router Router) swagger(fr fiber.Router) {
	fr.Get("/swagger/*", fiberSwagger.WrapHandler)
	slog.Info(
		fmt.Sprintf(
			"Swagger URL: http://%s%s/swagger/index.html",
			router.Config.HTTP.Domain,
			router.Config.HTTP.Port,
		),
	)
}
