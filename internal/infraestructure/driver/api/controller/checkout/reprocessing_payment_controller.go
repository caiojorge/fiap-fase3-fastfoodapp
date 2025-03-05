package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/reprocessing"
	"github.com/gin-gonic/gin"
)

type ReprocessingPaymentController struct {
	ctx     context.Context
	usecase portsusecase.ICheckoutReprocessingUseCase
}

func NewReprocessingPaymentController(ctx context.Context,
	usecase portsusecase.ICheckoutReprocessingUseCase) *ReprocessingPaymentController {
	return &ReprocessingPaymentController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostReprocessingPayment godoc
// @Summary Reprocessa o Checkout
// @Schemes
// @Description Reprocesso o pagamento. Ordens em checkout aprovado, com checkout criado podem ser reprocessadas. O reprocessamento é feito no gateway de pagamento. O checkout é atualizado com o status do pagamento e o pedido é notificado.
// @Tags Checkouts
// @Accept json
// @Produce json
// @Param        request   body     usecase.CheckoutReprocessingInputDTO  true  "reprocessa o Checkout"
// @Success 200 {object} usecase.CheckoutReprocessingOutputDTO
// @Failure 400 {object} string "invalid data"
// @Failure 500 {object} string "internal server error"
// @Router /checkouts/reprocessing/payment [post]
func (r *ReprocessingPaymentController) PostReprocessingPayment(c *gin.Context) {
	//// a.I Checkout Pedido que deverá receber os produtos solicitados e retornar à identificação do pedido
	var dto portsusecase.CheckoutReprocessingInputDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	output, err := r.usecase.ReprocessPayment(r.ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if output == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction on gateway"})
		return
	}

	c.JSON(http.StatusOK, output)
}
