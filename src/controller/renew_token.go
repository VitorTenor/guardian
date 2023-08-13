package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vitortenor/guardian/src/config/logger"
	"github.com/vitortenor/guardian/src/config/validation"
	"github.com/vitortenor/guardian/src/controller/model/request"
	"github.com/vitortenor/guardian/src/controller/model/response"
	"go.uber.org/zap"
)

func (ac *authenticationControllerInterface) RenewToken(c *gin.Context) {
	logger.Info("Starting RenewToken",
		zap.String("journey", "renewToken"),
	)

	var refreshToken request.AuthRenewToken
	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		logger.Error("Error when trying to bind JSON",
			err,
			zap.String("journey", "renewToken"),
		)
		restErr := validation.ValidateAuthError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	token, newRefreshToken, err := ac.service.RenewTokenServices(refreshToken.RefreshToken)
	if err != nil {
		logger.Error("Error when trying to renew token",
			err,
			zap.String("journey", "renewToken"),
		)
		c.JSON(err.Code, err)
		return
	}

	response := response.AuthResponse{
		AccessToken:  token,
		RefreshToken: newRefreshToken,
	}

	c.JSON(200, response)
}
