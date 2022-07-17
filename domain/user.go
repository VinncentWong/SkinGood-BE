package domain

import (
	"module/domain/enum"
	"time"

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
	Gender       enum.Gender
	BirthDate    time.Time
	Password     string
}

type Address struct {
	gorm.Model
	User   User
	UserID uint64
}
