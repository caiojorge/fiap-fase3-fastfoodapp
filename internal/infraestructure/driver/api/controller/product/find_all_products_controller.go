package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findall"
	"github.com/gin-gonic/gin"
)

type FindAllProductController struct {
	usecase portsusecase.FindAllProductsUseCase
	ctx     context.Context
}

func NewFindAllProductController(ctx context.Context, usecase portsusecase.FindAllProductsUseCase) *FindAllProductController {
	return &FindAllProductController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetAllProducts returns a list of all products
// @Summary Get all products
// @Description Get details of all products
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {array} usecase.FindAllProductOutputDTO
// @Failure 400 {object} map[string]string "Invalida data"
// @Failure 404 {object} map[string]string "No products foundr"
// @Router /products [get]
func (cr *FindAllProductController) GetAllProducts(c *gin.Context) {

	outputs, err := cr.usecase.FindAllProducts(cr.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(outputs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No products found"})
		return
	}

	c.JSON(http.StatusOK, outputs)
}
