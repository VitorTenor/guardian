package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vitortenor/guardian/src/controller"
)

func InitRoutes(r *gin.RouterGroup, tokenController controller.AuthenticationControllerInterface) {
	r.POST("/token", tokenController.GetToken)
	r.POST("/token/refresh", tokenController.RenewToken)
}
