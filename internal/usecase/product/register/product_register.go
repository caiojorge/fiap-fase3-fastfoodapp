package usecase

import (
	"context"
	"errors"
	"fmt"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
)

type ProductRegisterUseCase struct {
	repository portsrepository.ProductRepository
}

func NewProductRegister(repository portsrepository.ProductRepository) *ProductRegisterUseCase {
	return &ProductRegisterUseCase{
		repository: repository,
	}
}

// RegisterProduct registra um novo cliente.
func (cr *ProductRegisterUseCase) RegisterProduct(ctx context.Context, product *RegisterProductInputDTO) (*RegisterProductOutputDTO, error) {

	isOK := sharedconsts.IsCategoryValid(product.Category)
	if !isOK {
		return nil, errors.New("invalid category")
	}

	fmt.Println("usecase: verifica se o produto existe: " + product.Name)
	entityFound, err := cr.repository.FindByName(ctx, product.Name)
	if err != nil && err.Error() != "product not found" {
		fmt.Println("usecase: err: " + err.Error())
		return nil, err
	}

	if entityFound != nil {
		fmt.Println("usecase: producto já existe: " + product.Name)
		return nil, errors.New("product already exists")
	}

	fmt.Println("usecase: Criando produto: " + product.Name)
	entity := product.ToEntity()

	entity.DefineID()
	entity.FormatCategory()
	err = entity.Validate()
	if err != nil {
		return nil, err
	}

	err = cr.repository.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	output := RegisterProductOutputDTO{}
	output.FromEntity(*entity)

	return &output, nil
}
