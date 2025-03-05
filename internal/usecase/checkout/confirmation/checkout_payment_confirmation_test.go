package usecase

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocks "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestConfirmPayment_Success(t *testing.T) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	ctrl := gomock.NewController(t) // Cria um controlador do gomock
	defer ctrl.Finish()             // Libera os mocks após o teste

	// Mocks
	orderRepoMock := mocks.NewMockOrderRepository(ctrl)
	checkoutRepoMock := mocks.NewMockCheckoutRepository(ctrl)
	transactionManagerMock := mocks.NewMockTransactionManager(ctrl)

	// Use Case
	checkoutConfirmationUseCase := NewCheckoutConfirmation(orderRepoMock, checkoutRepoMock, transactionManagerMock, logger)

	// Dados de entrada e entidades simuladas
	ctx := context.Background()
	input := &CheckoutConfirmationInputDTO{
		OrderID: "order123",
	}

	status := entity.Status{
		Name: "confirmed",
	}

	order := &entity.Order{
		ID:     "order123",
		Status: status,
	}
	checkout := &entity.Checkout{
		ID:      "checkout123",
		OrderID: "order123",
	}

	// Configuração dos mocks
	orderRepoMock.EXPECT().Find(ctx, input.OrderID).Return(order, nil)
	checkoutRepoMock.EXPECT().FindbyOrderID(ctx, input.OrderID).Return(checkout, nil)
	orderRepoMock.EXPECT().UpdateStatus(ctx, "order123", gomock.Any()).Return(nil)
	// checkoutRepoMock.EXPECT().UpdateStatus(ctx, gomock.Any(), gomock.Any()).Return(nil)
	transactionManagerMock.EXPECT().
		RunInTransaction(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
			return fn(ctx) // Executa a função transacional passada
		})

	// Execução do teste
	output, err := checkoutConfirmationUseCase.ConfirmPayment(ctx, input)

	// Asserções
	assert.NoError(t, err)
	assert.NotNil(t, output)
}
