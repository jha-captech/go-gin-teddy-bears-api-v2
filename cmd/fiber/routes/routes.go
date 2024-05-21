package routes

import (
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"

	"github.com/gofiber/fiber/v2"

	_ "teddy_bears_api_v2/models" // needed for swaggo to identify model types
)

type Router struct {
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

func (r *Router) InitRouter(app *fiber.App) {
	appGroup := app.Group("/api")

	r.healthCheck(appGroup.Group("/health-check"))

	r.location(appGroup.Group("/location"))
	r.teddyBear(appGroup.Group("/teddy-bear"))

	r.swagger(app)
}
