package controller

import (
	"context"
	"net/http"

	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/monitor"
	"github.com/gin-gonic/gin"
)

type MonitorKitchenController struct {
	usecase usecase.IMonitorKitchenUseCase
	ctx     context.Context
}

func NewMonitorKitchenController(ctx context.Context, usecase usecase.IMonitorKitchenUseCase) *MonitorKitchenController {
	return &MonitorKitchenController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetMonitorKitchen Notifica a cozinha
// @Summary Busca ordens que estão na cozinha e estão sendo preparadas (ou na fila de preparo)
// @Description Retorna uma lista de ordens e seus status, ordenado por recebido, em preparo e pronto, e por ordem de chegada tbm, e o tempo estimado e o delivery number para retirada do pedido
// @Tags Kitchens
// @Accept  json
// @Produce  json
// @Success 200 {array} usecase.MonitorKitchenOutputDTO
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /kitchens/orders/monitor [get]
func (r *MonitorKitchenController) GetMonitorKitchen(c *gin.Context) {

	outputs, err := r.usecase.Monitor(r.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(outputs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, outputs)
}
