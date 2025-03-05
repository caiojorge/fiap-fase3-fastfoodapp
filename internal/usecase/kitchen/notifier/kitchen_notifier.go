package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"
	sharedtime "github.com/caiojorge/fiap-challenge-ddd/internal/shared/time"
	"golang.org/x/exp/rand"
)

type KitchenNotifierInputDTO struct{}
type KitchenNotifierOutputDTO struct { // pode ser usado no monitor para atender o requisito de mostrar as ordens na fila
	ID            string   `json:"id"`
	OrderID       string   `json:"order_id"`
	CustomerID    string   `json:"customer_id"`
	Queue         string   `json:"queue"` // ordem na fila de preparo
	EstimatedTime string   `json:"estimated_time"`
	Status        string   `json:"status"`
	Items         []string `json:"items"`
}

// IKitchenNotifierUseCase é uma interface para o caso de uso de notificação da cozinha
type IKitchenNotifierUseCase interface {
	Notify(ctx context.Context) ([]*KitchenNotifierOutputDTO, error)
}

// KitchenNotifierUseCase é responsável por notificar a cozinha
type KitchenNotifierUseCase struct {
	repoKitchen ports.KitchenRepository
	repoOrder   ports.OrderRepository
	repoProduct ports.ProductRepository
}

// NewKitchenNotifierUseCase cria um novo caso de uso de notificação da cozinha
func NewKitchenNotifierUseCase(repoKitchen ports.KitchenRepository, repoOrder ports.OrderRepository, repoProduct ports.ProductRepository) *KitchenNotifierUseCase {
	return &KitchenNotifierUseCase{
		repoKitchen: repoKitchen,
		repoOrder:   repoOrder,
		repoProduct: repoProduct,
	}
}

/*
Notifier é responsável por buscar as ordens pagas e não notificadas,
notificar a cozinha e atualizar o status da ordem para recebida pela cozinha.
*/
func (uc *KitchenNotifierUseCase) Notify(ctx context.Context) ([]*KitchenNotifierOutputDTO, error) {

	// 1. busca ordens pagas e não notificadas
	orders, err := uc.repoOrder.FindByParams(ctx, map[string]interface{}{"status": sharedconsts.OrderStatusPaymentApproved})
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, nil
	}

	// 2. notifica a cozinha, criando o objeto kitchen
	//var kitchenOrders []*entity.Kitchen
	var outputs []*KitchenNotifierOutputDTO
	rand.Seed(uint64(time.Now().UnixNano()))

	for i, order := range orders {
		randomNumber := rand.Intn(51) + 10
		kt := &entity.Kitchen{
			ID:            sharedgenerator.NewIDGenerator(),
			OrderID:       order.ID,
			Queue:         fmt.Sprintf("%03d", i+1),         // ordem na fila de preparo - é apenas uma ideia, e não irei implementar a solução completa
			EstimatedTime: fmt.Sprintf("%dm", randomNumber), // tbm, é apenas uma ideia - não irei implementar a solução completa
			CreatedAt:     sharedtime.GetBRTimeNow(),
		}

		var outputItems []string
		for _, item := range order.Items {
			product, err := uc.repoProduct.Find(ctx, item.ProductID)
			if err != nil {
				return nil, err
			}
			outputItems = append(outputItems, product.Name)
		}

		kt.AddItems(outputItems)

		output := &KitchenNotifierOutputDTO{
			ID:            kt.ID,
			OrderID:       order.ID,
			CustomerID:    order.CustomerCPF,
			Status:        sharedconsts.OrderReceivedByKitchen,
			Queue:         kt.Queue,
			EstimatedTime: kt.EstimatedTime,
			Items:         outputItems,
		}

		err := uc.repoKitchen.Create(ctx, kt)
		if err != nil {
			return nil, err
		}
		//kitchenOrders = append(kitchenOrders, kt)
		outputs = append(outputs, output)
	}

	// 3. atualiza a ordem para recebida pela cozinha
	for _, order := range orders {
		order.Status = *entity.NewStatus(sharedconsts.OrderReceivedByKitchen)
		err := uc.repoOrder.UpdateStatus(ctx, order.ID, order.Status.Name)
		if err != nil {
			return nil, err
		}
	}

	// 4. retorna as ordens notificadas
	return outputs, nil
}
