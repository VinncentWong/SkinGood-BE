package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	MembershipID uuid.UUID
	Gender       string
	BirthDate    string
	Password     string
	Picture      []byte
}

type Address struct {
	gorm.Model
	User   User
	UserID uint64
}
