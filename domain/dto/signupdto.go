package dto

type SignUpDto struct {
	Email     string `json:"email" binding:"required" validate:"email"`
	FirstName string `json:"firstname" binding:"required" validate:"min=4"`
	LastName  string `json:"lastname" binding:"required" validate:"min=4"`
	Password  string `json:"password" binding:"required" validate:"min=6"`
	Gender    string `json:"gender" binding:"required"`
	BirthDate string `json:"birth" binding:"required"`
}
