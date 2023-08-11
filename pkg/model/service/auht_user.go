package service

import (
	"github.com/vitortenor/guardian/pkg/config/logger"
	"github.com/vitortenor/guardian/pkg/config/rest_error"
	"github.com/vitortenor/guardian/pkg/model"
	"go.uber.org/zap"
)

func (ad *authDomainService) AuthUserServices(domain model.AuthDomainInterface) (string, string, *rest_error.Err) {
	logger.Info("Init AuthUserServices service",
		zap.String("journey", "AuthUserServices"),
	)

	user, err := ad.repository.GetByEmail(domain.GetEmail())
	if err != nil {
		logger.Info("Error on get user by email",
			zap.String("journey", "AuthUserServices"),
			zap.String("email", domain.GetEmail()),
			zap.Error(err),
		)
		return "", "", err
	}

	err = domain.CheckPassword(user.Password)
	if err != nil {
		logger.Error("Error on check password", err)
		return "", "", err
	}

	domain.SetId(user.ID)
	domain.SetName(user.Name)

	token, err := domain.GenerateToken()
	if err != nil {
		logger.Error("Error on generate token",
			err,
			zap.String("journey", "AuthUserServices"),
		)
		return "", "", err
	}

	logger.Info("AuthUserServices Service OK",
		zap.String("journey", "AuthUserServices"),
	)

	return token, "", nil
}
