package usecase

import (
	"context"
)

type FindAllKitchenUseCase interface {
	FindAllKitchen(ctx context.Context) ([]*KitchenFindAllAOutputDTO, error)
}
