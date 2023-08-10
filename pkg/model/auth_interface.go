package model

import "github.com/vitortenor/guardian/pkg/config/rest_error"

type AuthDomainInterface interface {
	GetId() string
	SetId(string)
	GetEmail() string
	GetPassword() string
	GetName() string
	GenerateToken() (string, *rest_error.Err)
	EncryptPassword() *rest_error.Err
	CheckPassword(string) *rest_error.Err
}

func NewAuthDomain(email, password, name string) AuthDomainInterface {
	return &authDomain{
		email:    email,
		password: password,
		name:     name,
	}
}

func AuthLoginDomain(email string, password string) AuthDomainInterface {
	return &authDomain{
		email:    email,
		password: password,
	}
}
