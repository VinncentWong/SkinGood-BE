package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"module/app/user/repository"
	"module/domain"
	"module/domain/dto"
	"module/middleware"
	"module/middleware/hash"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:8000/user/callback",
		Scopes:      []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:    google.Endpoint,
	}
	randomString = "random" // you can use random state as a token to prevent CSRF attack
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
	fmt.Println("Create user dipanggil")
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
	users.Password = *password
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
	if hash.CompareContent(tempUser.Password, users.Password) == nil {
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
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "user doesn't exist",
			"success": false,
		})
		return
	}
}

func (*UserService) LoginWithGoogle(c *gin.Context) {
	googleOauthConfig.ClientID = os.Getenv("CLIENT_ID")
	googleOauthConfig.ClientSecret = os.Getenv("CLIENT_SECRET")
	// random string for prevent CSRF attack
	url := googleOauthConfig.AuthCodeURL(randomString)
	fmt.Println(googleOauthConfig)
	fmt.Println("url = " + url)
	fmt.Println("user client id = " + os.Getenv("CLIENT_ID"))
	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func (*UserService) GetGoogleDetails(c *gin.Context) {
	if c.Request.FormValue("state") != randomString {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "state value doesn't same",
			"success": false,
		})
		return
	}
	fmt.Println("state = " + c.Request.FormValue("state"))
	token, err := googleOauthConfig.Exchange(context.Background(), c.Request.FormValue("code"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error occured when transfer authorization code into token",
			"success": false,
		})
		return
	}
	fmt.Println("code = " + c.Request.FormValue("code"))
	fmt.Println("token = " + token.AccessToken)
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error occured when trying get access token",
			"success": false,
		})
		return
	}
	fmt.Printf("response body = %v", response.Body)
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "io error",
			"success": false,
		})
		return
	}
	var container domain.AuthenticatedUser
	err = json.Unmarshal(content, &container)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "success",
		"success": true,
		"data":    container,
	})
}
