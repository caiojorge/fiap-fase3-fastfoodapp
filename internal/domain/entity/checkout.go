package entity

import (
	"errors"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"
)

type Checkout struct {
	ID        string
	OrderID   string
	Gateway   valueobject.Gateway
	Total     float64
	CreatedAt time.Time
	QRCode    string
}

func NewCheckout(orderID string, gatewayName string, gatewayToken string, total float64) (*Checkout, error) {

	return &Checkout{
		ID:      sharedgenerator.NewIDGenerator(),
		OrderID: orderID,
		Gateway: valueobject.NewGateway(gatewayName, gatewayToken),
		Total:   total,
	}, nil
}

func (c *Checkout) ConfirmTransaction(transactionID string, total float64, qrcode string) error {

	c.ID = sharedgenerator.NewIDGenerator()
	c.Gateway.GatewayTransactionID = transactionID
	c.Total = total
	c.QRCode = qrcode

	err := c.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (c *Checkout) Validate() error {
	if c.OrderID == "" {
		return errors.New("orderID is required")
	}

	return nil
}

func (c *Checkout) Reprocessing(gatewayID, QRCode string) error {

	c.QRCode = QRCode
	c.Gateway.GatewayTransactionID = gatewayID

	return nil
}
