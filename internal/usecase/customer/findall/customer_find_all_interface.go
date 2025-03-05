package usecase

import (
	"context"
)

type FindAllCustomersUseCase interface {
	FindAllCustomers(ctx context.Context) ([]*CustomerFindAllOutputDTO, error)
}
