package usecase

import (
	"context"
	"errors"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
)

type ProductUpdateUseCase struct {
	repository portsrepository.ProductRepository
}

func NewProductUpdate(repository portsrepository.ProductRepository) *ProductUpdateUseCase {
	return &ProductUpdateUseCase{
		repository: repository,
	}
}

// UpdateProduct atualiza um novo produto.
func (cr *ProductUpdateUseCase) UpdateProduct(ctx context.Context, product UpdateProductInputDTO) (*UpdateProductOutputDTO, error) {

	isOK := sharedconsts.IsCategoryValid(product.Category)
	if !isOK {
		return nil, errors.New("invalid category")
	}

	prd, err := cr.repository.Find(ctx, product.ID)
	if err != nil {
		return nil, err
	}

	if prd == nil {
		return nil, errors.New("product not found")
	}

	entity := product.ToEntity()
	err = entity.Validate()
	if err != nil {
		return nil, err
	}

	err = cr.repository.Update(ctx, product.ToEntity())
	if err != nil {
		return nil, err
	}

	output := &UpdateProductOutputDTO{}
	output.FromEntity(*product.ToEntity())

	return output, nil
}
