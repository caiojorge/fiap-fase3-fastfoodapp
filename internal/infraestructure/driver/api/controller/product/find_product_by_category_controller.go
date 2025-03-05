package controller

import (
	"context"
	"log"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findbycategory"
	"github.com/gin-gonic/gin"
)

type FindProductByCategoryController struct {
	usecase portsusecase.FindProductByCategoryUseCase
	ctx     context.Context
}

func NewFindProductByCategoryController(ctx context.Context, usecase portsusecase.FindProductByCategoryUseCase) *FindProductByCategoryController {
	return &FindProductByCategoryController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// @Summary Busca produtos por categoria
// @Description Busca produtos por categoria, ajusta a categoria para o plural e a primeira letra maiúscula
// @Tags Products
// @Accept  json
// @Produce  json
// @Param category path string true "Busca pelo nome da categoria; São aceitos apenas: Lanches, Bebidas, Sobremesas e Acompanhamentos"
// @Success 200 {array} usecase.FindProductByCategoryOutputDTO
// @Failure 404 {object} string "Product not found"
// @Failure 500 {object} string "Problems processing the request"
// @Router /products/category/{id} [get]
func (cr *FindProductByCategoryController) GetProductByCategory(c *gin.Context) {
	//id, ok := c.GetQuery("id")
	category := c.Param("id")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	log.Println("category: ", category)

	products, err := cr.usecase.FindProductByCategory(cr.ctx, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No products found"})
		return
	}

	c.JSON(http.StatusOK, products)
}
