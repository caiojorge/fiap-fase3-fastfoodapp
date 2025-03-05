package repositorygorm

import (
	"context"

	"gorm.io/gorm"
)

type txKeyType string

const txKey txKeyType = "tx"

type GormTransactionManager struct {
	DB *gorm.DB
}

func NewGormTransactionManager(db *gorm.DB) *GormTransactionManager {
	return &GormTransactionManager{DB: db}
}

func (tm *GormTransactionManager) RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	// Inicia a transação
	tx := tm.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Substitui o contexto pelo contexto da transação
	ctxWithTx := context.WithValue(ctx, txKey, tx)

	// Executa a função com o contexto da transação
	if err := fn(ctxWithTx); err != nil {
		tx.Rollback() // Faz rollback em caso de erro
		return err
	}

	// Commit se tudo der certo
	return tx.Commit().Error
}
