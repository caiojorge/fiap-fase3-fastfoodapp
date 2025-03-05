package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	customerrors "github.com/caiojorge/fiap-challenge-ddd/internal/shared/error"
)

type CheckPaymentUseCase struct {
	checkoutRepo ports.CheckoutRepository
	orderRepo    ports.OrderRepository
}

func NewCheckPaymentUseCase(checkoutRepo ports.CheckoutRepository, orderRepo ports.OrderRepository) *CheckPaymentUseCase {
	return &CheckPaymentUseCase{
		checkoutRepo: checkoutRepo,
		orderRepo:    orderRepo,
	}
}

// CheckPayment verifica se a ordem em questão foi paga; é um mix de checkout com ordem
func (cr *CheckPaymentUseCase) CheckPayment(ctx context.Context, orderID string) (*CheckPaymentOutputDTO, error) {

	checkout, err := cr.checkoutRepo.FindbyOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	order, err := cr.orderRepo.Find(ctx, orderID)
	if err != nil {
		return nil, err
	}

	// nesse caso faz sentido retornar um erro se não encontrar o pedido ou o checkout
	if order == nil {
		return nil, customerrors.ErrOrderNotFound
	}

	outputs := &CheckPaymentOutputDTO{
		OrderID:              checkout.OrderID,
		GatewayTransactionID: checkout.Gateway.GatewayTransactionID,
		Status:               order.Status.Name,
		PaymentApproved:      order.IsPaymentApproved(),
	}

	return outputs, nil
}
