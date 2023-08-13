package repository

import (
	"database/sql"
	"github.com/vitortenor/guardian/src/config/rest_error"
	"github.com/vitortenor/guardian/src/model/repository/entity"
)

type AuthRepository interface {
	GetByEmail(email string) (*entity.AuthEntity, *rest_error.Err)
}

func NewAuthRepository(databaseConnection *sql.DB) AuthRepository {
	return &authRepository{
		databaseConnection: databaseConnection,
	}
}

type authRepository struct {
	databaseConnection *sql.DB
}
