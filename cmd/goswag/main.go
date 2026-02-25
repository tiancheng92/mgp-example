package main

import "mgp_example/controller"

// @title						My Gin Plus Example API
// @version					2.0
// @Schemes					https http
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	controller.InitRouter().GenerateSwagger() //will generate your swagger
}
