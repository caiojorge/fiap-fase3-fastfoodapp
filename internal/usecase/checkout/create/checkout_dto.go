package usecase

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
)

// Checkout Pedido que deverá receber os produtos
// solicitados
// os itens da ordem serão recuperados a partir da ordem que esta sendo informada no checkout
type CheckoutInputDTO struct {
	OrderID         string  `json:"order_id"`         // o ID do pedido que será pago
	GatewayName     string  `json:"gateway_name"`     // nome do gateway (nesse caso, mercado livre fake)
	GatewayToken    string  `json:"gateway_token"`    // id ou token para uso no gateway
	NotificationURL string  `json:"notification_url"` // webhook para receber a confirmação do pagamento
	SponsorID       int     `json:"sponsor_id"`       // ID do patrocinador
	DiscontCoupon   float64 `json:"discont_coupon"`   // valor de desconto... só uma ideia nesse momento.
}

func (dto *CheckoutInputDTO) ToEntity() *entity.Checkout {
	return &entity.Checkout{
		//ID:                   dto.ID,
		OrderID: dto.OrderID,
		Gateway: valueobject.NewGateway(dto.GatewayName, dto.GatewayToken),
	}
}

type CheckoutOutputDTO struct {
	ID                   string `json:"id"`                     // ID do checkout
	GatewayTransactionID string `json:"gateway_transaction_id"` // ID de transação gerado pelo gateway
	OrderID              string `json:"order_id"`               // Identificação da ordem (pedido)
}

// retornar à identificação do pedido.

func (dto *CheckoutOutputDTO) FromEntity(entity entity.Checkout) {
	dto.ID = entity.ID
	dto.GatewayTransactionID = entity.Gateway.GatewayTransactionID
	dto.OrderID = entity.OrderID
}
