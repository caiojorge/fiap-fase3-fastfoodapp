package usecase

import (
	"context"
	"errors"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	customerrors "github.com/caiojorge/fiap-challenge-ddd/internal/shared/error"
)

type KitchenCookingInputDTO struct {
	OrderID string `json:"order_id"`
}

type KitchenCookingOutputDTO struct {
	OrderID       string   `json:"order_id"`
	KitchenID     string   `json:"kitchen_id"`
	CustomerID    string   `json:"customer_id"`
	Queue         string   `json:"queue"`
	EstimatedTime string   `json:"estimated_time"`
	Status        string   `json:"status"`
	Items         []string `json:"items"`
}

// IKitchenCookingUseCase é uma interface para o caso de uso de notificação da cozinha
type IKitchenCookingUseCase interface {
	Cook(ctx context.Context, input KitchenCookingInputDTO) (*KitchenCookingOutputDTO, error)
}

// KitchenCookingUseCase é responsável por notificar a cozinha
type KitchenCookingUseCase struct {
	repoKitchen ports.KitchenRepository
	repoOrder   ports.OrderRepository
	repoProduct ports.ProductRepository
}

// NewKitchenCookingUseCase cria um novo caso de uso de notificação da cozinha
func NewKitchenCookingUseCase(repoKitchen ports.KitchenRepository, repoOrder ports.OrderRepository, repoProduct ports.ProductRepository) *KitchenCookingUseCase {
	return &KitchenCookingUseCase{
		repoKitchen: repoKitchen,
		repoOrder:   repoOrder,
		repoProduct: repoProduct,
	}
}

/*
Cook é responsável por buscar a ordem e o ticket da cozinha; mover o status para o próximo da fase de preparo e fazer o delivery se estiver finalizado.
*/
func (uc *KitchenCookingUseCase) Cook(ctx context.Context, input KitchenCookingInputDTO) (*KitchenCookingOutputDTO, error) {

	// 1. busca a ordem especifica
	orders, err := uc.repoOrder.Find(ctx, input.OrderID)
	if err != nil {
		return nil, err
	}

	if orders == nil {
		return nil, customerrors.ErrOrderNotFound
	}

	if orders.Status.Name == sharedconsts.OrderReadyByKitchen {
		return nil, errors.New("this order is already ready")

	}

	if orders.Status.Name == sharedconsts.OrderFinalizedByKitchen {
		return nil, errors.New("this order is already finalized")
	}

	// 2. verifica se o status esta no range correto (a partir de recebido pela cozinha)
	// vai até OrderReadyByKitchen
	isOk, err := sharedconsts.IsStatusBetween(orders.Status.Name, 4, 5)
	if err != nil {
		return nil, err
	}

	if !isOk {
		return nil, errors.New("order not in the right phase")
	}

	var outputItems []string
	for _, item := range orders.Items {
		product, err := uc.repoProduct.Find(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}
		outputItems = append(outputItems, product.Name)
	}

	// 3. pega a próxima fase da ordem no fluxo da cozinha
	nextPhase, err := sharedconsts.GetNextStatus(orders.Status.Name)
	if err != nil {
		return nil, errors.New("final status already reached " + err.Error())
	}

	if nextPhase == "" {
		return nil, errors.New("final status already reached")
	}

	// 4. atualiza o status da ordem
	err = uc.repoOrder.UpdateStatus(ctx, input.OrderID, nextPhase)
	if err != nil {
		return nil, err
	}

	// 5. busca o ticket da cozinha relacionado à ordem
	kitchens, err := uc.repoKitchen.FindByParams(ctx, map[string]interface{}{"order_id": input.OrderID})
	if err != nil {
		return nil, err
	}

	// 6. cria a saida
	output := KitchenCookingOutputDTO{
		OrderID:       orders.ID,
		KitchenID:     kitchens[0].ID,
		Queue:         kitchens[0].Queue,
		EstimatedTime: kitchens[0].EstimatedTime,
		Status:        nextPhase,
		Items:         outputItems,
		CustomerID:    orders.CustomerCPF,
	}
	// 7. retorna as ordens notificadas
	return &output, nil
}
