package model

import "github.com/vitortenor/guardian/src/config/rest_error"

type AuthDomainInterface interface {
	GetId() int
	SetId(int)
	GetEmail() string
	GetPassword() string
	GetName() string
	SetName(string)
	CheckPassword([]byte) *rest_error.Err
	GenerateTokens() (string, string, *rest_error.Err)
	VerifyAndRefreshToken(string) (string, string, *rest_error.Err)
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
