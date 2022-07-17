package repository

import "github.com/gin-gonic/gin"

type UserDao interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
