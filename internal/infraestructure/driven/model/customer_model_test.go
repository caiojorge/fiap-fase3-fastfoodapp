package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerWithError1(t *testing.T) {
	// nome inválido
	model := Customer{
		CPF:   "123.456.789-09",
		Name:  "J",
		Email: "email@email.com",
	}

	assert.NotNil(t, model.Validate())

	model = Customer{
		CPF:   "12345678909",
		Name:  "J",
		Email: "email@email.com",
	}

	assert.NotNil(t, model.Validate())

}

func TestCustomerWithError2(t *testing.T) {
	// email inválido
	model := Customer{
		CPF:   "123.456.789-09",
		Name:  "John Doe",
		Email: "email.email.com",
	}

	assert.NotNil(t, model.Validate())
}

func TestCustomerWithError3(t *testing.T) {
	// cpf inválido
	model := Customer{
		CPF:   "123.456.789-0",
		Name:  "John Doe",
		Email: "email@email.com",
	}

	err := model.Validate()

	assert.NotNil(t, err)
}

func TestCustomerWithError5(t *testing.T) {
	// cpf inválido
	model := Customer{
		CPF:   "",
		Name:  "John Doe",
		Email: "email@email.com",
	}

	err := model.Validate()

	assert.NotNil(t, err)
}

func TestCustomerWithNoError(t *testing.T) {
	// customer válido
	model := Customer{
		CPF:   "123.456.789-09",
		Name:  "John Doe",
		Email: "email@email.com",
	}

	assert.Nil(t, model.Validate())

}

func TestCustomerWithNoError2(t *testing.T) {
	// customer válido
	model := Customer{
		CPF:   "123.456.789-09",
		Name:  "",
		Email: "",
	}

	assert.Nil(t, model.Validate())

}
