package usecase

import (
	"context"
)

type FindCustomerByCPFUseCase interface {
	FindCustomerByCPF(ctx context.Context, cpf string) (*CustomerFindByCpfOutputDTO, error)
}
