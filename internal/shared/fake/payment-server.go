package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type PaymentRequest struct {
	ExternalReference string  `json:"external_reference"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	NotificationURL   string  `json:"notification_url"`
	TotalAmount       float64 `json:"total_amount"`
	Items             []Item  `json:"items"`
	Sponsor           Sponsor `json:"sponsor"`
	CashOut           CashOut `json:"cash_out"`
}

type Item struct {
	SKUNumber   string  `json:"sku_number"`
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float64 `json:"total_amount"`
}

type Sponsor struct {
	ID int `json:"id"`
}

type CashOut struct {
	Amount float64 `json:"amount"`
}

type QrResponse struct {
	QrData            string `json:"qr_data"`
	InStoreOrderID    string `json:"in_store_order_id"`
	ExternalReference string `json:"external_reference"`
}

type WebhookRequest struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

func Run() {
	_ = godotenv.Load() // Carrega o .env se não estiver definido em variáveis de ambiente

	hostname := os.Getenv("HOST2_NAME")
	hostport := os.Getenv("HOST2_PORT")

	if hostname == "" || hostport == "" {
		panic("Variáveis de ambiente HOST2_NAME e HOST2_PORT não definidas")
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any

	r := gin.Default()

	// Endpoint para receber pedidos de pagamento
	r.POST("/instore/orders/qr/seller/collectors/:collectorID/pos/:posID/qrs", PostPaymentFake)

	// Endpoint para retornar os pagamentos confirmados
	r.GET("/payments/confirmed", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Pagamentos confirmados",
		})
	})

	logger.Info("Server running on " + hostname + ":" + hostport)

	//r.Run(":8081") // Porta do serviço fake de pagamento
	r.Run(":" + hostport)
}

func PostPaymentFake(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gerando um ID único para a ordem de pagamento
	inStoreOrderID := uuid.New().String()
	//inStoreOrderID := req.ExternalReference
	fmt.Println("InStoreOrderID:", inStoreOrderID)

	// Simulando geração do QR Data
	qrData := fmt.Sprintf("00020101021243650016COM.MERCADOLIBRE020130%s-%s", req.ExternalReference, inStoreOrderID)

	// Retornando o QR Code e ID da ordem de pagamento
	c.JSON(http.StatusOK, QrResponse{
		QrData:            qrData,
		InStoreOrderID:    inStoreOrderID,
		ExternalReference: req.ExternalReference,
	})

	// Simulando pagamento após um tempo
	go func() {
		time.Sleep(15 * time.Second)
		callWebhook(req.NotificationURL, req.ExternalReference)
	}()
}

func callWebhook(notificationURL, orderID string) {
	webhookBody := WebhookRequest{
		OrderID: orderID,
		Status:  "approved",
	}

	jsonBody, _ := json.Marshal(webhookBody)

	req, err := http.NewRequest(http.MethodPut, notificationURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Erro ao criar requisição PUT:", err)
		return
	}

	// Define o header Content-Type como application/json
	req.Header.Set("Content-Type", "application/json")

	// Executa a requisição HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao chamar webhook:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Resposta do webhook:", resp.Status)
	fmt.Println("Webhook chamado com sucesso!")

}
