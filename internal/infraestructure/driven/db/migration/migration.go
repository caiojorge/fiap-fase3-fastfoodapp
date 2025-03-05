package migration

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	"gorm.io/gorm"
)

type Migration interface {
	Execute() error
}

type MigrationGorm struct {
	db *gorm.DB
}

func NewMigration(db *gorm.DB) Migration {
	return &MigrationGorm{db: db}
}

func (m *MigrationGorm) Execute() error {
	if err := m.db.AutoMigrate(
		&model.Customer{},
		&model.Product{},
		&model.Order{},
		&model.OrderItem{},
		&model.Checkout{},
		&model.Kitchen{},
	); err != nil {
		return err
	}
	return nil
}
