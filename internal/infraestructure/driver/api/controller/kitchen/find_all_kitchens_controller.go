package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/findall"
	"github.com/gin-gonic/gin"
)

type FindKitchenAllController struct {
	usecase portsusecase.FindAllKitchenUseCase
	ctx     context.Context
}

func NewFindKitchenAllController(ctx context.Context, usecase portsusecase.FindAllKitchenUseCase) *FindKitchenAllController {
	return &FindKitchenAllController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetAllOrders retorna todas as ordens que estão na cozinha
// @Summary retorna todas as ordens que estão na cozinha
// @Description Retorna as ordens que estão na cozinha em alguma etapa do preparo
// @Tags Kitchens
// @Accept  json
// @Produce  json
// @Success 200 {array} usecase.KitchenFindAllAOutputDTO
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /kitchens/orders/flow [get]
func (r *FindKitchenAllController) GetAllOrdersInKitchen(c *gin.Context) {

	entities, err := r.usecase.FindAllKitchen(r.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(entities) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, entities)
}
