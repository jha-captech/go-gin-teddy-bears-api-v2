package routes

import (
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Logic  logic.Logic
	Config config.Configuration
}

type responseMessage struct {
	Message string `json:"message"`
}

type responseID struct {
	ObjectID int `json:"object_id"`
}

type responseError struct {
	Error string `json:"error"`
}

// Setup and return new routes.Router struct.
func NewHandler(logic logic.Logic, config config.Configuration) Handler {
	return Handler{
		Logic:  logic,
		Config: config,
	}
}

func RoutesInit(app *chi.Mux, h Handler) {
	app.Route("/api", func(r chi.Router) {
		r.Route("/health-check", h.healthCheck)

		r.Route("/location", h.location)
		r.Route("/teddy-bear", h.teddyBear)
	})

	app.Route("/swagger", h.swagger)
}
