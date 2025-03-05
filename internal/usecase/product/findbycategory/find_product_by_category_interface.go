package usecase

import (
	"context"
)

type FindProductByCategoryUseCase interface {
	FindProductByCategory(ctx context.Context, category string) ([]*FindProductByCategoryOutputDTO, error)
}
