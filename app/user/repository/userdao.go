package repository

import (
	"module/domain"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (conn *UserDao) CreateUser(user domain.User) error {
	err := conn.db.Create(&user)
	if err != nil {
		return err.Error
	}
	return nil
}

func (conn *UserDao) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := conn.db.Where("email = ?", email).Take(&user)
	if err != nil {
		return nil, err.Error
	}
	return &user, nil
}
