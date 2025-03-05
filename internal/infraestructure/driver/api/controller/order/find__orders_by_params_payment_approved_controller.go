package controller

import (
	"context"
	"net/http"

	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findbyparam"
	"github.com/gin-gonic/gin"
)

type FindByParamsPaymentApprovedController struct {
	usecase portsusecase.FindOrderByParamsUseCase
	ctx     context.Context
}

func NewFindByParamsPaymentApprovedController(ctx context.Context, usecase portsusecase.FindOrderByParamsUseCase) *FindByParamsPaymentApprovedController {
	return &FindByParamsPaymentApprovedController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetOrdersWithPaymentApproved returns a list of all paid orders
// @Summary Get all confirmed orders
// @Description Retorna todos os pedidos (orders) registrados no sistema que tenham status de pagamento confirmado. Se não houver pedidos, retorna um erro (404).
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} usecase.OrderFindByParamOutputDTO
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /orders/paid [get]
func (r *FindByParamsPaymentApprovedController) GetOrdersWithPaymentApproved(c *gin.Context) {

	orders, err := r.usecase.FindOrdersByParams(r.ctx, map[string]interface{}{"status": sharedconsts.OrderStatusPaymentApproved})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
