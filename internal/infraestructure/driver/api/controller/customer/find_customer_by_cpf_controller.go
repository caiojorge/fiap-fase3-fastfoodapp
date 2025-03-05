package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/findbycpf"
	"github.com/gin-gonic/gin"
)

type FindCustomerByCPFController struct {
	usecase portsusecase.FindCustomerByCPFUseCase
	ctx     context.Context
}

func NewFindCustomerByCPFController(ctx context.Context, usecase portsusecase.FindCustomerByCPFUseCase) *FindCustomerByCPFController {
	return &FindCustomerByCPFController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// @Summary Get a customer
// @Description Get details of a customer by cpf
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param cpf path string true "Customer cpf"
// @Success 200 {object} usecase.CustomerFindByCpfOutputDTO
// @Failure 404 {object} map[string]string "Customer not found"
// @Router /customers/{cpf} [get]
func (cr *FindCustomerByCPFController) GetCustomerByCPF(c *gin.Context) {
	//cpf, ok := c.GetQuery("cpf")
	cpf := c.Param("cpf")

	if cpf == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	customer, err := cr.usecase.FindCustomerByCPF(cr.ctx, formatter.RemoveMaskFromCPF(cpf))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}
