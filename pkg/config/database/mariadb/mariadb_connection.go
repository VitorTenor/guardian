package mariadb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vitortenor/guardian/pkg/config/logger"
	"go.uber.org/zap"
	"os"
)

var (
	DATABASE_CONNECTION = "DATABASE_CONNECTION"
)

func NewMariaDBConnection() (*sql.DB, error) {
	logger.Info("Initializing mariadb connection",
		zap.String("journey", "mariadbConnection"),
	)

	db, err := sql.Open("mysql", os.Getenv(DATABASE_CONNECTION))
	if err != nil {
		logger.Info("Error when trying to connect with database",
			zap.String("journey", "mariadbConnection"),
		)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Info("Error when trying to connect with database",
			zap.String("journey", "mariadbConnection"),
		)
		return nil, err
	}

	logger.Info("Mariadb connection initialized with success",
		zap.String("journey", "mariadbConnection"),
	)
	return db, nil
}
