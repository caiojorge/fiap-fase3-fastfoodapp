package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
)

func TestCreateOrder(t *testing.T) {
	// Passo 1: Inicializar o controlador do GoMock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Passo 2: Criar o mock para o repositório de ordem
	mockOrder := mocksrepository.NewMockOrderRepository(ctrl)
	mockCustomer := mocksrepository.NewMockCustomerRepository(ctrl)
	mockProduct := mocksrepository.NewMockProductRepository(ctrl)

	// Passo 3: Criar o caso de uso com o mock injetado
	createOrderUseCase := NewOrderCreate(mockOrder, mockCustomer, mockProduct)

	// Passo 4: Definir o input para o teste
	input := &OrderCreateInputDTO{
		CustomerCPF: "123.456.789-09",
		Items: []*OrderItemCreateInputDTO{
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

	// Passo 7: Executar o caso de uso
	output, err := createOrderUseCase.CreateOrder(context.Background(), input)

	// Passo 8: Verificar o resultado
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)
	assert.Equal(t, "order-confirmed", output.Status)
	assert.Equal(t, 2, len(output.Items))
	assert.Equal(t, 200.0, output.Total)

}
