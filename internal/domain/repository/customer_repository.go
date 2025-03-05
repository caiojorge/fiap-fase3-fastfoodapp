package portsrepository

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
)

// CustomerRepository defines the methods for interacting with the customer data.
type CustomerRepository interface {
	Create(ctx context.Context, customer *entity.Customer) error
	Update(ctx context.Context, customer *entity.Customer) error
	Find(ctx context.Context, id string) (*entity.Customer, error)
	FindAll(ctx context.Context) ([]*entity.Customer, error)
}
