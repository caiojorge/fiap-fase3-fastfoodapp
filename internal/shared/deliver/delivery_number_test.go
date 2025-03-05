package deliver

import (
	"regexp"
	"testing"
)

func TestGenerateDeliveryNumber(t *testing.T) {
	deliveryNumber := GenerateDeliveryNumber()

	match, _ := regexp.MatchString(`^[A-Z]{2}\d{3}$`, deliveryNumber)
	if !match {
		t.Errorf("Expected delivery number to match format 'AA999', but got %s", deliveryNumber)
	}

	deliveryNumber2 := GenerateDeliveryNumber()
	if deliveryNumber == deliveryNumber2 {
		t.Errorf("Expected delivery numbers to be unique, but got %s and %s", deliveryNumber, deliveryNumber2)
	}
}
