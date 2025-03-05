package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindOrderByID(t *testing.T) {

	t.Run("FindOrderByID_Success", func(t *testing.T) {
		// Inicializar o controlador do GoMock
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Criar o mock para o repositório de ordem
		mockOrderRepo := mocksrepository.NewMockOrderRepository(ctrl)

		// Criar o caso de uso com o mock injetado
		findOrderByIDUseCase := NewOrderFindByID(mockOrderRepo)

		// Definir o contexto e o ID da ordem
		ctx := context.Background()
		orderID := "order-123"

		status := entity.Status{
			Name: "confirmed",
		}

		// Criar uma ordem de teste que o mock retornará
		testOrder := &entity.Order{
			ID:          orderID,
			CustomerCPF: "123.456.789-00",
			Total:       100.0,
			CreatedAt:   time.Now(),
			Status:      status,
			Items: []*entity.OrderItem{
				{ID: "item-1", ProductID: "prod1", Quantity: 2, Price: 50.00},
			},
		}

		// Configurar o mock para retornar a ordem de teste
		mockOrderRepo.EXPECT().
			Find(gomock.Any(), orderID).
			Return(testOrder, nil).Times(1)

		// Executar o caso de uso
		output, err := findOrderByIDUseCase.FindOrderByID(ctx, orderID)

		// Verificar o resultado
		assert.Nil(t, err)                  // Verificar que não houve erro
		assert.NotNil(t, output)            // Verificar que o resultado não é nulo
		assert.Equal(t, orderID, output.ID) // Verificar que o ID da ordem está correto
		assert.Equal(t, "123.456.789-00", output.CustomerCPF)
		assert.Len(t, output.Items, 1) // Verificar que há 1 item na ordem
		assert.Equal(t, "prod1", output.Items[0].ProductID)
		assert.Equal(t, 100.0, output.Total)
	})

	t.Run("FindOrderByID_Error", func(t *testing.T) {
		// Inicializar o controlador do GoMock
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Criar o mock para o repositório de ordem
		mockOrderRepo := mocksrepository.NewMockOrderRepository(ctrl)

		// Criar o caso de uso com o mock injetado
		findOrderByIDUseCase := NewOrderFindByID(mockOrderRepo)

		// Definir o contexto e o ID da ordem
		ctx := context.Background()
		orderID := "order-123"

		// Configurar o mock para retornar um erro
		mockOrderRepo.EXPECT().
			Find(gomock.Any(), orderID).
			Return(nil, errors.New("order not found")).Times(1)

		// Executar o caso de uso
		output, err := findOrderByIDUseCase.FindOrderByID(ctx, orderID)

		// Verificar o resultado
		assert.Nil(t, output)                           // Verificar que o resultado é nulo
		assert.NotNil(t, err)                           // Verificar que houve um erro
		assert.Equal(t, "order not found", err.Error()) // Verificar a mensagem de erro
	})

}
