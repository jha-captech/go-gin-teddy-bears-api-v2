package routes

import (
	"net/http"

	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"

	_ "teddy_bears_api_v2/models" // needed for swaggo to identify model types
)

type Handler struct {
	Logic  *logic.Logic
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

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() Router {
	return Router{Mux: http.NewServeMux()}
}

func NewHandler(logic *logic.Logic, config config.Configuration) Handler {
	return Handler{
		Logic:  logic,
		Config: config,
	}
}

func (h *Handler) InitRouter(r Router) {
	r.group("/api", func(r Router) {
		r.group("/health-check", h.healthCheck)

		r.group("/location", h.location)
		r.group("/teddy-bear", h.teddyBear)
	})

	r.group("/swagger", h.swagger)
}
