package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/vitortenor/guardian/pkg/config/rest_error"
	"os"
	"time"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *authDomain) GenerateToken() (string, *rest_error.Err) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.ID,
		"email": ud.email,
		"name":  ud.name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_error.NewInternalServerError("Error when trying to generate token", err)
	}

	return tokenString, nil
}
