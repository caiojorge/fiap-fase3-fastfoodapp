package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindOrdersByParams_Success(t *testing.T) {
	// Inicializar o controlador do GoMock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Criar o mock do repositório
	mockOrderRepo := mocksrepository.NewMockOrderRepository(ctrl)

	// Criar o caso de uso com o mock injetado
	findOrdersByParamsUseCase := NewOrderFindByParams(mockOrderRepo)

	// Definir o contexto e os parâmetros de busca
	ctx := context.Background()
	params := map[string]interface{}{
		"customer_cpf": "123.456.789-09",
		"status":       "pending",
	}

	status := entity.Status{
		Name: "confirmed",
	}

	// Criar ordens de teste que o mock retornará
	testOrders := []*entity.Order{
		{
			ID:          "order-123",
			CustomerCPF: "123.456.789-09",
			Total:       100.0,
			CreatedAt:   time.Now(),
			Status:      status,
			Items: []*entity.OrderItem{
				{ID: "item-1", ProductID: "prod1", Quantity: 2, Price: 50.00},
			},
		},
		{
			ID:          "order-456",
			CustomerCPF: "123.456.789-09",
			Total:       200.0,
			CreatedAt:   time.Now(),
			Status:      status,
			Items: []*entity.OrderItem{
				{ID: "item-2", ProductID: "prod2", Quantity: 1, Price: 200.00},
			},
		},
	}

	// Configurar o mock para retornar as ordens de teste
	mockOrderRepo.EXPECT().
		FindByParams(gomock.Any(), params).
		Return(testOrders, nil).Times(1)

	// Executar o caso de uso
	output, err := findOrdersByParamsUseCase.FindOrdersByParams(ctx, params)

	// Verificar o resultado
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output, 2)
	assert.Equal(t, "order-123", output[0].ID)
	assert.Equal(t, "order-456", output[1].ID)
}
