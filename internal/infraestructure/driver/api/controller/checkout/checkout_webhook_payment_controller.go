package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/confirmation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebhookCheckoutController struct {
	ctx     context.Context
	usecase portsusecase.ICheckoutConfirmationUseCase
	logger  *zap.Logger
}

func NewWebhookCheckoutController(ctx context.Context, usecase portsusecase.ICheckoutConfirmationUseCase, logger *zap.Logger) *WebhookCheckoutController {
	return &WebhookCheckoutController{
		ctx:     ctx,
		usecase: usecase,
		logger:  logger,
	}
}

// @Summary Webhook to confirm payment
// @Schemes
// @Description Confirma o pagamento do cliente, via fake checkout nesse momento, e libera o pedido para preparação. A ordem muda de status nesse momento, para pagamento aprovado. Req #1 - Webhook para receber confirmação de pagamento aprovado ou recusado. A implementação deve ser clara quanto ao Webhook.
// @Tags Checkouts
// @Accept json
// @Produce json
// @Param        request   body     usecase.CheckoutConfirmationInputDTO  true  "Webhook para finalizar o Checkout"
// @Success 200 {object} usecase.CheckoutConfirmationOutputDTO
// @Failure 400 {object} string "invalid data"
// @Failure 500 {object} string "internal server error"
// @Router /checkouts/confirmation/payment [post]
func (cw *WebhookCheckoutController) PutConfirmPayment(c *gin.Context) {
	cw.logger.Info("PutConfirmPayment")
	// a.III Webhook para receber confirmação de pagamento aprovado ou recusado. A implementação deve ser clara quanto ao Webhook.
	var input portsusecase.CheckoutConfirmationInputDTO

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// determina se o pagamento foi aprovado ou não (claro, com base na resposta do gateway de pagamento)
	output, err := cw.usecase.ConfirmPayment(cw.ctx, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if output == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to confirm payment"})
		return
	}

	c.JSON(http.StatusOK, output)
}
