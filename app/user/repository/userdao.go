package repository

import (
	"fmt"
	"module/domain"
	"module/infrastructure"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{
		db: infrastructure.GetDb(),
	}
}

func (conn *UserDao) CreateUser(user domain.User) error {
	fmt.Println("create-user-dao")
	err := conn.db.Create(&user)
	if err != nil {
		return err.Error
	}
	return nil
}

func (conn *UserDao) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := conn.db.Where("email = ?", email).Take(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}
