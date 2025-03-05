package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type OrderFindByIDUseCase struct {
	repository ports.OrderRepository
}

func NewOrderFindByID(repository ports.OrderRepository) *OrderFindByIDUseCase {
	return &OrderFindByIDUseCase{
		repository: repository,
	}
}

func (cr *OrderFindByIDUseCase) FindOrderByID(ctx context.Context, id string) (*OrderFindByIdOutputDTO, error) {

	order, err := cr.repository.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	output := FromEntity(*order)

	return &output, nil
}
