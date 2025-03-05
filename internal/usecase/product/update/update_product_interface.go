package usecase

import (
	"context"
)

type UpdateProductUseCase interface {
	UpdateProduct(ctx context.Context, customer UpdateProductInputDTO) (*UpdateProductOutputDTO, error)
}
