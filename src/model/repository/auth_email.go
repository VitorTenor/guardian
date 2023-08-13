package repository

import (
	"github.com/vitortenor/guardian/src/config/logger"
	"github.com/vitortenor/guardian/src/config/rest_error"
	"github.com/vitortenor/guardian/src/model/repository/entity"
	"go.uber.org/zap"
)

func (ar *authRepository) GetByEmail(email string) (*entity.AuthEntity, *rest_error.Err) {
	logger.Info("Init GetByEmail Repository",
		zap.String("journey", "GetByEmail"),
		zap.String("email", email),
	)
	userEntity := &entity.AuthEntity{}

	err := ar.databaseConnection.QueryRow("SELECT id, email, password, name FROM users WHERE email = ?", email).Scan(&userEntity.ID, &userEntity.Email, &userEntity.Password, &userEntity.Name)
	if err != nil {

		switch err.Error() {
		case "sql: no rows in result set":
			logger.Info("User not found",
				zap.String("journey", "GetByEmail"),
				zap.String("email", email),
			)

			return nil, rest_error.NewNotFoundError("User not found")
		default:
			logger.Error("Error on get user by email",
				err,
				zap.String("journey", "GetByEmail"),
				zap.String("email", email),
			)

			return nil, rest_error.NewInternalServerError("Error on get user by email", err)
		}
	}

	logger.Info("GetByEmail Repository OK",
		zap.String("journey", "GetByEmail"),
		zap.String("email", email),
	)

	return userEntity, nil
}
