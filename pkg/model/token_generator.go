package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/vitortenor/guardian/pkg/config/logger"
	"github.com/vitortenor/guardian/pkg/config/rest_error"
	"go.uber.org/zap"
	"os"
	"time"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *authDomain) GenerateTokens() (string, string, *rest_error.Err) {
	logger.Info("Init GenerateTokens",
		zap.String("journey", "generateTokens"),
	)

	secret := os.Getenv(JWT_SECRET_KEY)

	accessClaims := jwt.MapClaims{
		"id":    ud.ID,
		"email": ud.email,
		"name":  ud.name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	accessTokenString, errr := accessToken.SignedString([]byte(secret))
	if errr != nil {
		logger.Error("Error when trying to generate access token",
			errr,
			zap.String("journey", "generateTokens"),
		)
		return "", "", rest_error.NewInternalServerError("Error when trying to generate access token", errr)
	}

	logger.Info("Access Token Generated With Success, Starting to generate refresh token",
		zap.String("journey", "generateTokens"),
	)

	refreshTokenString, err := ud.generateRefreshToken(secret)
	if err != nil {
		logger.Error("Error when trying to generate refresh token",
			err,
			zap.String("journey", "generateTokens"),
		)
		restErr := rest_error.NewInternalServerError("Error when trying to generate refresh token", err)
		return "", "", restErr
	}

	logger.Info("Refresh Token Generated With Success",
		zap.String("journey", "generateTokens"),
	)

	logger.Info("GenerateTokens OK",
		zap.String("journey", "generateTokens"),
	)

	return accessTokenString, refreshTokenString, nil
}

func (ud *authDomain) generateRefreshToken(secret string) (string, *rest_error.Err) {
	logger.Info("Init generateRefreshToken",
		zap.String("journey", "generateRefreshToken"),
	)

	refreshClaims := jwt.MapClaims{
		"id":    ud.ID,
		"email": ud.email,
		"name":  ud.name,
		"exp":   time.Now().Add(time.Hour * 7 * 24).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		logger.Error("Error when trying to generate refresh token",
			err,
			zap.String("journey", "generateRefreshToken"),
		)

		return "", rest_error.NewInternalServerError("Error when trying to generate refresh token", err)
	}

	logger.Info("GenerateRefreshToken OK",
		zap.String("journey", "generateRefreshToken"),
	)

	return refreshTokenString, nil
}

func (ud *authDomain) VerifyAndRefreshToken(refreshTokenString string) (string, string, *rest_error.Err) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", "", rest_error.NewUnauthorizedError("Invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", rest_error.NewUnauthorizedError("Invalid refresh token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		return "", "", rest_error.NewUnauthorizedError("Refresh token expired")
	}

	ud.ID = int(uint64(claims["id"].(float64)))
	ud.email = claims["email"].(string)
	ud.name = claims["name"].(string)

	newAccessToken, newRefreshToken, erro := ud.GenerateTokens()
	if erro != nil {
		return "", "", rest_error.NewInternalServerError("Error when trying to generate new tokens", erro)
	}

	return newAccessToken, newRefreshToken, nil
}
