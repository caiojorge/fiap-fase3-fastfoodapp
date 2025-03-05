package usecase

import (
	"context"
)

type FindAllOrderUseCase interface {
	FindAllOrders(ctx context.Context) ([]*OrderFindAllOutputDTO, error)
}
