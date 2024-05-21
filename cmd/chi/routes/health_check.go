package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) healthCheck(r chi.Router) {
	r.Get("/", h.runHealthCheck)
}

// @Summary		Health check response
// @Description	Health check response
// @Tags		health-check
// @Accept		json
// @Produce		json
// @Success		200				{object}	routes.responseMessage
// @Router		/health-check 	[GET]
func (h *Handler) runHealthCheck(w http.ResponseWriter, r *http.Request) {
	encode(
		w,
		http.StatusOK,
		responseMessage{Message: "Health check response."},
	)
	// encode(
	// 	w,
	// 	,
	// 	,
	// )
}
