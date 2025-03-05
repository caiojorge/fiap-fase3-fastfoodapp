package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/update"
	"github.com/gin-gonic/gin"
)

type UpdateCustomerController struct {
	usecase usecase.UpdateCustomerUseCase
	ctx     context.Context
}

func NewUpdateCustomerController(ctx context.Context, usecase usecase.UpdateCustomerUseCase) *UpdateCustomerController {
	return &UpdateCustomerController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PutUpdateCustomer updates a customer by cpf
// @Summary Update a customer
// @Description Update details of a customer by cpf
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param cpf path string true "Customer cpf"
// @Param Customer body usecase.CustomerUpdateInputDTO true "Customer data"
// @Success 200 {object} usecase.CustomerUpdateOutputDTO
// @Failure 400 {object} map[string]string "Invalid data"
// @Failure 404 {object} map[string]string "Customer not found"
// @Router /customers/{cpf} [put]
func (r *UpdateCustomerController) PutUpdateCustomer(c *gin.Context) {
	cpf := c.Param("cpf")
	if cpf == "" {
		c.JSON(http.StatusBadRequest, gin.H{"param: error": "Invalid data"})
		return
	}

	var dto usecase.CustomerUpdateInputDTO
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"bind: error": "Invalid data"})
		return
	}

	dto.CPF = formatter.RemoveMaskFromCPF(cpf)

	err := r.usecase.UpdateCustomer(r.ctx, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"update: error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto)
}
