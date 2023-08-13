package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vitortenor/guardian/src/config/logger"
	"github.com/vitortenor/guardian/src/config/validation"
	"github.com/vitortenor/guardian/src/controller/model/request"
	"github.com/vitortenor/guardian/src/controller/model/response"
	"github.com/vitortenor/guardian/src/model"
	"go.uber.org/zap"
)

func (ac *authenticationControllerInterface) GetToken(c *gin.Context) {

	logger.Info("Init GetToken Controller",
		zap.String("journey", "getToken"),
	)
	var authLogin request.AuthLogin

	if err := c.ShouldBindJSON(&authLogin); err != nil {
		logger.Error("Error when trying to bind JSON",
			err,
			zap.String("journey", "getToken"),
		)
		restErr := validation.ValidateAuthError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	authDomain := model.AuthLoginDomain(authLogin.Email, authLogin.Password)

	token, refreshToken, err := ac.service.AuthUserServices(authDomain)
	if err != nil {
		logger.Error("Error when trying to authenticate user",
			err,
			zap.String("journey", "getToken"),
		)
		c.JSON(err.Code, err)
		return
	}

	response := response.AuthResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	c.JSON(200, response)
}
