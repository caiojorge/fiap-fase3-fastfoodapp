package model

import (
	"fmt"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
)

// Checkout belongs to an order. It is the final step of the order process
type Checkout struct {
	ID                   string    `gorm:"type:char(36);primaryKey"`
	OrderID              string    `gorm:"not null;index:idx_order_id"`
	Order                Order     `gorm:"foreignKey:OrderID"`
	GatewayName          string    `gorm:"not null;type:varchar(255)"`
	GatewayToken         string    `gorm:"not null;type:varchar(255)"`
	GatewayTransactionID string    `gorm:"type:varchar(255)"`
	Total                float64   `gorm:"not null"`
	CreatedAt            time.Time `gorm:"not null"`
	QRCode               string    `gorm:"type:varchar(255)"`
}

func (c *Checkout) FromEntity(entity *Checkout) {
	c.ID = entity.ID
	c.OrderID = entity.OrderID
	c.GatewayName = entity.GatewayName
	c.GatewayToken = entity.GatewayToken
	c.GatewayTransactionID = entity.GatewayTransactionID
	c.Total = entity.Total
	c.CreatedAt = entity.CreatedAt
}

func (c *Checkout) ToEntity() *entity.Checkout {
	return &entity.Checkout{
		ID:        c.ID,
		OrderID:   c.OrderID,
		Gateway:   valueobject.NewGateway(c.GatewayName, c.GatewayToken),
		Total:     c.Total,
		CreatedAt: c.CreatedAt,
	}
}

func (c *Checkout) Validate() error {

	if c.ID == "" {
		return fmt.Errorf("id is required")
	}

	if c.OrderID == "" {
		return fmt.Errorf("orderID is required")
	}

	if c.GatewayName == "" {
		return fmt.Errorf("gatewayName is required")
	}

	if c.GatewayToken == "" {
		return fmt.Errorf("gatewayToken is required")
	}

	if c.Total == 0 {
		return fmt.Errorf("total is required")
	}

	return nil
}
