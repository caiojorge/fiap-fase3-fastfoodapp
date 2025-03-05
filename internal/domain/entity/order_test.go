package entity

import (
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"
	"github.com/stretchr/testify/assert"
)

// Test Order

func TestOrder(t *testing.T) {
	// Customer
	cpf, err := valueobject.NewCPF("19528476562")
	assert.Nil(t, err)
	customer, err := NewCustomer(*cpf, "Caio", "email@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "Caio", customer.Name)

	// Product
	product, err := NewProduct("Lanche xpto", "Pão, carne e queijo", "lanches", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	product2, err := NewProduct("Coca Cola", "Água com gás e xarope de coca", "refrigerante", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product2)

	// Order
	//order := OrderInit(customer.GetCPF().Value)

	orderItem, err := NewOrderItem(product.GetID(), 1, product.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem)
	orderItem2, err := NewOrderItem(product2.GetID(), 1, product2.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem2)

	order, _ := NewOrder(customer.GetCPF().Value, []*OrderItem{orderItem, orderItem2})
	assert.Equal(t, "19528476562", order.CustomerCPF)
	assert.Equal(t, 2, len(order.Items))
	assert.Equal(t, 20.00, order.Total)

}

func TestOrderWithNoItens(t *testing.T) {
	// Customer
	cpf, err := valueobject.NewCPF("19528476562")
	assert.Nil(t, err)
	customer, err := NewCustomer(*cpf, "Caio", "email@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "Caio", customer.Name)

	// Product
	product, err := NewProduct("Lanche xpto", "Pão, carne e queijo", "lanches", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	product2, err := NewProduct("Coca Cola", "Água com gás e xarope de coca", "refrigerante", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product2)

	// Order
	order, err := NewOrder(customer.GetCPF().Value, []*OrderItem{})
	assert.NotNil(t, err)
	assert.Nil(t, order)

}

func TestOrderWithNoCustomer(t *testing.T) {

	// Product
	product, err := NewProduct("Lanche xpto", "Pão, carne e queijo", "lanches", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	product2, err := NewProduct("Coca Cola", "Água com gás e xarope de coca", "refrigerante", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product2)

	orderItem, err := NewOrderItem(product.GetID(), 1, product.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem)
	orderItem2, err := NewOrderItem(product2.GetID(), 1, product2.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem2)

	// Order
	order, err := NewOrder("", []*OrderItem{orderItem, orderItem2})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(order.Items))
	assert.Equal(t, 20.00, order.Total)

}

func TestOrderWithNoRegistration(t *testing.T) {
	// Customer
	cpf, err := valueobject.NewCPF("19528476562")
	assert.Nil(t, err)
	customer, err := NewCustomerWithCPFOnly(cpf)
	assert.Nil(t, err)

	// Product
	product, err := NewProduct("Lanche xpto", "Pão, carne e queijo", "lanches", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	product2, err := NewProduct("Coca Cola", "Água com gás e xarope de coca", "refrigerante", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product2)

	orderItem, err := NewOrderItem(product.GetID(), 1, product.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem)

	orderItem2, err := NewOrderItem(product2.GetID(), 1, product2.Price)
	assert.Nil(t, err)
	assert.NotNil(t, orderItem2)

	// cancelando o item2, então o valor do pedido deveria ser 10
	orderItem2.Cancel()

	// Order
	order, err := NewOrder(customer.GetCPF().Value, []*OrderItem{orderItem, orderItem2})
	assert.Nil(t, err)
	assert.Equal(t, "19528476562", order.CustomerCPF)

	assert.Equal(t, 2, len(order.Items))
	assert.Equal(t, 10.00, order.Total)

	assert.Equal(t, sharedconsts.OrderStatusConfirmed, order.Status.Name)

}

func TestConfirmedOrder(t *testing.T) {
	cpf := "75419654059"

	lanche := Product{
		ID:          sharedgenerator.NewIDGenerator(),
		Name:        "Burger Kong",
		Description: "Pão, carne e queijo",
		Category:    "lanches",
		Price:       50.0,
	}

	refri := Product{
		ID:          sharedgenerator.NewIDGenerator(),
		Name:        "Pepsicola",
		Description: "Peptococa",
		Category:    "refrigerante",
		Price:       10.0,
	}

	item := OrderItem{
		ProductID: lanche.ID,
		Quantity:  1,
		Price:     lanche.Price,
	}

	item2 := OrderItem{
		ProductID: refri.ID,
		Quantity:  1,
		Price:     lanche.Price,
	}

	order := Order{
		Items:       []*OrderItem{&item, &item2},
		CustomerCPF: cpf,
	}

	err := order.Confirm()
	assert.Nil(t, err)

	assert.Equal(t, sharedconsts.OrderStatusConfirmed, order.Status.Name)

}
