package main

import (
	"database/sql"
	"fmt"
	"github.com/vitortenor/guardian/src/controller"
	"github.com/vitortenor/guardian/src/model/repository"
	"github.com/vitortenor/guardian/src/model/service"
)

func initAuthControllerDependencies(database *sql.DB) controller.AuthenticationControllerInterface {
	fmt.Sprintf("Initializing dependencies...")

	repo := repository.NewAuthRepository(database)

	svc := service.NewAuthDomainService(repo)

	return controller.NewAuthenticationControllerInterface(svc)
}
