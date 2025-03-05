package portsrepository

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
)

// OrderRepository defines the methods for interacting with the product data.
type OrderRepository interface {
	Create(ctx context.Context, product *entity.Order) error
	Update(ctx context.Context, product *entity.Order) error
	Find(ctx context.Context, id string) (*entity.Order, error)
	FindByParams(ctx context.Context, params map[string]interface{}) ([]*entity.Order, error)
	FindAll(ctx context.Context) ([]*entity.Order, error)
	Delete(ctx context.Context, id string) error
	UpdateStatus(ctx context.Context, id string, status string) error
}
