package converter

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
)

type OrderConverter struct{}

func NewOrderConverter() *OrderConverter {
	return &OrderConverter{}
}

func (pc *OrderConverter) FromEntity(entity *entity.Order) *model.Order {
	return &model.Order{
		ID:             entity.ID,
		Items:          pc.fromEntityItems(entity.Items),
		Total:          entity.Total,
		Status:         entity.Status.Name,
		CustomerCPF:    &entity.CustomerCPF,
		CreatedAt:      entity.CreatedAt,
		DeliveryNumber: entity.DeliveryNumber,
	}
}

func (pc *OrderConverter) ToEntity(model *model.Order) *entity.Order {

	status := entity.Status{
		Name: model.Status,
	}

	return &entity.Order{
		ID:             model.ID,
		Items:          pc.toEntityItems(model.Items),
		Total:          model.Total,
		Status:         status,
		CustomerCPF:    pc.checkCustomerCPF(model.CustomerCPF),
		CreatedAt:      model.CreatedAt,
		DeliveryNumber: model.DeliveryNumber,
	}
}

func (pc *OrderConverter) checkCustomerCPF(cpf *string) string {
	if cpf == nil {
		return ""
	}
	return *cpf

}

func (pc *OrderConverter) fromEntityItems(items []*entity.OrderItem) []*model.OrderItem {
	modelItems := make([]*model.OrderItem, len(items))
	for i, item := range items {
		modelItems[i] = &model.OrderItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Status:    item.Status,
		}
	}
	return modelItems
}

func (pc *OrderConverter) toEntityItems(items []*model.OrderItem) []*entity.OrderItem {
	entityItems := make([]*entity.OrderItem, len(items))
	for i, item := range items {
		entityItems[i] = &entity.OrderItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Status:    item.Status,
		}
	}
	return entityItems
}
