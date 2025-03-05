package usecase

import (
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
	"github.com/stretchr/testify/assert"
)

func TestToEntity_ValidDTOWithMaskedCPF(t *testing.T) {
	// Arrange
	dto := &OrderCreateInputDTO{
		CustomerCPF: "167.132.236-30",
		Items: []*OrderItemCreateInputDTO{
			{
				ProductID: "1",
				Quantity:  1,
			},
		},
	}
	wrapper := &CreateOrderWrapper{
		dto: dto,
	}
	// Act
	order, err := wrapper.ToEntity()

	// Assert
	assert.NotNil(t, order)
	assert.Nil(t, err)

	expectedCPF := "16713223630"
	assert.Equal(t, expectedCPF, order.CustomerCPF)
	assert.Len(t, order.Items, 1)
	assert.Equal(t, "1", order.Items[0].ProductID)
	assert.Equal(t, 1, order.Items[0].Quantity)

}

func TestCPFValidator_IsValid(t *testing.T) {
	// Arrange
	v := validator.NewCPFValidator()

	// Act
	err := v.Validate("167.132.236-30")

	// Assert
	assert.Nil(t, err)
}
