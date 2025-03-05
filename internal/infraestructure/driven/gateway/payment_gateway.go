package service

import (
	"context"
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
)

type QrResponse struct {
	QrData            string `json:"qr_data"`
	InStoreOrderID    string `json:"in_store_order_id"`
	ExternalReference string `json:"external_reference"`
}

// PaymentGateway provides methods for payment operations.
// Vai se conectar com o gateway de pagamento, nesse caso, FAKE.
type PaymentGateway struct {
	paymentService IPaymentService
}

func NewPaymentGateway(paymentService IPaymentService) *PaymentGateway {
	return &PaymentGateway{paymentService: paymentService}
}

// CreateCheckout creates a new checkout. This method should be implemented by the payment gateway.
func (p *PaymentGateway) ConfirmPayment(ctx context.Context, checkout *entity.Checkout, order *entity.Order, productList []*entity.Product, notificationURL string, sponsorID int) (*entity.Payment, error) {
	payment, err := entity.NewPayment(*checkout, *order, productList, notificationURL, sponsorID)
	if err != nil {
		return nil, err
	}

	// a.I Checkout Pedido que deverá receber os produtos solicitados e retornar à identificação do pedido
	//err = sendPaymentRequest("collector123", "pos456", payment)
	err = p.paymentService.SendPaymentRequest("collector123", "pos456", payment)
	if err != nil {
		fmt.Printf("Erro ao enviar pagamento: %v\n", err)
		return nil, err
	}

	return payment, nil
}

// CancelTransaction cancels a transaction. This method should be implemented by the payment gateway.
func (p *PaymentGateway) CancelPayment(ctx context.Context, id string) error {
	return nil
}

// func sendPaymentRequest(collectorID, posID string, payment *entity.Payment) error {

// 	url := fmt.Sprintf("http://payment-api-url/instore/orders/qr/seller/collectors/%s/pos/%s/qrs", collectorID, posID)

// 	// Serializar o payload para JSON
// 	body, err := json.Marshal(payment)
// 	if err != nil {
// 		return fmt.Errorf("erro ao serializar o payload: %v", err)
// 	}

// 	// Criar a requisição POST
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return fmt.Errorf("erro ao criar a requisição: %v", err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	// Cliente HTTP com timeout
// 	client := &http.Client{Timeout: 10 * time.Second}

// 	// Enviar a requisição
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("erro ao enviar a requisição: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("requisição falhou com status: %s", resp.Status)
// 	}

// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("erro ao ler o corpo da resposta: %v", err)
// 	}

// 	var qrResponse QrResponse
// 	err = json.Unmarshal(respBody, &qrResponse)
// 	if err != nil {
// 		return fmt.Errorf("erro ao desserializar a resposta: %v", err)
// 	}

// 	// registra as respostas enviadas pelo gateway
// 	payment.QrData = qrResponse.QrData
// 	payment.InStoreOrderID = qrResponse.InStoreOrderID

// 	fmt.Println("Pagamento enviado com sucesso!")
// 	return nil
// }
