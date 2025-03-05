package usecase

import "github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"

// RegisterProductInputDTO para atender o use case de registro de produto
type DeleteProductInputDTO struct {
	ID string `json:"id"`
}

type DeleteProductOutputDTO struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

func (dto *DeleteProductOutputDTO) FromEntity(product entity.Product) {
	dto.ID = product.ID
	dto.Name = product.Name
	dto.Description = product.Description
	dto.Category = product.Category
	dto.Price = product.Price
}

// ...
