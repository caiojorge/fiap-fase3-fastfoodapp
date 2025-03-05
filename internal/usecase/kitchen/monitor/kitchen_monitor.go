package usecase

import (
	"context"
	"time"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type MonitorKitchenInputDTO struct{}

type MonitorKitchenOutputDTO struct {
	ID             string    `json:"id"`
	OrderID        string    `json:"order_id"`
	Queue          string    `json:"queue"`
	EstimatedTime  string    `json:"estimated_time"`
	CreatedAt      time.Time `json:"created_at"`
	Status         string    `json:"status"`
	DeliveryNumber string    `json:"delivery_number"`
}

type IMonitorKitchenUseCase interface {
	Monitor(ctx context.Context) ([]*MonitorKitchenOutputDTO, error)
}

type MonitorKitchenUseCase struct {
	repoKitchen portsrepository.KitchenRepository
	repoOrder   portsrepository.OrderRepository
}

func NewMonitorKitchenUseCase(repoKitchen portsrepository.KitchenRepository, repoOrder portsrepository.OrderRepository) *MonitorKitchenUseCase {
	return &MonitorKitchenUseCase{
		repoKitchen: repoKitchen,
		repoOrder:   repoOrder,
	}
}

func (m *MonitorKitchenUseCase) Monitor(ctx context.Context) ([]*MonitorKitchenOutputDTO, error) {
	kitchens, err := m.repoKitchen.Monitor(ctx)
	if err != nil {
		return nil, err
	}

	// OrderReceivedByKitchen,       // 4
	// OrderInPreparationByKitchen,  // 5
	// OrderReadyByKitchen,          // 6

	var outputs []*MonitorKitchenOutputDTO

	for _, kitchen := range kitchens {

		output := &MonitorKitchenOutputDTO{
			ID:             kitchen.ID,
			OrderID:        kitchen.OrderID,
			Queue:          kitchen.Queue,
			EstimatedTime:  kitchen.EstimatedTime,
			CreatedAt:      kitchen.CreatedAt,
			Status:         kitchen.Status,
			DeliveryNumber: kitchen.DeliveryNumber,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}
