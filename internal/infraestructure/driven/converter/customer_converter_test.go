package converter_test

import (
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	"github.com/stretchr/testify/assert"
)

func TestFromEntity(t *testing.T) {
	// TODO: Add test cases for FromEntity function

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.NotNil(t, cpf)
	assert.Nil(t, err)
	customer, err := entity.NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.NotNil(t, customer)
	assert.Nil(t, err)

	model := converter.FromEntity(customer)
	assert.NotNil(t, model)
	assert.Equal(t, "12345678909", model.CPF)
	assert.Equal(t, customer.GetName(), model.Name)
	assert.Equal(t, customer.GetEmail(), model.Email)

	entity := converter.ToEntity(model)
	assert.NotNil(t, entity)
	assert.Equal(t, customer, entity)

}
