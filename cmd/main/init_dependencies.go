package main

import (
	"database/sql"
	"fmt"
	"github.com/vitortenor/guardian/pkg/controller"
	"github.com/vitortenor/guardian/pkg/model/repository"
	"github.com/vitortenor/guardian/pkg/model/service"
)

func initAuthControllerDependencies(database *sql.DB) controller.AuthenticationControllerInterface {
	fmt.Sprintf("Initializing dependencies...")

	repo := repository.NewAuthRepository(database)

	svc := service.NewAuthDomainService(repo)

	return controller.NewAuthenticationControllerInterface(svc)
}
