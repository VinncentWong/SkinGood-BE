package rest

import (
	"module/app/user/handler"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.Engine, handler *handler.UserService) {
	rGroup := router.Group("/user")
	rGroup.POST("/registration", handler.CreateUser)
	rGroup.POST("/login")
}
