package controller

import (
	"context"
	"net/http"

	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/delivery"
	"github.com/gin-gonic/gin"
)

type DeliveryKitchenController struct {
	usecase usecase.IKitchenDeliveryUseCase
	ctx     context.Context
}

func NewDeliveryKitchenController(ctx context.Context, usecase usecase.IKitchenDeliveryUseCase) *DeliveryKitchenController {
	return &DeliveryKitchenController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostDeliveryKitchen move o status para finalizado
// @Summary Busca ordem que esta pronta, registra o delivery e finaliza a ordem
// @Description Busca ordem que esta pronta, registra o delivery e finaliza a ordem
// @Tags Kitchens
// @Accept  json
// @Produce  json
// @Param        request   body     usecase.KitchenDeliveryInputDTO  true  "indica a ordem que será finalizada"
// @Success 200 {array} usecase.KitchenDeliveryOutputDTO
// @Failure 400 {object} string "Bad Request"
// @Router /kitchens/orders/delivery [post]
func (r *DeliveryKitchenController) PostDeliveryKitchen(c *gin.Context) {
	var input usecase.KitchenDeliveryInputDTO

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	output, err := r.usecase.Delivery(r.ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
