package usecase

import (
	"context"
)

type UpdateCustomerUseCase interface {
	UpdateCustomer(ctx context.Context, customer CustomerUpdateInputDTO) error
}
