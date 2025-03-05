package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocks "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/create"
	"github.com/gin-gonic/gin"
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
	mockGatewayService := usecase.NewMLFakePaymentService()
	assert.NotNil(t, mockGatewayService)
	mockKitchenRepository := mocks.NewMockKitchenRepository(ctrl)
	assert.NotNil(t, mockKitchenRepository)

	useCase := usecase.NewCheckoutCreate(
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
	checkoutInput := &usecase.CheckoutInputDTO{
		OrderID:         "order123",
		GatewayName:     "mercadopago", //TODO: colocar uma validaçao para o nome do gateway
		GatewayToken:    "01234567890",
		NotificationURL: "http://localhost:8080/checkout/notification", // TODO: essa URL deveria vir por parametro
		SponsorID:       1,                                             // TODO: descobrir o que é esse sponsorID
		DiscontCoupon:   0.0,                                           // Não é bem um cupom de desconto, mas sim um valor de desconto
	}

	status := &entity.Status{
		Name: sharedconsts.OrderStatusConfirmed,
	}

	// Define entities for the mocks to return
	order := &entity.Order{
		ID:     "order123",
		Status: *status,
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
		Return(nil, nil)

	mockOrderRepository.EXPECT().
		Find(ctx, "order123").
		Return(order, nil)

	// mockOrderRepository.EXPECT().
	// 	UpdateStatus(ctx, gomock.Any(), gomock.Any()).
	// 	Return(nil) // Order found and not paid

	mockOrderRepository.EXPECT().
		UpdateStatus(ctx, gomock.Any(), gomock.Any()).
		Return(nil) // Order found and not paid

	mockProductRepository.EXPECT().
		Find(ctx, "prod123").
		Return(product, nil) // Product found

	mockCheckoutRepository.EXPECT().
		Create(ctx, gomock.Any()).
		Return(nil)

	controller := NewCreateCheckoutController(context.Background(), useCase)

	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register the handler
	r.POST("/checkout", controller.PostCreateCheckout)

	jsonData, err := json.MarshalIndent(checkoutInput, "", "  ")
	assert.NoError(t, err)

	// Create a JSON body
	requestBody := bytes.NewBuffer(jsonData)

	// Create the HTTP request with JSON body
	req, err := http.NewRequest("POST", "/checkout", requestBody)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code, "Expected response code to be 200")

}
