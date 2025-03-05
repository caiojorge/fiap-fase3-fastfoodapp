package usecase

import (
	"context"
)

type CreateOrderUseCase interface {
	CreateOrder(ctx context.Context, order *OrderCreateInputDTO) (*OrderCreateOutputDTO, error)
}
