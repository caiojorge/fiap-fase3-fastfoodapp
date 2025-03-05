package usecase

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	shared "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/shared"
)

// FromEntity converts an entity.Order to an OrderFindByIdOutputDTO
func FromEntity(orders []*entity.Order) []*OrderFindByParamOutputDTO {

	outputList := []*OrderFindByParamOutputDTO{}
	itemsDto := []*shared.OrderItemDTO{}

	for _, order := range orders {
		output := OrderFindByParamOutputDTO{
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

	return outputList

}
