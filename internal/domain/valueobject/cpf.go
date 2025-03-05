package valueobject

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
)

type CPF struct {
	Value string
}

func NewCPF(value string) (*CPF, error) {

	cpf := &CPF{
		Value: value,
	}

	err := cpf.Validate()
	if err != nil {
		return nil, err
	}

	return cpf, nil
}

func (c *CPF) GetValue() string {
	return c.Value
}

func (c *CPF) Validate() error {
	cpf := c.Value

	cpfValidator := validator.CPFValidator{}

	return cpfValidator.Validate(cpf)
}

func (c *CPF) Format() (string, error) {
	cpf := c.Value

	cpfFormatter, err := formatter.PutMaskOnCPF(cpf)
	if err != nil {
		return "", err
	}

	return cpfFormatter, nil
}

func (c *CPF) RemoveFormat() string {
	cpf := c.Value

	return formatter.RemoveMaskFromCPF(cpf)
}
