package controller

import "github.com/gin-gonic/gin"

func NewAuthenticationControllerInterface() AuthenticationControllerInterface {
	return &authenticationControllerInterface{}
}

type AuthenticationControllerInterface interface {
	GetToken(c *gin.Context)
	RenewToken(c *gin.Context)
}
type authenticationControllerInterface struct {
}
