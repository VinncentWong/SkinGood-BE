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

func (dao *UserDao) CreateUser(user domain.User) error {
	err := dao.db.Create(&user)
	if err != nil {
		return err.Error
	}
	return nil
}
