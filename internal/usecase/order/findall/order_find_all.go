package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	shared "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/shared"
)

type OrderFindAllUseCase struct {
	repository ports.OrderRepository
}

func NewOrderFindAll(repository ports.OrderRepository) *OrderFindAllUseCase {
	return &OrderFindAllUseCase{
		repository: repository,
	}
}

// FindAllOrder busca todas as ordens
func (cr *OrderFindAllUseCase) FindAllOrders(ctx context.Context) ([]*OrderFindAllOutputDTO, error) {

	orders, err := cr.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	outputList := []*OrderFindAllOutputDTO{}
	itemsDto := []*shared.OrderItemDTO{}

	for _, order := range orders {
		output := OrderFindAllOutputDTO{
			ID:          order.ID,
			CustomerCPF: order.CustomerCPF,
			Total:       order.Total,
			CreatedAt:   order.CreatedAt,
			Status:      order.Status.Name,
		}

		for _, item := range order.Items {
			itemDto := shared.OrderItemDTO{
				ID:        item.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
				Status:    item.Status,
			}
			itemsDto = append(itemsDto, &itemDto)
		}

		output.Items = itemsDto

		outputList = append(outputList, &output)

	}

	return outputList, nil
}
