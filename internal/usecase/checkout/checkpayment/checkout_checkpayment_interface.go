package usecase

import "context"

type ICheckPaymentUseCase interface {
	CheckPayment(ctx context.Context, orderID string) (*CheckPaymentOutputDTO, error)
}
