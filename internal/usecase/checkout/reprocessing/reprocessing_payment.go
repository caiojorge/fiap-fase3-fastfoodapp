package usecase

import (
	"context"
	"strconv"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	portsservice "github.com/caiojorge/fiap-challenge-ddd/internal/domain/gateway"
	repository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	customerrors "github.com/caiojorge/fiap-challenge-ddd/internal/shared/error"
	"go.uber.org/zap"
)

const Approved = "approved"

type CheckoutReprocessingInputDTO struct {
	NotificationURL string `json:"notification_url"` // webhook para receber a confirmação do pagamento
	SponsorID       string `json:"sponsor_id"`
}

type CheckoutReprocessingOutputDTO struct {
	CheckoutID           string `json:"checkout_id" binding:"required"`
	OrderID              string `json:"order_id" binding:"required"`
	Status               string `json:"status" binding:"required"`
	GatewayTransactionID string `json:"gateway_transaction_id" binding:"required"`
	QRCode               string `json:"qrcode" binding:"required"`
}

type ICheckoutReprocessingUseCase interface {
	ReprocessPayment(ctx context.Context, input *CheckoutReprocessingInputDTO) (*[]CheckoutReprocessingOutputDTO, error)
}

type CheckoutReprocessingUseCase struct {
	orderRepository    repository.OrderRepository
	checkoutRepository repository.CheckoutRepository
	productRepository  repository.ProductRepository
	gatewayService     portsservice.GatewayTransactionService
	tm                 repository.TransactionManager
	logger             *zap.Logger
}

func NewCheckoutReprocessing(orderRepository repository.OrderRepository,
	checkoutRepository repository.CheckoutRepository,
	productRepository repository.ProductRepository,
	gatewayService portsservice.GatewayTransactionService,
	tm repository.TransactionManager,
	logger *zap.Logger) ICheckoutReprocessingUseCase {
	return &CheckoutReprocessingUseCase{
		orderRepository:    orderRepository,
		checkoutRepository: checkoutRepository,
		productRepository:  productRepository,
		gatewayService:     gatewayService,
		tm:                 tm,
		logger:             logger,
	}
}

func (cr *CheckoutReprocessingUseCase) ReprocessPayment(ctx context.Context, input *CheckoutReprocessingInputDTO) (*[]CheckoutReprocessingOutputDTO, error) {

	// 1. Verifica se existem ordens em checkout
	orders, err := cr.orderRepository.FindByParams(ctx, map[string]interface{}{"status": sharedconsts.OrderStatusCheckoutConfirmed})
	if err != nil {
		cr.logger.Error("Error to find order", zap.Error(err))
		return nil, err
	}

	if len(orders) == 0 {
		cr.logger.Error("Order not found")
		return nil, customerrors.ErrOrderNotFound
	}

	var outputs []CheckoutReprocessingOutputDTO
	// 2.Verifica se os ckeckouts existem
	for _, order := range orders {

		// só ordens com checkout confirmado

		checkout, err := cr.checkoutRepository.FindbyOrderID(ctx, order.ID)
		if err != nil {
			cr.logger.Error("Error to find checkout", zap.Error(err))
			return nil, err
		}
		if checkout == nil {
			cr.logger.Error("Checkout not found")
			return nil, customerrors.ErrCheckoutNotFound

		}
		// 3. Pega os produtos da ordem
		var products []*entity.Product

		for _, item := range order.Items {
			product, err := cr.productRepository.Find(ctx, item.ProductID)
			if err != nil {
				return nil, err
			}

			products = append(products, product)
		}

		sponsorID, err := strconv.Atoi(input.SponsorID)
		if err != nil {
			return nil, err
		}

		payment, err := cr.gatewayService.ConfirmPayment(ctx, checkout, order, products, input.NotificationURL, sponsorID)
		if err != nil {
			order.InformPaymentNotApproval()
			_ = cr.orderRepository.Update(ctx, order)
			return nil, err
		}

		// 4. Atualiza o checkout
		checkout.Reprocessing(payment.InStoreOrderID, payment.QrData)
		err = cr.checkoutRepository.Update(ctx, checkout)
		if err != nil {
			return nil, err
		}

		// 5. Retorna o output
		output := CheckoutReprocessingOutputDTO{
			CheckoutID:           checkout.ID,
			OrderID:              order.ID,
			Status:               order.Status.Name,
			GatewayTransactionID: payment.InStoreOrderID,
			QRCode:               payment.QrData,
		}

		outputs = append(outputs, output)

	}

	return &outputs, nil
}
