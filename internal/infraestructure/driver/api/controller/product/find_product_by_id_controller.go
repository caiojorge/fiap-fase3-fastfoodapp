package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findbyid"
	"github.com/gin-gonic/gin"
)

type FindProductByIDController struct {
	usecase portsusecase.FindProductByIDUseCase
	ctx     context.Context
}

func NewFindProductByIDController(ctx context.Context, usecase portsusecase.FindProductByIDUseCase) *FindProductByIDController {
	return &FindProductByIDController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// @Summary Get a Product
// @Description Get details of a Product by id
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path string true "Product id"
// @Success 200 {object} usecase.FindProductByIDOutputDTO
// @Failure 404 {object} string "Product not found"
// @Router /products/{id} [get]
func (cr *FindProductByIDController) GetProductByID(c *gin.Context) {
	//id, ok := c.GetQuery("id")
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	product, err := cr.usecase.FindProductByID(cr.ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}
