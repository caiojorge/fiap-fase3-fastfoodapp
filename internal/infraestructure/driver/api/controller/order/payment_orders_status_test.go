package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocks "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findbyparam"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConfirmedOrder(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gin.SetMode(gin.TestMode) // Modo de teste para o Gin

	mockOrderRepository := mocks.NewMockOrderRepository(ctrl)
	assert.NotNil(t, mockOrderRepository)

	findOrdersByParamsUseCase := usecase.NewOrderFindByParams(mockOrderRepository)

	ctx := context.Background()

	t.Run("should return 200 - Confirmed", func(t *testing.T) {
		status := &entity.Status{
			Name: sharedconsts.OrderStatusConfirmed,
		}

		testOrders := []*entity.Order{
			{
				ID:          "order-123",
				CustomerCPF: "123.456.789-09",
				Total:       100.0,
				CreatedAt:   time.Now(),
				Status:      *status,
				Items: []*entity.OrderItem{
					{ID: "item-1", ProductID: "prod1", Quantity: 2, Price: 50.00},
				},
			},
			{
				ID:          "order-456",
				CustomerCPF: "123.456.789-09",
				Total:       200.0,
				CreatedAt:   time.Now(),
				Status:      *status,
				Items: []*entity.OrderItem{
					{ID: "item-2", ProductID: "prod2", Quantity: 1, Price: 200.00},
				},
			},
		}

		// Configurar o mock para retornar as ordens de teste
		mockOrderRepository.EXPECT().
			FindByParams(gomock.Any(), gomock.Any()).
			Return(testOrders, nil).Times(1)

		controller := NewFindByParamsConfirmedController(ctx, findOrdersByParamsUseCase)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/orders/confirmed", controller.GetOrdersConfirmed)

		req, err := http.NewRequest(http.MethodGet, "/orders/confirmed", nil)
		assert.Nil(t, err)
		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		responseBody := w.Body.String()
		t.Logf("Response JSON: %s", responseBody)

		var orders []*usecase.OrderFindByParamOutputDTO

		err = json.Unmarshal([]byte(responseBody), &orders)
		assert.Nil(t, err)
		assert.Len(t, orders, len(testOrders))
		assert.Equal(t, testOrders[0].ID, orders[0].ID)
		assert.Equal(t, testOrders[1].ID, orders[1].ID)

	})

	t.Run("should return 200 - Not Confirmed", func(t *testing.T) {
		status := &entity.Status{
			Name: sharedconsts.OrderStatusNotConfirmed,
		}
		testOrders := []*entity.Order{
			{
				ID:          "order-123",
				CustomerCPF: "123.456.789-09",
				Total:       100.0,
				CreatedAt:   time.Now(),
				Status:      *status,
				Items: []*entity.OrderItem{
					{ID: "item-1", ProductID: "prod1", Quantity: 2, Price: 50.00},
				},
			},
			{
				ID:          "order-456",
				CustomerCPF: "123.456.789-09",
				Total:       200.0,
				CreatedAt:   time.Now(),
				Status:      *status,
				Items: []*entity.OrderItem{
					{ID: "item-2", ProductID: "prod2", Quantity: 1, Price: 200.00},
				},
			},
		}

		// Configurar o mock para retornar as ordens de teste
		mockOrderRepository.EXPECT().
			FindByParams(gomock.Any(), gomock.Any()).
			Return(testOrders, nil).Times(1)

		controller := NewFindByParamsNotConfirmedController(ctx, findOrdersByParamsUseCase)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/orders/pending", controller.GetOrdersNotConfirmed)

		req, err := http.NewRequest(http.MethodGet, "/orders/pending", nil)
		assert.Nil(t, err)
		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		responseBody := w.Body.String()
		t.Logf("Response JSON: %s", responseBody)

		var orders []*usecase.OrderFindByParamOutputDTO

		err = json.Unmarshal([]byte(responseBody), &orders)
		assert.Nil(t, err)
		assert.Len(t, orders, len(testOrders))
		assert.Equal(t, testOrders[0].ID, orders[0].ID)
		assert.Equal(t, testOrders[1].ID, orders[1].ID)
	})

	t.Run("should return 200 - Paid", func(t *testing.T) {
		status := &entity.Status{
			Name: sharedconsts.OrderStatusPaymentApproved,
		}
		testOrders := []*entity.Order{
			{
				ID:          "order-123",
				CustomerCPF: "123.456.789-09",
				Total:       100.0,
				CreatedAt:   time.Now(),
				Status:      *status,
				Items: []*entity.OrderItem{
					{ID: "item-1", ProductID: "prod1", Quantity: 2, Price: 50.00},
				},
			},
			{
				ID:          "order-456",
				CustomerCPF: "123.456.789-09",
				Total:       200.0,
				CreatedAt:   time.Now(),
				Status:      *status,
				Items: []*entity.OrderItem{
					{ID: "item-2", ProductID: "prod2", Quantity: 1, Price: 200.00},
				},
			},
		}

		// Configurar o mock para retornar as ordens de teste
		mockOrderRepository.EXPECT().
			FindByParams(gomock.Any(), gomock.Any()).
			Return(testOrders, nil).Times(1)

		controller := NewFindByParamsPaymentApprovedController(ctx, findOrdersByParamsUseCase)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/orders/paid", controller.GetOrdersWithPaymentApproved)

		req, err := http.NewRequest(http.MethodGet, "/orders/paid", nil)
		assert.Nil(t, err)
		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		responseBody := w.Body.String()
		t.Logf("Response JSON: %s", responseBody)

		var orders []*usecase.OrderFindByParamOutputDTO

		err = json.Unmarshal([]byte(responseBody), &orders)
		assert.Nil(t, err)
		assert.Len(t, orders, len(testOrders))
		assert.Equal(t, testOrders[0].ID, orders[0].ID)
		assert.Equal(t, testOrders[1].ID, orders[1].ID)
		assert.Equal(t, sharedconsts.OrderStatusPaymentApproved, orders[0].Status)
		assert.Equal(t, sharedconsts.OrderStatusPaymentApproved, orders[1].Status)

	})

}
