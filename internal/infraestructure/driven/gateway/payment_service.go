package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	sharedurl "github.com/caiojorge/fiap-challenge-ddd/internal/shared/url"
	"github.com/joho/godotenv"
)

type IPaymentService interface {
	SendPaymentRequest(collectorID, posID string, payment *entity.Payment) error
}

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) SendPaymentRequest(collectorID, posID string, payment *entity.Payment) error {
	_ = godotenv.Load() // Carrega o .env se não estiver definido em variáveis de ambiente

	hostname := os.Getenv("APP_HOST_K8S")
	hostport := os.Getenv("HOST_PORT_CONTAINER")

	//url := fmt.Sprintf("http://%s:%s/instore/orders/qr/seller/collectors/%s/pos/%s/qrs", hostname, hostport, collectorID, posID)
	u := sharedurl.NewURL(hostname, hostport)
	url := u.GetPaymentURL(collectorID, posID)

	// Serializar o payload para JSON
	body, err := json.Marshal(payment)
	if err != nil {
		return fmt.Errorf("erro ao serializar o payload: %v", err)
	}

	// Criar a requisição POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("erro ao criar a requisição: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Cliente HTTP com timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Enviar a requisição
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao enviar a requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("requisição falhou com status: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("erro ao ler o corpo da resposta: %v", err)
	}

	var qrResponse QrResponse
	err = json.Unmarshal(respBody, &qrResponse)
	if err != nil {
		return fmt.Errorf("erro ao desserializar a resposta: %v", err)
	}

	// Registra as respostas enviadas pelo gateway
	payment.QrData = qrResponse.QrData
	payment.InStoreOrderID = qrResponse.InStoreOrderID

	return nil
}
