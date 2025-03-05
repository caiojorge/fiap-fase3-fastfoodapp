package usecase

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	shared "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/shared"
)

func FromEntity(order entity.Order) OrderFindByIdOutputDTO {

	output := OrderFindByIdOutputDTO{
		ID:          order.ID,
		CustomerCPF: order.CustomerCPF,
		Total:       order.Total,
		CreatedAt:   order.CreatedAt,
		Status:      order.Status.Name,
	}

	itemsDto := []*shared.OrderItemDTO{}

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

	return output

}
