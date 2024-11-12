package utilities

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, salt int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)

	return string(bytes), err
}

func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func NewSalt() int {
	r := rand.Intn(29)
	return r
}
