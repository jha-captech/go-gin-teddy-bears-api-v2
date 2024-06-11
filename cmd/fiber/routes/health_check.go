package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) healthCheck(r fiber.Router) {
	r.Get("/", h.runHealthCheck)
}

// @Summary		Health check response
// @Description	Health check response
// @Tags		health-check
// @Accept		json
// @Produce		json
// @Success		200				{object}	routes.responseMessage
// @Router		/health-check 	[GET]
func (h *Handler) runHealthCheck(c *fiber.Ctx) error {
	return c.
		Status(http.StatusOK).
		JSON(responseMessage{Message: "Health check response."})
}
