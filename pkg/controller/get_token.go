package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vitortenor/guardian/pkg/config/logger"
	"github.com/vitortenor/guardian/pkg/config/validation"
	"github.com/vitortenor/guardian/pkg/controller/model/request"
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

	logger.Info(fmt.Sprintf(authLogin.Email))
	logger.Info(fmt.Sprintf(authLogin.Password))

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
