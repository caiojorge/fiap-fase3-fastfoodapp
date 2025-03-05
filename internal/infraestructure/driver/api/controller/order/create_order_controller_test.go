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
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/create"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {

	/*
		type CreateOrderController struct {
			usecase usecaseorder.CreateOrderUseCase
			ctx     context.Context
		}
		type OrderCreateUseCase struct {
			orderRepository    domainRepository.OrderRepository
			customerRepository domainRepository.CustomerRepository
			productRepository  domainRepository.ProductRepository
		}
	*/

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gin.SetMode(gin.TestMode) // Modo de teste para o Gin

	mockOrderRepository := mocks.NewMockOrderRepository(ctrl)
	assert.NotNil(t, mockOrderRepository)

	ctx := context.Background()

	t.Run("should return 200", func(t *testing.T) {
		// Passo 2: Criar o mock para o repositório de ordem
		mockOrder := mocks.NewMockOrderRepository(ctrl)
		mockCustomer := mocks.NewMockCustomerRepository(ctrl)
		mockProduct := mocks.NewMockProductRepository(ctrl)

		// Passo 3: Criar o caso de uso com o mock injetado
		createOrderUseCase := usecase.NewOrderCreate(mockOrder, mockCustomer, mockProduct)

		// Passo 4: Definir o input para o teste
		input := &usecase.OrderCreateInputDTO{
			CustomerCPF: "123.456.789-09",
			Items: []*usecase.OrderItemCreateInputDTO{
				{ProductID: "prod1", Quantity: 2},
				{ProductID: "prod2", Quantity: 1},
			},
		}

		// Passo 5: Criar a entidade de ordem que será usada no mock
		// Configurar um cliente de teste e um produto de teste
		cpf := valueobject.CPF{
			Value: "123.456.789-09",
		}

		err := cpf.Validate()
		assert.Nil(t, err)

		testCustomer := &entity.Customer{
			CPF:   cpf,
			Name:  "John Doe",
			Email: "johndoe@example.com",
		}
		testProduct := &entity.Product{
			ID:          "prod1",
			Name:        "Produto 1",
			Description: "Produto 1",
			Category:    "Produto 1",
			Price:       50.00,
		}

		testProduct2 := &entity.Product{
			ID:          "prod2",
			Name:        "Produto 2",
			Description: "Produto 2",
			Category:    "Produto 2",
			Price:       100.00,
		}

		mockCustomer.EXPECT().
			Find(gomock.Any(), "12345678909").
			Return(testCustomer, nil).Times(1)

		// Expectativa: Chamar FindProduct para buscar os produtos da ordem
		mockProduct.EXPECT().
			Find(gomock.Any(), "prod1").
			Return(testProduct, nil).Times(1)

		// Configurar para o segundo produto (prod2), retornando outro produto
		mockProduct.EXPECT().
			Find(gomock.Any(), "prod2").
			Return(testProduct2, nil).Times(1)

		// Expectativa: Criar a ordem com sucesso
		mockOrder.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).Times(1)

		controller := NewCreateOrderController(ctx, createOrderUseCase)

		gin.SetMode(gin.TestMode)
		r := gin.Default()

		r.POST("/order", controller.PostCreateOrder)

		jsonData, err := json.MarshalIndent(input, "", "  ")
		assert.Nil(t, err)

		requestBody := bytes.NewBuffer(jsonData)

		req, err := http.NewRequest("POST", "/order", requestBody)
		if err != nil {
			t.Fatalf("Couldn't create request: %v\n", err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

}
