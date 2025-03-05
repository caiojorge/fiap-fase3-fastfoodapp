package usecase

import (
	"context"
)

type RegisterProductUseCase interface {
	//RegisterProduct(ctx context.Context, customer *entity.Product) (string, error)
	RegisterProduct(ctx context.Context, product *RegisterProductInputDTO) (*RegisterProductOutputDTO, error)
}
