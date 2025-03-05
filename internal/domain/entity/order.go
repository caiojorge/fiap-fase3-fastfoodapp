package entity

import (
	"errors"
	"time"

	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/deliver"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/validator"
)

/*
Simplicidade: O pedido (Order) representa o agregado principal no domínio. Saber se o pedido foi pago ou não é fundamental para determinar se ele pode avançar para outras etapas (ex: envio, entrega).
Estados do Pedido: O status de pagamento faz parte do ciclo de vida do pedido. Estados como "pendente" e "pago" são naturais em um pedido.
Agregação: A entidade Order é o ponto central para acompanhar o progresso do pedido, enquanto Checkout é apenas um processo auxiliar para concretizar o pagamento.
*/

// Order representa um pedido.
type Order struct {
	ID             string
	Items          []*OrderItem
	Total          float64
	Status         Status
	CustomerCPF    string
	CreatedAt      time.Time
	DeliveryNumber string
}

// NewOrder cria um novo pedido. TODO Não esta sendo usada.
func NewOrder(cpf string, items []*OrderItem) (*Order, error) {

	order := Order{
		ID:          sharedgenerator.NewIDGenerator(),
		CustomerCPF: cpf,
		Items:       items,
		Status:      *NewStatus(sharedconsts.OrderStatusConfirmed),
	}

	if len(order.Items) > 0 {
		order.CalculateTotal()
	}

	err := order.Validate()
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *Order) GetOrderItemByProductID(productID string) *OrderItem {

	for _, item := range o.Items {
		if item.ProductID == productID {
			return item
		}
	}

	return nil
}

// Confirm confirma o pedido. Tem muita lógica de negócio aqui.
// Toda preparação necessária, validação de cpf, cálculo do total e validação dos itens.
// As regras aplicadas impactam apenas os dados da ordem / item.
func (o *Order) Confirm() error {

	o.ID = sharedgenerator.NewIDGenerator()

	// o status da ordem é confirmado
	o.Status = *NewStatus(sharedconsts.OrderStatusConfirmed)

	// o número de entrega é gerado - o cliente usa esse número para retirar o pedido
	o.DeliveryNumber = deliver.GenerateDeliveryNumber()

	for _, item := range o.Items {
		item.Confirm()
	}

	// Calcula o total do pedido se o item for confirmado
	o.CalculateTotal()

	// Valida o pedido
	err := o.Validate()
	if err != nil {
		return errors.New("failed to validate order")
	}

	return nil
}

func (o *Order) ConfirmItemsPrice(products []*Product) {

	for i, item := range o.Items {
		for _, product := range products {
			if item.ProductID == product.ID {
				o.Items[i].ConfirmPrice(product.Price)
			}
		}
	}
}

func (o *Order) ConfirmCheckout() {
	o.Status = *NewStatus(sharedconsts.OrderStatusCheckoutConfirmed)
}

func (o *Order) ConfirmPayment() {
	o.Status = *NewStatus(sharedconsts.OrderStatusPaymentApproved)
}

func (o *Order) IsCustomerInformed() bool {
	return o.CustomerCPF != ""
}

func (o *Order) RemoveMaksFromCPF() {
	if o.CustomerCPF != "" {
		o.CustomerCPF = formatter.RemoveMaskFromCPF(o.CustomerCPF)
	}
}

func (o *Order) GetID() string {
	return o.ID
}

func (o *Order) Validate() error {

	if o.CustomerCPF != "" && len(o.CustomerCPF) == 11 {
		cpfValidator := validator.CPFValidator{}

		err := cpfValidator.Validate(o.CustomerCPF)
		if err != nil {
			return err
		}
	}

	if len(o.Items) == 0 {
		return errors.New("invalid order items")
	}

	return nil
}

func (o *Order) AddItem(item *OrderItem) {
	o.Items = append(o.Items, item)
}

func (o *Order) RemoveItem(item *OrderItem) {
	for i, v := range o.Items {
		if v == item {
			o.Items = append(o.Items[:i], o.Items[i+1:]...)
		}
	}
}

func (o *Order) CalculateTotal() {

	o.Total = 0

	for _, item := range o.Items {
		if item.Status == sharedconsts.OrderItemStatusConfirmed {
			o.Total += (item.Price * float64(item.Quantity))
		}
	}
}

func (o *Order) IsPaymentApproved() bool {
	name := o.Status.Name
	return name == sharedconsts.OrderStatusPaymentApproved
}

func (o *Order) InformPaymentApproval() {
	o.Status = *NewStatus(sharedconsts.OrderStatusPaymentApproved)
}

func (o *Order) InformPaymentNotApproval() {
	o.Status = *NewStatus(sharedconsts.OrderStatusNotApproved)
}

func (o *Order) Received() {
	o.Status = *NewStatus(sharedconsts.OrderReceivedByKitchen)
}

func (o *Order) InPreparation() {
	o.Status = *NewStatus(sharedconsts.OrderInPreparationByKitchen)
}

func (o *Order) Ready() {
	o.Status = *NewStatus(sharedconsts.OrderReadyByKitchen)
}

func (o *Order) Delivered() {
	o.Status = *NewStatus(sharedconsts.OrderFinalizedByKitchen)

}
