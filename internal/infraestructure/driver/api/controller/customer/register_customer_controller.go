package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	portsusecaseregister "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/register"
	"github.com/gin-gonic/gin"
)

var ErrCustomerAlreadyExists = errors.New("customer already exists")

type RegisterCustomerController struct {
	usecase portsusecaseregister.RegisterCustomerUseCase
	ctx     context.Context
}

func NewRegisterCustomerController(ctx context.Context, usecase portsusecaseregister.RegisterCustomerUseCase) *RegisterCustomerController {
	return &RegisterCustomerController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostRegisterCustomer godoc
// @Summary Create Customer
// @Schemes
// @Description Create Customer in DB
// @Tags Customers
// @Accept json
// @Produce json
// @Param        request   body     usecase.CustomerRegisterInputDTO  true  "cria novo cliente"
// @Success 200 {object} usecase.CustomerRegisterOutputDTO
// @Failure 400 {object} map[string]string "invalid data"
// @Failure 409 {object} map[string]string "customer already exists"
// @Failure 500 {object} map[string]string "internal server error"
// @Router /customers [post]
func (r *RegisterCustomerController) PostRegisterCustomer(c *gin.Context) {
	var dto portsusecaseregister.CustomerRegisterInputDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	fmt.Println("controller: Criando cliente: " + dto.CPF)
	err := r.usecase.RegisterCustomer(r.ctx, dto)
	if err != nil {
		if err == ErrCustomerAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Customer already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, dto)
}
