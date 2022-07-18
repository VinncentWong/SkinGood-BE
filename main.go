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
	r := gin.Default()
	rest.CheckHealth(r)
	rest.InitUserRoutes(r, userHandler)
	r.Run(":8000")
}
