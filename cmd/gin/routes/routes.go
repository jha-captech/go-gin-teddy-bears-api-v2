package routes

import (
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"

	"github.com/gin-gonic/gin"
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

func RoutesInit(app *gin.Engine, handler Handler) {
	r := app.Group("api")

	handler.healthCheck(r.Group("health-check"))

	handler.location(r.Group("location"))
	handler.teddyBear(r.Group("teddy-bear"))

	handler.swagger(app)
}
