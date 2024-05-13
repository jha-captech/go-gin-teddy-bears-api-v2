package routes

import (
	"teddy_bears_api_v2/config"
	"teddy_bears_api_v2/logic"

	"github.com/gin-gonic/gin"

	_ "teddy_bears_api_v2/models" // needed for swaggo to identify model types
)

type Router struct {
	Logic  *logic.Logic
	Config *config.Config
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

func InitRouter(app *gin.Engine, router *Router) {
	r := app.Group("api")

	router.healthCheck(r.Group("health-check"))

	router.location(r.Group("location"))
	router.teddyBear(r.Group("teddy-bear"))

	router.swagger(app)
}
