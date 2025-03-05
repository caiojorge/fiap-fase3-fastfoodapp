package usecase

import (
	"time"

	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/shared"
)

type OrderFindByParamInputDTO struct{}

type OrderFindByParamOutputDTO struct {
	ID             string                  `json:"id"`
	Items          []*usecase.OrderItemDTO `json:"items"`
	Total          float64                 `json:"total"`
	Status         string                  `json:"status"`
	CustomerCPF    string                  `json:"customercpf"`
	CreatedAt      time.Time               `json:"created_at"`
	DeliveryNumber string                  `json:"delivery_number"`
}
