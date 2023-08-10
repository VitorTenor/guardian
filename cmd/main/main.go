package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vitortenor/guardian/pkg/config/database/mariadb"
	"github.com/vitortenor/guardian/pkg/config/logger"
	"github.com/vitortenor/guardian/pkg/controller"
	"github.com/vitortenor/guardian/pkg/routes"
	"log"
)

func main() {
	logger.Info("Starting the application...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	databaseConnection, err := mariadb.NewMariaDBConnection()
	if err != nil {
		log.Fatal("Error when trying to connect to database", err)
	}
	router := gin.Default()

	fmt.Sprintf("databaseConnection: %v", databaseConnection)

	routes.InitRoutes(&router.RouterGroup, controller.NewAuthenticationControllerInterface())

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error when trying to start the application", err)
	}
}
