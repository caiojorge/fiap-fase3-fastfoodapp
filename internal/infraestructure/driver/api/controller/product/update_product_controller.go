package controller

import (
	"context"
	"net/http"

	updateusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/update"
	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	usecase updateusecase.UpdateProductUseCase
	ctx     context.Context
}

func NewUpdateProductController(ctx context.Context, usecase updateusecase.UpdateProductUseCase) *UpdateProductController {
	return &UpdateProductController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PutUpdateProduct updates a Product by id
// @Summary Update a Product
// @Description Update details of a Product by id
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path string true "Product id"
// @Param Product body usecase.UpdateProductInputDTO true "Product data"
// @Success 200 {object} usecase.UpdateProductOutputDTO "Product updated"
// @Failure 400 {object} string "Invalid data"
// @Failure 404 {object} string "Product not found"
// @Router /products/{id} [put]
func (r *UpdateProductController) PutUpdateProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var dto updateusecase.UpdateProductInputDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if dto.ID == "" || dto.ID != id {
		dto.ID = id
	}

	output, err := r.usecase.UpdateProduct(r.ctx, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
