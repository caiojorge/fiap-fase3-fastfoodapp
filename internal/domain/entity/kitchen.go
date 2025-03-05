package entity

import (
	"time"

	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"
	sharedtime "github.com/caiojorge/fiap-challenge-ddd/internal/shared/time"
)

type Kitchen struct { // aqui seria mais um kitchen ticket... um ticket por item
	ID             string
	OrderID        string    // id do pedido
	Queue          string    // ordem na fila de preparo
	EstimatedTime  string    // tempo estimado para preparo
	CreatedAt      time.Time // data / hora de qdo o pedido foi recebido pela cozinha
	Items          []string  // lista de itens do pedido
	Status         string    // status do pedido na cozinha
	DeliveryNumber string    // numero do pedido de entrega
}

func NewKitchen(orderID string) *Kitchen {

	return &Kitchen{
		ID:        sharedgenerator.NewIDGenerator(),
		OrderID:   orderID,
		CreatedAt: sharedtime.GetBRTimeNow(),
	}
}

func (k *Kitchen) SetQueue(queue string) {
	k.Queue = queue
}
func (k *Kitchen) SetEstimatedTime(estimatedTime string) {
	k.EstimatedTime = estimatedTime
}

func (k *Kitchen) AddItems(items []string) {
	k.Items = items
}
