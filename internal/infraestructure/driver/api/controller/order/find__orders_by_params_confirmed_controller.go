package controller

import (
	"context"
	"net/http"

	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findbyparam"
	"github.com/gin-gonic/gin"
)

type FindByParamsConfirmedController struct {
	usecase portsusecase.FindOrderByParamsUseCase
	ctx     context.Context
}

func NewFindByParamsConfirmedController(ctx context.Context, usecase portsusecase.FindOrderByParamsUseCase) *FindByParamsConfirmedController {
	return &FindByParamsConfirmedController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetConfirmedOrders returns a list of all confirmed orders
// @Summary Get all confirmed orders
// @Description Retorna todos os pedidos (orders) registrados no sistema. Se não houver pedidos, retorna um erro (404).
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} usecase.OrderFindByParamOutputDTO
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /orders/confirmed [get]
func (r *FindByParamsConfirmedController) GetOrdersConfirmed(c *gin.Context) {

	orders, err := r.usecase.FindOrdersByParams(r.ctx, map[string]interface{}{"status": sharedconsts.OrderStatusConfirmed})
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
