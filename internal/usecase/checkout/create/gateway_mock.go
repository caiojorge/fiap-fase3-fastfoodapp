package usecase

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
)

type MLFakePaymentService struct {
}

func NewMLFakePaymentService() *MLFakePaymentService {
	return &MLFakePaymentService{}
}

// CreateCheckout creates a new checkout. This method should be implemented by the payment gateway.
func (p *MLFakePaymentService) ConfirmPayment(ctx context.Context, checkout *entity.Checkout, order *entity.Order, productList []*entity.Product, notificationURL string, sponsorID int) (*entity.Payment, error) {
	payment, err := entity.NewPayment(*checkout, *order, productList, notificationURL, sponsorID)
	if err != nil {
		return nil, err
	}

	// TODO: connectar no server de pagamento
	// enviar dados de pagamento para o gateway
	// tratar a resposta do gateway

	return payment, nil
}

// CancelTransaction cancels a transaction. This method should be implemented by the payment gateway.
func (p *MLFakePaymentService) CancelPayment(ctx context.Context, id string) error {
	return nil
}
