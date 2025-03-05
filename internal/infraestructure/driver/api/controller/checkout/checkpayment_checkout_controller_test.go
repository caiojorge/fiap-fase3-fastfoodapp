package controller

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocks "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/checkpayment"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCheckPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gin.SetMode(gin.TestMode) // Modo de teste para o Gin

	mockCheckoutRepository := mocks.NewMockCheckoutRepository(ctrl)
	assert.NotNil(t, mockCheckoutRepository)
	mockOrderRepository := mocks.NewMockOrderRepository(ctrl)
	assert.NotNil(t, mockOrderRepository)

	ctx := context.Background()

	t.Run("should return 200", func(t *testing.T) {

		checkout := &entity.Checkout{
			ID:        "checkout123",
			OrderID:   "order123",
			Gateway:   valueobject.Gateway{GatewayName: "mercadopago", GatewayTransactionID: "01234567890"},
			Total:     100.0,
			CreatedAt: time.Now(),
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

		// Set up mock expectations for a successful checkout
		mockCheckoutRepository.EXPECT().
			FindbyOrderID(ctx, "order123").
			Return(checkout, nil)

		mockOrderRepository.EXPECT().
			Find(ctx, "order123").
			Return(order, nil)

		mockUseCase := usecase.NewCheckPaymentUseCase(mockCheckoutRepository, mockOrderRepository)

		router := gin.Default()

		controller := &CheckPaymentCheckoutController{
			usecase: mockUseCase,
			ctx:     ctx,
		}

		router.GET("/checkout/check/:id", controller.GetCheckPaymentCheckout)

		// Simular requisição sem ID
		req, _ := http.NewRequest(http.MethodGet, "/checkout/check/order123", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		//assert.JSONEq(t, `{"error":"Order ID is required"}`, resp.Body.String())
	})

	t.Run("should return 404 if order not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := usecase.NewCheckPaymentUseCase(mockCheckoutRepository, mockOrderRepository)

		mockCheckoutRepository.EXPECT().
			FindbyOrderID(ctx, gomock.Any()).
			Return(nil, nil)

		mockOrderRepository.EXPECT().
			Find(ctx, gomock.Any()).
			Return(nil, nil)

		router := gin.Default()
		controller := &CheckPaymentCheckoutController{
			usecase: mockUseCase,
			ctx:     ctx,
		}

		router.GET("/checkout/check/:id", controller.GetCheckPaymentCheckout)

		req, _ := http.NewRequest(http.MethodGet, "/checkout/check/123", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		assert.JSONEq(t, `{"error":"order not found"}`, resp.Body.String())
	})

	t.Run("should return 500 if use case returns unexpected error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := usecase.NewCheckPaymentUseCase(mockCheckoutRepository, mockOrderRepository)

		mockCheckoutRepository.EXPECT().
			FindbyOrderID(ctx, gomock.Any()).
			Return(nil, nil)

		mockOrderRepository.EXPECT().
			Find(ctx, gomock.Any()).
			Return(nil, errors.New("unexpected error"))

		router := gin.Default()
		controller := &CheckPaymentCheckoutController{
			usecase: mockUseCase,
			ctx:     ctx,
		}

		router.GET("/checkout/check/:id", controller.GetCheckPaymentCheckout)

		req, _ := http.NewRequest(http.MethodGet, "/checkout/check/123", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.JSONEq(t, `{"error":"unexpected error"}`, resp.Body.String())
	})

}
