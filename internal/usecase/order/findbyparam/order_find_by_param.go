package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type OrderFindByParamsUseCase struct {
	repository ports.OrderRepository
}

func NewOrderFindByParams(repository ports.OrderRepository) *OrderFindByParamsUseCase {
	return &OrderFindByParamsUseCase{
		repository: repository,
	}
}

// FindAllOrder busca todas as ordens
func (cr *OrderFindByParamsUseCase) FindOrdersByParams(ctx context.Context, params map[string]interface{}) ([]*OrderFindByParamOutputDTO, error) {

	orders, err := cr.repository.FindByParams(ctx, params)
	if err != nil {
		return nil, err
	}

	outputs := FromEntity(orders)

	return outputs, nil
}
