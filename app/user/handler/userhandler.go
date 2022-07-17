package handler

import "module/app/user/repository"

type UserService struct {
	repository.UserDao
}
