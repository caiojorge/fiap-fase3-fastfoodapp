package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
)

type CustomerFindByCPFUseCase struct {
	repository ports.CustomerRepository
}

func NewCustomerFindByCPF(repository ports.CustomerRepository) *CustomerFindByCPFUseCase {
	return &CustomerFindByCPFUseCase{
		repository: repository,
	}
}

// RegisterCustomer registra um novo cliente.
func (cr *CustomerFindByCPFUseCase) FindCustomerByCPF(ctx context.Context, cpf string) (*CustomerFindByCpfOutputDTO, error) {
	_, err := valueobject.NewCPF(cpf)
	if err != nil {
		return nil, err
	}

	customer, err := cr.repository.Find(ctx, cpf)
	if err != nil {
		return nil, err
	}

	dto := &CustomerFindByCpfOutputDTO{}
	dto.FromEntity(*customer)

	return dto, nil
}
