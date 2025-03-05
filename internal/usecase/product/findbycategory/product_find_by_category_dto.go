package usecase

import "github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"

// RegisterProductInputDTO para atender o use case de registro de produto
type FindProductByCategoryInputDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

func (dto *FindProductByCategoryInputDTO) ToEntity() *entity.Product {
	return &entity.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Category:    dto.Category,
		Price:       dto.Price,
	}
}

type FindProductByCategoryOutputDTO struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

func (dto *FindProductByCategoryOutputDTO) FromEntity(product entity.Product) {
	dto.ID = product.ID
	dto.Name = product.Name
	dto.Description = product.Description
	dto.Category = product.Category
	dto.Price = product.Price
}

// ...
