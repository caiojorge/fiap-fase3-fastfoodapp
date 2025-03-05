package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatCPF(t *testing.T) {
	cpf := "12345678900"
	expected := "123.456.789-00"

	result, err := PutMaskOnCPF(cpf)
	assert.Nil(t, err)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestRemoveFormatCPF(t *testing.T) {
	cpf := "123.456.789-00"
	expected := "12345678900"

	result := RemoveMaskFromCPF(cpf)
	assert.NotNil(t, result)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
