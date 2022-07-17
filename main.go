package main

import (
	"module/app/user/implementation"
	"module/config"
	"module/infrastructure"
	"module/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Env
	config.InitEnv()

	// Database Connection
	infrastructure.ConnectDb()
	implementation.InitDb()

	// Running App
	r := gin.Default()
	rest.InitUserRoutes(r)
	r.Run()
}
