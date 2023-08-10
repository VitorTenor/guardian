package model

import (
	"github.com/vitortenor/guardian/pkg/config/rest_error"
	"golang.org/x/crypto/bcrypt"
)

func (ud *authDomain) EncryptPassword() *rest_error.Err {
	hashedPassword, err := hashPassword(ud.password)
	if err != nil {
		return rest_error.NewInternalServerError("Error when trying to encrypt password", err)
	}
	ud.password = hashedPassword
	return nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (ud *authDomain) CheckPassword(hashedPassword string) *rest_error.Err {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(ud.password))
	if err != nil {
		return rest_error.NewUnauthorizedError("Invalid credentials")
	}
	return nil
}
