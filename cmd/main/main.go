package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitortenor/guardian/pkg/controller"
	"github.com/vitortenor/guardian/pkg/routes"
	"log"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller.NewAuthenticationControllerInterface())

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error when trying to start the application", err)
	}
}
