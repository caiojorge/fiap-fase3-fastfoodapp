package controller

import (
	"context"
	"net/http"

	sharedconsts "github.com/caiojorge/fiap-challenge-ddd/internal/shared/consts"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findbyparam"
	"github.com/gin-gonic/gin"
)

type FindByParamsNotConfirmedController struct {
	usecase portsusecase.FindOrderByParamsUseCase
	ctx     context.Context
}

func NewFindByParamsNotConfirmedController(ctx context.Context, usecase portsusecase.FindOrderByParamsUseCase) *FindByParamsNotConfirmedController {
	return &FindByParamsNotConfirmedController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetOrdersNotConfirmed returns a list of all paid orders
// @Summary Get all orders not confirmed
// @Description Retorna todos os pedidos (orders) registrados no sistema sem o pagamento confirmado. Se não houver pedidos, retorna um erro (404).
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} usecase.OrderFindByParamOutputDTO
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /orders/pending [get]
func (r *FindByParamsNotConfirmedController) GetOrdersNotConfirmed(c *gin.Context) {

	orders, err := r.usecase.FindOrdersByParams(r.ctx, map[string]interface{}{"status": sharedconsts.OrderStatusNotConfirmed})
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
