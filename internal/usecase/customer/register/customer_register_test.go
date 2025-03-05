package usecase

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRegister(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocksrepository.NewMockCustomerRepository(ctrl)

	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Customer{})).
		Return(nil)

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := entity.NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	mockRepo.EXPECT().
		Find(gomock.Any(), customer.CPF.Value).
		Return(nil, nil).
		Times(1)

	inputDto := CustomerRegisterInputDTO{
		CPF:   customer.CPF.Value,
		Name:  customer.Name,
		Email: customer.Email,
	}

	register := NewCustomerRegister(mockRepo)
	assert.NotNil(t, register)

	err = register.RegisterCustomer(context.Background(), inputDto)
	assert.Nil(t, err)

}
