package portsrepository

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
)

// KitchenRepository defines the methods for interacting with the kitchen data.
type KitchenRepository interface {
	Create(ctx context.Context, kt *entity.Kitchen) error
	Update(ctx context.Context, kt *entity.Kitchen) error
	Find(ctx context.Context, id string) (*entity.Kitchen, error)
	FindAll(ctx context.Context) ([]*entity.Kitchen, error)
	Delete(ctx context.Context, id string) error
	FindByParams(ctx context.Context, params map[string]interface{}) ([]*entity.Kitchen, error)
	Monitor(ctx context.Context) ([]*entity.Kitchen, error)
}
