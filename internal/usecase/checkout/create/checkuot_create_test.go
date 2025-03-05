package usecase

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocks "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCheckout(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCheckoutRepository := mocks.NewMockCheckoutRepository(ctrl)
	assert.NotNil(t, mockCheckoutRepository)
	mockOrderRepository := mocks.NewMockOrderRepository(ctrl)
	assert.NotNil(t, mockOrderRepository)
	mockProductRepository := mocks.NewMockProductRepository(ctrl)
	assert.NotNil(t, mockProductRepository)
	mockGatewayService := NewMLFakePaymentService()
	assert.NotNil(t, mockGatewayService)
	mockKitchenRepository := mocks.NewMockKitchenRepository(ctrl)
	assert.NotNil(t, mockKitchenRepository)

	useCase := NewCheckoutCreate(
		mockOrderRepository,
		mockCheckoutRepository,
		mockGatewayService,
		mockKitchenRepository,
		mockProductRepository,
	)
	assert.NotNil(t, useCase)

	ctx := context.Background()
	assert.NotNil(t, ctx)

	// Define input DTO
	checkoutInput := &CheckoutInputDTO{
		OrderID:         "order123",
		GatewayName:     "mercadopago", //TODO: colocar uma validaçao para o nome do gateway
		GatewayToken:    "01234567890",
		NotificationURL: "http://localhost:8080/checkout/notification", // TODO: essa URL deveria vir por parametro
		SponsorID:       1,                                             // TODO: descobrir o que é esse sponsorID
		DiscontCoupon:   0.0,                                           // Não é bem um cupom de desconto, mas sim um valor de desconto
	}

	status := entity.Status{
		Name: "confirmed",
	}

	// Define entities for the mocks to return
	order := &entity.Order{
		ID:     "order123",
		Status: status,
		Items: []*entity.OrderItem{
			{ProductID: "prod123", Quantity: 1, Status: sharedconsts.OrderItemStatusConfirmed, Price: 100.0},
		},
	}

	order.CalculateTotal()

	product := &entity.Product{
		ID:    "prod123",
		Name:  "Test Product",
		Price: 100.0,
	}

	// Set up mock expectations for a successful checkout
	mockCheckoutRepository.EXPECT().
		FindbyOrderID(ctx, "order123").
		Return(nil, nil) // No duplicate checkout found

	mockOrderRepository.EXPECT().
		Find(ctx, "order123").
		Return(order, nil) // Order found and not paid

	mockOrderRepository.EXPECT().
		UpdateStatus(ctx, gomock.Any(), gomock.Any()).
		Return(nil) // Order found and not paid

	mockProductRepository.EXPECT().
		Find(ctx, "prod123").
		Return(product, nil) // Product found

	mockCheckoutRepository.EXPECT().
		Create(ctx, gomock.Any()).
		Return(nil) // Checkout creation successful

	result, err := useCase.CreateCheckout(ctx, checkoutInput)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.ID)
	assert.NotNil(t, result.GatewayTransactionID)
	assert.NotNil(t, result.OrderID)
	assert.Equal(t, order.ID, result.OrderID)                                     // #3 Checkout Pedido que deverá receber os produtos solicitados e retornar à identificação do pedido.
	assert.Equal(t, order.Status.Name, sharedconsts.OrderStatusCheckoutConfirmed) // #4 Checkout O pagamento deve ser confirmado e o ID da transação do gateway deve ser retornado.
}
