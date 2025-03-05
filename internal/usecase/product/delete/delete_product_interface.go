package usecase

import (
	"context"
)

type DeleteProductUseCase interface {
	DeleteProduct(ctx context.Context, id string) error
}
