package usecase

import (
	"context"
	"fmt"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type CustomerUpdateUseCase struct {
	repository portsrepository.CustomerRepository
}

func NewCustomerUpdate(repository portsrepository.CustomerRepository) *CustomerUpdateUseCase {
	return &CustomerUpdateUseCase{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerUpdateUseCase) UpdateCustomer(ctx context.Context, customer CustomerUpdateInputDTO) error {
	entity := customer.ToEntity()
	err := entity.Validate()
	if err != nil {
		return err
	}

	fmt.Println("usecase: verifica se o cliente existe: " + customer.CPF)
	c, err := cr.repository.Find(ctx, customer.CPF)
	if err != nil {
		fmt.Println("usecase: err: " + err.Error())
		return err
	}

	if c == nil {
		return fmt.Errorf("customer not found")
	}

	fmt.Println("usecase: atualizando cliente: " + customer.CPF)

	// Cria o cliente
	err = cr.repository.Update(ctx, customer.ToEntity())
	if err != nil {
		return err
	}

	return nil
}
