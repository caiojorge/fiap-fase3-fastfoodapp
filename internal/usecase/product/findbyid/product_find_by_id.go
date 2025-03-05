package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type ProductFindByIDUseCase struct {
	repository ports.ProductRepository
}

func NewProductFindByID(repository ports.ProductRepository) *ProductFindByIDUseCase {
	return &ProductFindByIDUseCase{
		repository: repository,
	}
}

func (cr *ProductFindByIDUseCase) FindProductByID(ctx context.Context, id string) (*FindProductByIDOutputDTO, error) {

	product, err := cr.repository.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	productDTO := &FindProductByIDOutputDTO{}
	productDTO.FromEntity(*product)

	return productDTO, nil
}
