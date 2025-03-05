package converter

import (
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	"github.com/stretchr/testify/assert"
)

func TestOrderConverter_FromEntity(t *testing.T) {
	converter := NewOrderConverter()

	entityOrder := &entity.Order{
		ID:          "order123",
		Total:       100.0,
		Status:      *entity.NewStatus("pending"),
		CustomerCPF: "12345678900",
		Items: []*entity.OrderItem{
			{ID: "item1", ProductID: "product1", Quantity: 2, Price: 50.0, Status: "new"},
		},
	}

	modelOrder := converter.FromEntity(entityOrder)

	assert.Equal(t, entityOrder.ID, modelOrder.ID)
	assert.Equal(t, entityOrder.Total, modelOrder.Total)
	assert.Equal(t, entityOrder.Status.Name, modelOrder.Status)
	assert.Equal(t, entityOrder.CustomerCPF, *modelOrder.CustomerCPF)
	assert.Len(t, modelOrder.Items, len(entityOrder.Items))

	for i, item := range modelOrder.Items {
		entityItem := entityOrder.Items[i]
		assert.Equal(t, entityItem.ID, item.ID)
		assert.Equal(t, entityItem.ProductID, item.ProductID)
		assert.Equal(t, entityItem.Quantity, item.Quantity)
		assert.Equal(t, entityItem.Price, item.Price)
		assert.Equal(t, entityItem.Status, item.Status)
	}
}

func TestOrderConverter_ToEntity(t *testing.T) {
	customerCPF := "12345678900"
	status := "pending"
	converter := NewOrderConverter()
	modelOrder := &model.Order{
		ID:          "order123",
		Total:       100.0,
		Status:      status,
		CustomerCPF: &customerCPF,
		Items: []*model.OrderItem{
			{ID: "item1", ProductID: "product1", Quantity: 2, Price: 50.0, Status: "new"},
		},
	}

	entityOrder := converter.ToEntity(modelOrder)

	assert.Equal(t, modelOrder.ID, entityOrder.ID)
	assert.Equal(t, modelOrder.Total, entityOrder.Total)
	assert.Equal(t, modelOrder.Status, entityOrder.Status.Name)
	assert.Equal(t, *modelOrder.CustomerCPF, entityOrder.CustomerCPF)
	assert.Len(t, entityOrder.Items, len(modelOrder.Items))

	for i, item := range entityOrder.Items {
		modelItem := modelOrder.Items[i]
		assert.Equal(t, modelItem.ID, item.ID)
		assert.Equal(t, modelItem.ProductID, item.ProductID)
		assert.Equal(t, modelItem.Quantity, item.Quantity)
		assert.Equal(t, modelItem.Price, item.Price)
		assert.Equal(t, modelItem.Status, item.Status)
	}
}
