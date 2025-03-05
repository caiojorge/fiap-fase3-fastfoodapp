package usecase

import (
	"context"
)

// ICreateCheckoutUseCase is the interface that wraps the CreateCheckout method.
type ICreateCheckoutUseCase interface {
	CreateCheckout(ctx context.Context, checkout *CheckoutInputDTO) (*CheckoutOutputDTO, error)
}
