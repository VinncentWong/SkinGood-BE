package implementation

import (
	"module/infrastructure"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserImplementation struct {
	db *gorm.DB
}

func NewUser() *UserImplementation {
	return &UserImplementation{}
}

func InitDb() {
	NewUser().db = infrastructure.GetDb()
}

func (user *UserImplementation) CreateUser(c *gin.Context) {

}
