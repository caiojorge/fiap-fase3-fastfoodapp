package usecase

import (
	"context"

	repository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	customerrors "github.com/caiojorge/fiap-challenge-ddd/internal/shared/error"
	"go.uber.org/zap"
)

const Approved = "approved"

type CheckoutConfirmationInputDTO struct {
	OrderID string `json:"order_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}

type CheckoutConfirmationOutputDTO struct {
	CheckoutID           string `json:"checkout_id" binding:"required"`
	OrderID              string `json:"order_id" binding:"required"`
	Status               string `json:"status" binding:"required"`
	GatewayTransactionID string `json:"gateway_transaction_id" binding:"required"`
	QRCode               string `json:"qrcode" binding:"required"`
}

type ICheckoutConfirmationUseCase interface {
	ConfirmPayment(ctx context.Context, checkout *CheckoutConfirmationInputDTO) (*CheckoutConfirmationOutputDTO, error)
}

type CheckoutConfirmationUseCase struct {
	orderRepository    repository.OrderRepository
	checkoutRepository repository.CheckoutRepository
	tm                 repository.TransactionManager
	logger             *zap.Logger
}

func NewCheckoutConfirmation(orderRepository repository.OrderRepository,
	checkoutRepository repository.CheckoutRepository,
	tm repository.TransactionManager,
	logger *zap.Logger) *CheckoutConfirmationUseCase {
	return &CheckoutConfirmationUseCase{
		orderRepository:    orderRepository,
		checkoutRepository: checkoutRepository,
		tm:                 tm,
		logger:             logger,
	}
}

func (cr *CheckoutConfirmationUseCase) ConfirmPayment(ctx context.Context, input *CheckoutConfirmationInputDTO) (*CheckoutConfirmationOutputDTO, error) {

	// 1. Verificar se a ordem existe
	order, err := cr.orderRepository.Find(ctx, input.OrderID)
	if err != nil {
		cr.logger.Error("Error to find order", zap.Error(err))
		return nil, err
	}

	if order == nil {
		cr.logger.Error("Order not found")
		return nil, customerrors.ErrOrderNotFound
	}
	cr.logger.Info("Order found", zap.Any("order", order))

	// 2. Verificar se o checkout existe, com base no id da ordem
	checkout, err := cr.checkoutRepository.FindbyOrderID(ctx, input.OrderID)
	if err != nil {
		cr.logger.Error("Error to find checkout", zap.Error(err))
		return nil, err
	}

	if checkout == nil {
		return nil, customerrors.ErrCheckoutNotFound
	}

	cr.logger.Info("Checkout found", zap.Any("checkout", checkout))

	if input.Status == Approved {
		// 3. Mudar status da ordem
		order.ConfirmPayment()
		// 4. Mudar status do checkout
		//checkout.ConfirmPayment()

		cr.logger.Info("Order and Checkout confirmed")
	} else {
		order.InformPaymentNotApproval()
		//checkout.InformPaymentNotApproval()
		cr.logger.Info("Order and Checkout not confirmed")
	}

	cr.logger.Info("Iniciando a transação", zap.Any("order", order), zap.Any("checkout", checkout))
	// Abre uma transação para salvar ordem e checkout
	err = cr.tm.RunInTransaction(ctx, func(ctx context.Context) error {
		// 5. Salvar novo status da ordem
		//err := cr.orderRepository.Update(ctx, order)
		err := cr.orderRepository.UpdateStatus(ctx, order.ID, order.Status.Name)
		if err != nil {
			cr.logger.Error("Error to update order", zap.Error(err))
			return err
		}

		// 6. Salvar novo status do checkout
		// err = cr.checkoutRepository.UpdateStatus(ctx, checkout.ID, checkout.Status)
		// if err != nil {
		// 	cr.logger.Error("Error to update checkout", zap.Error(err))
		// 	return err
		// }

		return nil
	})
	if err != nil {
		cr.logger.Error("Error to run transaction", zap.Error(err))
		return nil, err
	}

	cr.logger.Info("Finalizando a transação", zap.Any("order", order), zap.Any("checkout", checkout))

	output := &CheckoutConfirmationOutputDTO{
		CheckoutID:           checkout.ID,
		OrderID:              order.ID,
		Status:               order.Status.Name,
		GatewayTransactionID: checkout.Gateway.GatewayTransactionID,
		QRCode:               checkout.QRCode,
	}

	return output, nil
}
