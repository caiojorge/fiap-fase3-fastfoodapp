package usecase

import (
	"context"
	"errors"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	customerrors "github.com/caiojorge/fiap-challenge-ddd/internal/shared/error"
)

type KitchenDeliveryInputDTO struct {
	OrderID string `json:"order_id"`
}

type KitchenDeliveryOutputDTO struct {
	OrderID       string   `json:"order_id"`
	KitchenID     string   `json:"kitchen_id"`
	CustomerID    string   `json:"customer_id"`
	Queue         string   `json:"queue"`
	EstimatedTime string   `json:"estimated_time"`
	Status        string   `json:"status"`
	Items         []string `json:"items"`
}

// IKitchenDeliveryUseCase é uma interface para o caso de uso de notificação da cozinha
type IKitchenDeliveryUseCase interface {
	Delivery(ctx context.Context, input KitchenDeliveryInputDTO) (*KitchenDeliveryOutputDTO, error)
}

// KitchenDeliveryUseCase é responsável por notificar a cozinha
type KitchenDeliveryUseCase struct {
	repoKitchen ports.KitchenRepository
	repoOrder   ports.OrderRepository
	repoProduct ports.ProductRepository
}

// NewKitchenDeliveryUseCase cria um novo caso de uso de notificação da cozinha
func NewKitchenDeliveryUseCase(repoKitchen ports.KitchenRepository, repoOrder ports.OrderRepository, repoProduct ports.ProductRepository) *KitchenDeliveryUseCase {
	return &KitchenDeliveryUseCase{
		repoKitchen: repoKitchen,
		repoOrder:   repoOrder,
		repoProduct: repoProduct,
	}
}

/*
Delivery é responsável por buscar a ordem e o ticket da cozinha; mover o status para o próximo da fase de preparo e fazer o delivery se estiver finalizado.
*/
func (uc *KitchenDeliveryUseCase) Delivery(ctx context.Context, input KitchenDeliveryInputDTO) (*KitchenDeliveryOutputDTO, error) {

	// 1. busca a ordem especifica
	orders, err := uc.repoOrder.Find(ctx, input.OrderID)
	if err != nil {
		return nil, err
	}

	if orders == nil {
		return nil, customerrors.ErrOrderNotFound
	}

	if orders.Status.Name == sharedconsts.OrderFinalizedByKitchen {
		return nil, errors.New("this order is already finalized")
	}

	// 2. verifica se o status esta no range correto (a partir de recebido pela cozinha)
	// vai até OrderReadyByKitchen
	isOk, err := sharedconsts.IsStatusBetween(orders.Status.Name, 6, 6)
	if err != nil {
		return nil, err
	}

	if !isOk {
		return nil, errors.New("order not in the right phase for delivery")
	}

	var outputItems []string
	for _, item := range orders.Items {
		product, err := uc.repoProduct.Find(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}
		outputItems = append(outputItems, product.Name)
	}

	// 4. atualiza o status da ordem
	err = uc.repoOrder.UpdateStatus(ctx, input.OrderID, sharedconsts.OrderFinalizedByKitchen)
	if err != nil {
		return nil, err
	}

	// 5. busca o ticket da cozinha relacionado à ordem
	kitchens, err := uc.repoKitchen.FindByParams(ctx, map[string]interface{}{"order_id": input.OrderID})
	if err != nil {
		return nil, err
	}

	// 6. cria a saida
	output := KitchenDeliveryOutputDTO{
		OrderID:       orders.ID,
		KitchenID:     kitchens[0].ID,
		Queue:         kitchens[0].Queue,
		EstimatedTime: kitchens[0].EstimatedTime,
		Status:        sharedconsts.OrderFinalizedByKitchen,
		Items:         outputItems,
		CustomerID:    orders.CustomerCPF,
	}
	// 7. retorna as ordens notificadas
	return &output, nil
}
