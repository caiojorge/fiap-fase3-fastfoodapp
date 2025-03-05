package usecase

import (
	"context"
)

type FindAllProductsUseCase interface {
	FindAllProducts(ctx context.Context) ([]*FindAllProductOutputDTO, error)
}
