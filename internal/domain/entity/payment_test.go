package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayment(t *testing.T) {

	checkout, err := NewCheckout("order123", "gatewayteste", "gatewaytoken1234567890", 100)
	assert.Nil(t, err)
	assert.NotNil(t, checkout)

	product, err := NewProduct("prod123", "product", "product category", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	orderItem, err := NewOrderItem(product.ID, 1, 100)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem)
	assert.NotEmpty(t, orderItem.ProductID)

	order, err := NewOrder("order123", []*OrderItem{orderItem})
	assert.Nil(t, err)
	assert.NotNil(t, order)
	//order.CalculateTotal()

	payment, err := NewPayment(
		*checkout,
		*order,
		[]*Product{product},
		"http://localhost:8080/checkout/notification",
		1,
	)
	assert.Nil(t, err)
	assert.NotNil(t, payment)

}
