package main

import (
	"module/app/user/handler"
	"module/app/user/repository"
	"module/config"
	"module/infrastructure"
	"module/rest"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load Env
	config.InitEnv()

	// Database Connection
	err := infrastructure.ConnectDb()
	if err != nil {
		panic(err.Error())
	}

	userRepository := repository.NewUserDao()
	userHandler := handler.NewUserService(*userRepository)
	// Running App
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(CorsMiddleware())

	rest.CheckHealth(r)
	rest.InitUserRoutes(r, userHandler)
	r.Run(":8000")
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
