package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findbyid"
	"github.com/gin-gonic/gin"
)

type FindOrderByIDController struct {
	usecase portsusecase.FindOrderByIDUseCase
	ctx     context.Context
}

func NewFindOrderByIDController(ctx context.Context, usecase portsusecase.FindOrderByIDUseCase) *FindOrderByIDController {
	return &FindOrderByIDController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// @Summary Get a Order by id
// @Description Get details of a Order and their items by id
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param id path string true "Order id"
// @Success 200 {object} usecase.OrderFindByIdOutputDTO
// @Failure 404 {object} string "Order not found"
// @Failure 400 {object} string "Bad Request"
// @Router /orders/{id} [get]
func (cr *FindOrderByIDController) GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	order, err := cr.usecase.FindOrderByID(cr.ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
