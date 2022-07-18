package hash

import "golang.org/x/crypto/bcrypt"

func GenerateHashedContent(rawPassword string) (*string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	var newPassword = string(bytePassword)
	return &newPassword, nil
}

func CompareContent(rawPassword string, hashedPassword string) error {
	valid := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return valid
}
