package controller

import (
	"context"
	"net/http"

	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/cooking"
	"github.com/gin-gonic/gin"
)

type CookingKitchenController struct {
	usecase usecase.IKitchenCookingUseCase
	ctx     context.Context
}

func NewCookingKitchenController(ctx context.Context, usecase usecase.IKitchenCookingUseCase) *CookingKitchenController {
	return &CookingKitchenController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostCookingKitchen move o status para o próximo da fase de preparo
// @Summary Busca a ordem e o ticket da cozinha; move o status para o próximo da fase de preparo e faz o delivery se estiver finalizado.
// @Description Busca a ordem e o ticket da cozinha; move o status para o próximo da fase de preparo e faz o delivery se estiver finalizado.
// @Tags Kitchens
// @Accept  json
// @Produce  json
// @Param        request   body     usecase.KitchenCookingInputDTO  true  "indica a ordem a ser trabalhada pela cozinha"
// @Success 200 {array} usecase.KitchenCookingOutputDTO
// @Failure 400 {object} string "Bad Request"
// @Router /kitchens/orders/cooking [post]
func (r *CookingKitchenController) PostCookingKitchen(c *gin.Context) {
	var input usecase.KitchenCookingInputDTO

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	output, err := r.usecase.Cook(r.ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
