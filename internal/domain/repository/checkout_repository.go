package portsrepository

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
)

// CheckoutRepository defines the methods for interacting with the product data.
type CheckoutRepository interface {
	Create(ctx context.Context, checkout *entity.Checkout) error
	Update(ctx context.Context, checkout *entity.Checkout) error
	Find(ctx context.Context, id string) (*entity.Checkout, error)
	FindAll(ctx context.Context) ([]*entity.Checkout, error)
	Delete(ctx context.Context, id string) error
	FindbyOrderID(ctx context.Context, id string) (*entity.Checkout, error)
	UpdateStatus(ctx context.Context, id string, status string) error
}
