package usecase

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
)

// DTO
type CustomerFindAllInputDTO struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CustomerFindAllInputDTO) ToEntity() *entity.Customer {
	return &entity.Customer{
		CPF: valueobject.CPF{
			Value: dto.CPF,
		},
		Name:  dto.Name,
		Email: dto.Email,
	}
}

type CustomerFindAllOutputDTO struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CustomerFindAllOutputDTO) FromEntity(customer entity.Customer) {
	dto.CPF = customer.CPF.Value
	dto.Name = customer.Name
	dto.Email = customer.Email
}

// ...
