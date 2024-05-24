package routes

import (
	"net/http"

	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"
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

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() Router {
	return Router{Mux: http.NewServeMux()}
}

func NewHandler(logic logic.Logic, config config.Configuration) Handler {
	return Handler{
		Logic:  logic,
		Config: config,
	}
}

func RoutesInit(app Router, handler Handler) {
	app.group("/api", func(r Router) {
		r.group("/health-check", handler.healthCheck)

		r.group("/location", handler.location)
		r.group("/teddy-bear", handler.teddyBear)
	})

	app.group("/swagger", handler.swagger)
}
