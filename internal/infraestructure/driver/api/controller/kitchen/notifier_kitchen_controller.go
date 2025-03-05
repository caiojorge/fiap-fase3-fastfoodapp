package controller

import (
	"context"
	"net/http"

	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/notifier"
	"github.com/gin-gonic/gin"
)

type NotifyKitchenController struct {
	usecase usecase.IKitchenNotifierUseCase
	ctx     context.Context
}

func NewNotifyKitchenController(ctx context.Context, usecase usecase.IKitchenNotifierUseCase) *NotifyKitchenController {
	return &NotifyKitchenController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostNotifyKitchen Notifica a cozinha
// @Summary Busca ordens pagas e não notificadas, notifica a cozinha e atualiza o status da ordem para recebida pela cozinha. A ideia é que algum job chame esse endpoint para ficar buscando ordens para a cozinha
// @Description Retorna as ordens pagas e não notificadas, notifica a cozinha e atualiza o status da ordem para recebida pela cozinha
// @Tags Kitchens
// @Accept  json
// @Produce  json
// @Success 200 {array} usecase.KitchenNotifierOutputDTO
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /kitchens/orders/notifier [post]
func (r *NotifyKitchenController) PostNotifyKitchen(c *gin.Context) {

	outputs, err := r.usecase.Notify(r.ctx)
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
