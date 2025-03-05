package controller

import (
	"context"
	"errors"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/create"
	"github.com/gin-gonic/gin"
)

var ErrAlreadyExists = errors.New("order already exists")

type CreateCheckoutController struct {
	ctx     context.Context
	usecase portsusecase.ICreateCheckoutUseCase
}

func NewCreateCheckoutController(ctx context.Context,
	usecase portsusecase.ICreateCheckoutUseCase) *CreateCheckoutController {
	return &CreateCheckoutController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostCreateCheckout godoc
// @Summary Cria o checkout da ordem (inicia o processo de pagamento e comunicação com o gateway)
// @Schemes
// @Description Efetiva o pagamento do cliente, via fake checkout nesse momento, e deixa o pedindo em espera da confirmação do pagamento. A ordem muda de status nesse momento para checkout-confirmado. Req #1 - Checkout Pedido que deverá receber os produtos solicitados e retornar à identificação do pedido.
// @Tags Checkouts
// @Accept json
// @Produce json
// @Param        request   body     usecase.CheckoutInputDTO  true  "cria novo Checkout"
// @Success 200 {object} usecase.CheckoutOutputDTO
// @Failure 400 {object} string "invalid data"
// @Failure 500 {object} string "internal server error"
// @Router /checkouts [post]
func (r *CreateCheckoutController) PostCreateCheckout(c *gin.Context) {
	//// a.I Checkout Pedido que deverá receber os produtos solicitados e retornar à identificação do pedido
	var dto portsusecase.CheckoutInputDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	output, err := r.usecase.CreateCheckout(r.ctx, &dto)
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
