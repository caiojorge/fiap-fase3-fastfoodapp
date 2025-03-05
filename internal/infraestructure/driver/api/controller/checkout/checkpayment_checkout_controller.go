package controller

import (
	"context"
	"errors"
	"net/http"

	customerrors "github.com/caiojorge/fiap-challenge-ddd/internal/shared/error"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/checkpayment"
	"github.com/gin-gonic/gin"
)

//var someError = errors.New("order already exists")

type CheckPaymentCheckoutController struct {
	ctx     context.Context
	usecase portsusecase.ICheckPaymentUseCase
}

func NewCheckPaymentCheckoutController(ctx context.Context,
	usecase portsusecase.ICheckPaymentUseCase) *CheckPaymentCheckoutController {
	return &CheckPaymentCheckoutController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// @Summary Check Payment by Order ID
// @Description Get details of an Checkout and Status of Payment by Order id. Req #2 - Consultar status de pagamento pedido, que informa se o pagamento foi aprovado ou não.
// @Tags Checkouts
// @Accept  json
// @Produce  json
// @Param id path string true "Order id"
// @Success 200 {object} usecase.CheckPaymentOutputDTO
// @Failure 404 {object} string "Order | Checkout not found"
// @Failure 400 {object} string "Order ID is required"
// @Failure 500 {object} string "Internal Server Error"
// @Router /checkouts/{id}/check/payment [get]
func (r *CheckPaymentCheckoutController) GetCheckPaymentCheckout(c *gin.Context) {
	// a.II Consultar status de pagamento pedido, que informa se o pagamento foi aprovado ou não.
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrOrderIDIsRequired.Error()})
		return
	}

	output, err := r.usecase.CheckPayment(r.ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, customerrors.ErrCheckoutNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if output == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customerrors.ErrCheckoutNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
