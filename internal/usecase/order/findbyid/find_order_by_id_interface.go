package usecase

import (
	"context"
)

type FindOrderByIDUseCase interface {
	FindOrderByID(ctx context.Context, id string) (*OrderFindByIdOutputDTO, error)
}
