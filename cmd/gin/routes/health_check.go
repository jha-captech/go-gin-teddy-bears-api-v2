package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router Handler) healthCheck(r *gin.RouterGroup) {
	r.GET("/", router.runHealthCheck)
}

// @Summary		Health check response
// @Description	Health check response
// @Tags		health-check
// @Accept		json
// @Produce		json
// @Success		200				{object}	routes.responseMessage
// @Router		/health-check 	[GET]
func (router Handler) runHealthCheck(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		responseMessage{Message: "Health check response."},
	)
}
