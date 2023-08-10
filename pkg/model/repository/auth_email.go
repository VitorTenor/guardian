package repository

import (
	"github.com/vitortenor/guardian/pkg/config/logger"
	"github.com/vitortenor/guardian/pkg/config/rest_error"
	"github.com/vitortenor/guardian/pkg/model/repository/entity"
	"go.uber.org/zap"
)

func (ar *authRepository) GetByEmail(email string) (*entity.AuthEntity, *rest_error.Err) {
	logger.Info("GetByEmail",
		zap.String("journey", "GetByEmail"),
		zap.String("email", email),
	)
	userEntity := &entity.AuthEntity{}

	err := ar.databaseConnection.QueryRow("SELECT id, email, password, name FROM users WHERE email = ?", email).Scan(&userEntity.ID, &userEntity.Email, &userEntity.Password, &userEntity.Name)
	if err != nil {
		logger.Error("Error on get user by email", err)
		return nil, nil
	}

	return userEntity, nil
}
