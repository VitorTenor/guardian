package service

import (
	"github.com/vitortenor/guardian/src/config/rest_error"
	"github.com/vitortenor/guardian/src/model"
)

func (ad *authDomainService) RenewTokenServices(refreshToken string) (string, string, *rest_error.Err) {

	domain := model.NewAuthDomain("", "", "")

	token, newRefreshToken, err := domain.VerifyAndRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	return token, newRefreshToken, nil
}
