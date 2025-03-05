package usecase

import (
	"context"
)

type FindProductByIDUseCase interface {
	FindProductByID(ctx context.Context, id string) (*FindProductByIDOutputDTO, error)
}
