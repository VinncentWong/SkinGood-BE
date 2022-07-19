package rest

import (
	"module/app/user/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealth(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}

func InitUserRoutes(router *gin.Engine, handler *handler.UserService) {
	rGroup := router.Group("/user")
	rGroup.POST("/registration", handler.CreateUser)
	rGroup.POST("/login", handler.Login)
	rGroup.GET("/login-w-google", handler.LoginWithGoogle)
	rGroup.GET("/callback", handler.GetGoogleDetails)
}
