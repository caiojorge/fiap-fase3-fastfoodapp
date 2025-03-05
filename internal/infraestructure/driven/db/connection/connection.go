package setup

import (
	"errors"
	"os"

	infra "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/db"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Connection interface {
	Execute() (*gorm.DB, error)
}

type ConnectionGorm struct {
	logger       *zap.Logger
	dataBaseType string // mysql, sqlite e ou memory
}

func NewConnection(dataBaseType string, logger *zap.Logger) Connection {
	return &ConnectionGorm{dataBaseType: dataBaseType, logger: logger}
}

func (s *ConnectionGorm) Execute() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db := infra.NewDB(host, port, user, password, dbName)

	s.logger.Info("Database connection established")
	s.logger.Info(dbName, zap.String("host", host), zap.String("port", port), zap.String("user", user))
	// get a connection
	connection := db.GetConnection(s.dataBaseType)
	if connection == nil {
		return nil, errors.New("expected a non-nil MySQL connection, but got nil")
	}

	return connection, nil
}
