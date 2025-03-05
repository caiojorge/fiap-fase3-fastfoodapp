package customerrors

import "errors"

var (
	ErrOrderNotFound           = errors.New("order not found")
	ErrOrderIDIsRequired       = errors.New("order ID is required")
	ErrCheckoutNotFound        = errors.New("checkout not found")
	ErrFailedCheckPaymentError = errors.New("failed to check payment")
	// Add other custom errors as needed
)
