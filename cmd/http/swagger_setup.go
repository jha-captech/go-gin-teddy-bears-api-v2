package main

import (
	"teddy_bears_api_v2/cmd/http/docs"
	"teddy_bears_api_v2/config"
)

func SwaggerInit(config *config.Config) {
	docs.SwaggerInfo.Title = "Teddy Bears Go Gin API"
	docs.SwaggerInfo.Description = "Practice Go Gin API using data from .Net Tech Challenge"
	docs.SwaggerInfo.Version = "1.0"

	docs.SwaggerInfo.Host = config.HTTP.Domain + config.HTTP.Port
	docs.SwaggerInfo.BasePath = "/api"

	docs.SwaggerInfo.Schemes = []string{"http"}
}
