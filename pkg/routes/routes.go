package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vitortenor/guardian/pkg/controller"
)

func InitRoutes(r *gin.RouterGroup, tokenController controller.AuthenticationControllerInterface) {
	r.POST("/token", tokenController.GetToken)
	r.POST("/token/refresh", tokenController.RenewToken)
}
