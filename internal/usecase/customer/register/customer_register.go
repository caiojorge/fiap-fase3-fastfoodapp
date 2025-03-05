package usecase

import (
	"context"
	"errors"
	"fmt"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type CustomerRegisterUseCase struct {
	repository portsrepository.CustomerRepository
}

func NewCustomerRegister(repository portsrepository.CustomerRepository) *CustomerRegisterUseCase {
	return &CustomerRegisterUseCase{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerRegisterUseCase) RegisterCustomer(ctx context.Context, customer CustomerRegisterInputDTO) error {

	entity := customer.ToEntity()

	err := entity.Validate()
	if err != nil {
		return err
	}

	fmt.Println("usecase: verifica se o cliente existe: " + entity.GetCPF().Value)
	customerFound, err := cr.repository.Find(ctx, entity.GetCPF().Value)
	if err != nil && err.Error() != "customer not found" {
		fmt.Println("usecase: err: " + err.Error())
		return err
	}

	if customerFound != nil {
		fmt.Println("usecase: Cliente já existe: " + entity.GetCPF().Value)
		return errors.New("customer already exists")
	}

	fmt.Println("usecase: Criando cliente: " + entity.GetCPF().Value)
	// Cria o cliente
	err = cr.repository.Create(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}
