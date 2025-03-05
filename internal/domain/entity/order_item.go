package entity

import (
	"errors"

	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"
)

type OrderItem struct {
	ID        string
	ProductID string
	Quantity  int
	Price     float64
	Status    string
}

// NewOrderItem creates a new OrderItem
func NewOrderItem(productID string, quantity int, price float64) (*OrderItem, error) {
	item := OrderItem{
		ID:        sharedgenerator.NewIDGenerator(),
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
		Status:    sharedconsts.OrderItemStatusConfirmed,
	}

	err := item.Validate()
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (i *OrderItem) Confirm() {
	i.ID = sharedgenerator.NewIDGenerator()
	i.Status = sharedconsts.OrderItemStatusConfirmed

}

func (i *OrderItem) ConfirmPrice(price float64) {
	i.Price = price
}

func (i *OrderItem) Validate() error {
	if i.ProductID == "" {
		return errors.New("product id is required")
	}

	if i.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	if i.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	return nil
}

func (i *OrderItem) Cancel() {
	i.Status = sharedconsts.OrderItemStatusCanceled
}
