package usecase

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocks "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	checkoutUseCase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/create"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCheckoutRepository := mocks.NewMockCheckoutRepository(ctrl)
	assert.NotNil(t, mockCheckoutRepository)
	mockOrderRepository := mocks.NewMockOrderRepository(ctrl)
	assert.NotNil(t, mockOrderRepository)
	mockProductRepository := mocks.NewMockProductRepository(ctrl)
	assert.NotNil(t, mockProductRepository)
	mockGatewayService := checkoutUseCase.NewMLFakePaymentService()
	assert.NotNil(t, mockGatewayService)
	mockKitchenRepository := mocks.NewMockKitchenRepository(ctrl)
	assert.NotNil(t, mockKitchenRepository)

	useCase := checkoutUseCase.NewCheckoutCreate(
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
	checkoutInput := &checkoutUseCase.CheckoutInputDTO{
		OrderID:         "order123",
		GatewayName:     "mercadopago", //TODO: colocar uma validaçao para o nome do gateway
		GatewayToken:    "01234567890",
		NotificationURL: "http://localhost:8080/checkout/notification", // TODO: essa URL deveria vir por parametro
		SponsorID:       1,                                             // TODO: descobrir o que é esse sponsorID
		DiscontCoupon:   0.0,                                           // Não é bem um cupom de desconto, mas sim um valor de desconto
	}

	status := entity.Status{
		Name: sharedconsts.OrderStatusConfirmed,
	}
	// Define entities for the mocks to return
	order := &entity.Order{
		ID:     "order123",
		Status: status,
		Items: []*entity.OrderItem{
			{ProductID: "prod123", Quantity: 1, Status: sharedconsts.OrderItemStatusConfirmed, Price: 100.0},
		},
	}

	// como criei a ordem na mão, preciso calcular o total dela.
	order.CalculateTotal()

	product := &entity.Product{
		ID:    "prod123",
		Name:  "Test Product",
		Price: 100.0,
	}

	// Set up mock expectations for a successful checkout

	mockOrderRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(order, nil).AnyTimes() // Order found and not paid

	mockOrderRepository.EXPECT().
		UpdateStatus(ctx, gomock.Any(), gomock.Any()).
		Return(nil) // Order found and not paid

	mockProductRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(product, nil) // Product found

	checkoutEntity, err := entity.NewCheckout(order.ID, checkoutInput.GatewayName, checkoutInput.GatewayToken, order.Total)
	assert.NoError(t, err)
	assert.NotNil(t, checkoutEntity)

	firstcall := mockCheckoutRepository.EXPECT().
		FindbyOrderID(ctx, gomock.Any()).
		Return(nil, nil) // No duplicate checkout found
	secondcall := mockCheckoutRepository.EXPECT().
		FindbyOrderID(ctx, gomock.Any()).
		Return(checkoutEntity, nil) // No duplicate checkout found

	gomock.InOrder(firstcall, secondcall)

	mockCheckoutRepository.EXPECT().
		Create(ctx, gomock.Any()).
		Return(nil) // Checkout creation successful

	// mockKitchenRepository.EXPECT().
	// 	Create(ctx, gomock.Any()).
	// 	Return(nil) // Kitchen entry creation successful

	// o checkout recede a ordem, que tem os itens e os produtos.
	// o payment é criado no padrão do gateway de pagamento, com a lista de produtos e a ordem.
	// o teste prova que o output do usecase recebe e retorna os dados solicitados.
	result, err := useCase.CreateCheckout(ctx, checkoutInput)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.ID)
	assert.NotNil(t, result.GatewayTransactionID)
	assert.NotNil(t, result.OrderID)
	assert.Equal(t, order.ID, result.OrderID)

	checkPayment := NewCheckPaymentUseCase(
		mockCheckoutRepository,
		mockOrderRepository,
	)

	checkedPayment, err := checkPayment.CheckPayment(ctx, result.OrderID)
	assert.NoError(t, err)
	assert.NotNil(t, checkedPayment)
	assert.NotNil(t, checkedPayment.Status)
	assert.NotNil(t, checkedPayment.GatewayTransactionID)
	assert.NotNil(t, checkedPayment.OrderID)
	assert.Equal(t, order.ID, checkedPayment.OrderID)
	assert.Equal(t, checkedPayment.PaymentApproved, false)

}
