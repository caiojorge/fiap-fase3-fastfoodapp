package sharedconsts

import (
	"errors"
	"fmt"
)

// Order status
const (
	OrderStatusNotConfirmed      = "order-not-confirmed"  // status inicial da ordem
	OrderStatusConfirmed         = "order-confirmed"      // order confirmed by the customer
	OrderStatusCheckoutConfirmed = "checkout-confirmed"   // checkout confirmado e aguardando pagamento
	OrderStatusPaymentApproved   = "payment-approved"     // payment approved by the payment gateway
	OrderStatusNotApproved       = "payment-not-approved" // em caso de recusa do pagamento pelo gateway

	OrderReceivedByKitchen      = "order-received-by-kitchen"       // pedido recebido pela cozinha (recebido)
	OrderInPreparationByKitchen = "order-in-preparation-by-kitchen" // pedido em preparo na cozinha (em preparo)
	OrderReadyByKitchen         = "order-ready-by-kitchen"          // pedido pronto na cozinha (pronto)
	OrderFinalizedByKitchen     = "order-finalized-by-kitchen"      // pedido finalizado na cozinha (finalizado)
)

// Order Item status
const (
	OrderItemStatusConfirmed = "item-confirmed"
	OrderItemStatusCanceled  = "item-canceled"
)

// Ordem cronológica dos status
var StatusFlow = []string{
	OrderStatusNotConfirmed,      // 0
	OrderStatusConfirmed,         // 1
	OrderStatusCheckoutConfirmed, // 2
	OrderStatusPaymentApproved,   // 3
	OrderReceivedByKitchen,       // 4
	OrderInPreparationByKitchen,  // 5
	OrderReadyByKitchen,          // 6
	OrderFinalizedByKitchen,      // 7
}

// Função para obter o próximo status
func GetNextStatus(currentStatus string) (string, error) {
	for i, status := range StatusFlow {
		if status == currentStatus {
			if i+1 < len(StatusFlow) {
				return StatusFlow[i+1], nil
			}
			return "", errors.New("status final já alcançado")
		}
	}
	return "", errors.New("status inválido")
}

// Função para obter a posição de um status
func GetStatusIndex(status string) (int, error) {
	for i, s := range StatusFlow {
		if s == status {
			return i, nil
		}
	}
	return -1, fmt.Errorf("status '%s' não encontrado", status)
}

// Função para validar se o status está entre duas posições
func IsStatusBetween(status string, start, end int) (bool, error) {
	index, err := GetStatusIndex(status)
	if err != nil {
		return false, err
	}

	return index >= start && index <= end, nil
}
