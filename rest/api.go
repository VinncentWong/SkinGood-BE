package rest

import (
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.Engine) {
	rGroup := router.Group("/user")
	rGroup.POST("/registration")
	rGroup.POST("/login")
}
