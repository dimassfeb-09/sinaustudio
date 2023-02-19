package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
