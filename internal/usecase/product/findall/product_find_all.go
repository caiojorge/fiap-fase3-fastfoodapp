package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type ProductFindAllUseCase struct {
	repository ports.ProductRepository
}

func NewProductFindAll(repository ports.ProductRepository) *ProductFindAllUseCase {
	return &ProductFindAllUseCase{
		repository: repository,
	}
}

// FindAllProducts busca todos os produtos.
func (cr *ProductFindAllUseCase) FindAllProducts(ctx context.Context) ([]*FindAllProductOutputDTO, error) {

	products, err := cr.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var output []*FindAllProductOutputDTO

	for _, product := range products {
		dto := &FindAllProductOutputDTO{}
		dto.FromEntity(*product)
		output = append(output, dto)
	}

	return output, nil
}
