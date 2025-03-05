package usecase

import (
	"context"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type CustomerFindAllUseCase struct {
	repository portsrepository.CustomerRepository
}

func NewCustomerFindAll(repository portsrepository.CustomerRepository) *CustomerFindAllUseCase {
	return &CustomerFindAllUseCase{
		repository: repository,
	}
}

func (cr *CustomerFindAllUseCase) FindAllCustomers(ctx context.Context) ([]*CustomerFindAllOutputDTO, error) {

	customers, err := cr.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var output []*CustomerFindAllOutputDTO

	for _, customer := range customers {
		dto := &CustomerFindAllOutputDTO{}
		dto.FromEntity(*customer)
		output = append(output, dto)
	}

	return output, nil
}
