package service

import (
	"github.com/vitortenor/guardian/src/config/rest_error"
	"github.com/vitortenor/guardian/src/model"
	"github.com/vitortenor/guardian/src/model/repository"
)

func NewAuthDomainService(repo repository.AuthRepository) AuthDomainService {
	return &authDomainService{
		repository: repo,
	}
}

type authDomainService struct {
	repository repository.AuthRepository
}

type AuthDomainService interface {
	AuthUserServices(model.AuthDomainInterface) (string, string, *rest_error.Err)
	RenewTokenServices(string) (string, string, *rest_error.Err)
}
