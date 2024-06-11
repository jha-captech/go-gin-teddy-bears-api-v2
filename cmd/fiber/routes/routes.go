package routes

import (
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"

	"github.com/gofiber/fiber/v2"
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

func NewHandler(logic logic.Logic, config config.Configuration) Handler {
	return Handler{
		Logic:  logic,
		Config: config,
	}
}

func RoutesInit(app *fiber.App, handler Handler) {
	appGroup := app.Group("/api")

	handler.healthCheck(appGroup.Group("/health-check"))

	handler.location(appGroup.Group("/location"))
	handler.teddyBear(appGroup.Group("/teddy-bear"))

	handler.swagger(app)
}
