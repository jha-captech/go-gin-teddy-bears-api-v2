package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r *Router) healthCheck(fr fiber.Router) {
	fr.Get("/", r.runHealthCheck)
}

// @Summary		Health check response
// @Description	Health check response
// @Tags		health-check
// @Accept		json
// @Produce		json
// @Success		200				{object}	routes.responseMessage
// @Router		/health-check 	[GET]
func (r *Router) runHealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).
		JSON(responseMessage{Message: "Health check response."})
}
