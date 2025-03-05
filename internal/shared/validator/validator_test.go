package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailValidatorTrue(t *testing.T) {
	email := "email@email.com"

	v := EmailValidator{}
	result := v.IsValid(email)
	assert.True(t, result)
}

func TestEmailValidatorFalse(t *testing.T) {
	email := "email.email.com"

	v := EmailValidator{}
	result := v.IsValid(email)
	assert.False(t, result)
}
