package usecase

import "github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"

// RegisterProductInputDTO para atender o use case de registro de produto
type RegisterProductInputDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

func (dto *RegisterProductInputDTO) ToEntity() *entity.Product {
	return &entity.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Category:    dto.Category,
		Price:       dto.Price,
	}
}

type RegisterProductOutputDTO struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

func (dto *RegisterProductOutputDTO) FromEntity(product entity.Product) {
	dto.ID = product.ID
	dto.Name = product.Name
	dto.Description = product.Description
	dto.Category = product.Category
	dto.Price = product.Price
}

// ...
