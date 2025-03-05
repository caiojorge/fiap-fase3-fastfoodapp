package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPF(t *testing.T) {
	// Test cases for NewCPF function
	t.Run("NewCPF_ValidValue_ReturnsCPF", func(t *testing.T) {
		// Test logic here
		cpf, err := NewCPF("123.456.789-09")
		if err != nil {
			t.Errorf("NewCPF() error = %v, want nil", err)
		}
		if cpf == nil {
			t.Errorf("NewCPF() = %v, want CPF", cpf)
		}

	})

	t.Run("FormatCPF_ReturnsNonAndFormatted", func(t *testing.T) {
		// Test logic here
		cpf, err := NewCPF("123.456.789-09")
		assert.Nil(t, err)
		assert.NotNil(t, cpf)

		formattedCPF, err := cpf.Format()
		assert.Nil(t, err)
		assert.NotNil(t, formattedCPF)
		assert.Equal(t, "123.456.789-09", cpf.GetValue())

		nonFormattedCPF := cpf.RemoveFormat()
		assert.NotNil(t, nonFormattedCPF)
		assert.Equal(t, "12345678909", nonFormattedCPF)

	})

	t.Run("NewCPF_InvalidValue_ReturnsError", func(t *testing.T) {
		// Test logic here
		cpf, err := NewCPF("123.456.789-0")
		if err == nil {
			t.Errorf("NewCPF() error = nil, want error")
		}
		if cpf != nil {
			t.Errorf("NewCPF() = CPF, want nil")
		}

	})

	// Test cases for GetValue function
	t.Run("GetValue_ReturnsValue", func(t *testing.T) {
		// Test logic here
		cpf := &CPF{
			Value: "123.456.789-09",
		}
		if got := cpf.GetValue(); got != "123.456.789-09" {
			t.Errorf("CPF.GetValue() = %v, want 123.456.789-09", got)
		}

	})

	// Test cases for Validate function
	t.Run("Validate_ValidCPF_ReturnsNil", func(t *testing.T) {
		// Test logic here
		cpf := &CPF{
			Value: "123.456.789-09",
		}
		if err := cpf.Validate(); err != nil {
			t.Errorf("CPF.Validate() error = %v, want nil", err)
		}

	})

	t.Run("Validate_InvalidCPF_ReturnsError", func(t *testing.T) {
		// Test logic here
		cpf := &CPF{
			Value: "123.456.789-0",
		}
		if err := cpf.Validate(); err == nil {
			t.Errorf("CPF.Validate() error = nil, want error")
		}

	})

	// Test cases for allSameDigits function
	t.Run("Validate_AllDigitsSame_ReturnsTrue", func(t *testing.T) {
		// Test logic here

		cpf := &CPF{
			Value: "111.111.111-11",
		}

		if err := cpf.Validate(); err == nil {
			t.Errorf("CPF.Validate() error = nil, want error")
		}

	})

}
