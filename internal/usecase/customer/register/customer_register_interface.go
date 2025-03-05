package usecase

import (
	"context"
)

type RegisterCustomerUseCase interface {
	RegisterCustomer(ctx context.Context, customer CustomerRegisterInputDTO) error
}
