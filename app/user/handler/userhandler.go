package handler

import (
	"module/app/user/repository"
	"module/domain"
	"module/middleware/hash"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserService struct {
	dao repository.UserDao
}

func NewUserService(dao repository.UserDao) *UserService {
	return &UserService{
		dao: dao,
	}
}
func (user *UserService) CreateUser(c *gin.Context) {
	var tempUser domain.User
	err := c.ShouldBindJSON(&tempUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  err.Error(),
			"success": false,
		})
		return
	}
	tempUser.MembershipID, err = uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  err.Error(),
			"success": false,
		})
		return
	}
	password, err := hash.GenerateHashedContent(tempUser.Password)
	tempUser.Password = *password
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  err.Error(),
			"success": false,
		})
		return
	}
	result := user.dao.CreateUser(tempUser)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  err.Error(),
			"success": false,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   tempUser,
	})
}

func (user *UserService) Login(c *gin.Context) {

}
