package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vitortenor/guardian/pkg/model/service"
)

func NewAuthenticationControllerInterface(svc service.AuthDomainService) AuthenticationControllerInterface {
	return &authenticationControllerInterface{
		service: svc,
	}
}

type AuthenticationControllerInterface interface {
	GetToken(c *gin.Context)
	RenewToken(c *gin.Context)
}
type authenticationControllerInterface struct {
	service service.AuthDomainService
}
