package main

import (
	"mgp_example/internal/controller"

	"github.com/gin-gonic/gin"
)

// @title						My Gin Plus Example API
// @version					2.0
// @Schemes					https http
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	gin.SetMode(gin.ReleaseMode)
	controller.InitRouter().GenerateSwagger() //will generate your swagger
}
