package controller

import (
	"context"
	"errors"
	"net/http"

	usecaseorder "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/create"
	"github.com/gin-gonic/gin"
)

var ErrAlreadyExists = errors.New("order already exists")

type CreateOrderController struct {
	usecase usecaseorder.CreateOrderUseCase
	ctx     context.Context
}

func NewCreateOrderController(ctx context.Context, usecase usecaseorder.CreateOrderUseCase) *CreateOrderController {
	return &CreateOrderController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostCreateOrder godoc
// @Summary Create Order
// @Schemes
// @Description Cria um peddo (order) no sistema. O cliente (customer) pode ou não de identificar. Se o cliente não se identificar, o pedido será registrado como anônimo. O produto, porém, deve ter sido previamente cadastrado.
// @Tags Orders
// @Accept json
// @Produce json
// @Param        request   body     usecase.OrderCreateInputDTO  true  "cria nova Order"
// @Success 200 {object} usecase.OrderCreateOutputDTO
// @Failure 400 {object} string "invalid data"
// @Failure 409 {object} string "Order already exists"
// @Failure 500 {object} string "internal server error"
// @Router /orders [post]
func (r *CreateOrderController) PostCreateOrder(c *gin.Context) {
	var input usecaseorder.OrderCreateInputDTO

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	output, err := r.usecase.CreateOrder(r.ctx, &input)
	if err != nil {
		if err == ErrAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Order already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, output)
}
