package usecase

import (
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
)

type KitchenInputDTO struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (dto *KitchenInputDTO) ToEntity() *entity.Kitchen {
	return &entity.Kitchen{
		ID:        dto.ID,
		OrderID:   dto.OrderID,
		CreatedAt: dto.CreatedAt,
	}
}

type KitchenFindAllAOutputDTO struct {
	ID            string    `json:"id"`
	OrderID       string    `json:"order_id"`
	Queue         string    `json:"queue"`
	EstimatedTime string    `json:"estimated_time"`
	CreatedAt     time.Time `json:"created_at"`
}

func (dto *KitchenFindAllAOutputDTO) FromEntity(customer entity.Kitchen) {
	dto.ID = customer.ID
	dto.OrderID = customer.OrderID
	dto.CreatedAt = customer.CreatedAt

}
