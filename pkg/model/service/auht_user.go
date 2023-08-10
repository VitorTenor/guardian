package service

import (
	"github.com/vitortenor/guardian/pkg/config/logger"
	"github.com/vitortenor/guardian/pkg/config/rest_error"
	"github.com/vitortenor/guardian/pkg/model"
	"go.uber.org/zap"
)

func (ad *authDomainService) AuthUserServices(domainInterface model.AuthDomainInterface) (string, string, *rest_error.Err) {
	logger.Info("AuthUserServices",
		zap.String("journey", "AuthUserServices"),
	)

	user, _ := ad.repository.GetByEmail(domainInterface.GetEmail())

	return user.Name, user.Email, nil
}
