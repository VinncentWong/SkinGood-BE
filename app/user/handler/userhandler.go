package handler

import (
	"module/app/user/repository"
	"module/domain"
	"module/domain/dto"
	"module/middleware"
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
	var tempUser dto.SignUpDto
	err := c.ShouldBindJSON(&tempUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  err.Error(),
			"success": false,
		})
		return
	}
	users := domain.User{
		FirstName: tempUser.FirstName,
		LastName:  tempUser.LastName,
		Email:     tempUser.Email,
		Password:  tempUser.Password,
		Gender:    tempUser.Gender,
		BirthDate: tempUser.BirthDate,
	}
	users.MembershipID, err = uuid.NewRandom()
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
	result := user.dao.CreateUser(users)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  err.Error(),
			"success": false,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   users,
	})
}

func (user *UserService) Login(c *gin.Context) {
	var tempUser dto.LoginDto
	err := c.ShouldBindJSON(&tempUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  err.Error(),
			"success": false,
		})
		return
	}
	users, err := user.dao.GetByEmail(tempUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  err.Error(),
			"success": false,
		})
		return
	}
	if users == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "user doesn't exist",
			"success": false,
		})
		return
	}
	if users.Password == tempUser.Password {
		token, err := middleware.GenerateToken(*users)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  err.Error(),
				"success": false,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   *users,
			"token":  token,
		})
	}
}
