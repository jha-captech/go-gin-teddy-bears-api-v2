package routes

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (router Handler) swagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	slog.Info(
		fmt.Sprintf(
			"Swagger URL: http://%s%s/swagger/index.html",
			router.Config.HTTP.Domain,
			router.Config.HTTP.Port,
		),
	)
}
