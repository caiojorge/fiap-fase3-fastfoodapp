package usecase

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/jinzhu/copier"
)

// CreateOrderWrapper wraps the OrderCreateInputDTO to Order entity
// A ordem é um objeto complexo, então achei mais fácil criar um wrapper para fazer a conversão.
type CreateOrderWrapper struct {
	dto *OrderCreateInputDTO
}

func (co *CreateOrderWrapper) ToEntity() (*entity.Order, error) {
	var order entity.Order
	if err := copier.Copy(&order, &co.dto); err != nil {
		return nil, err
	}

	// just in case... remove a máscara do cpf
	order.RemoveMaksFromCPF()

	if err := order.Validate(); err != nil {
		return nil, err
	}

	return &order, nil

}
