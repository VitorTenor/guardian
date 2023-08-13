package model

import (
	"github.com/vitortenor/guardian/src/config/rest_error"
	"golang.org/x/crypto/bcrypt"
)

func (ud *authDomain) CheckPassword(hashedPassword []byte) *rest_error.Err {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(ud.password))
	if err != nil {
		return rest_error.NewUnauthorizedError("Invalid credentials")
	}
	return nil
}
