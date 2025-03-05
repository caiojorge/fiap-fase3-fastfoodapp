package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/shared"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/register"
	"github.com/gin-gonic/gin"
)

var ErrAlreadyExists = errors.New("product already exists")

type RegisterProductController struct {
	usecase portsusecase.RegisterProductUseCase
	ctx     context.Context
}

func NewRegisterProductController(ctx context.Context, usecase portsusecase.RegisterProductUseCase) *RegisterProductController {
	return &RegisterProductController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostRegisterProduct godoc
// @Summary Cria um novo produto
// @Schemes http
// @Description Cria um novo produto; As categorias são fixas: Lanches, Bebidas, Acompanhamentos e Sobremesas
// @Tags Products
// @Accept json
// @Produce json
// @Param        request   body     usecase.RegisterProductInputDTO  true  "Novo produto"
// @Success 200 {object} usecase.RegisterProductOutputDTO "Criado com sucesso"
// @Failure 400 {object} shared.ErrorResponse "Invalid data format or missing fields"
// @Failure 409 {object} shared.ErrorResponse "Product already exists"
// @Failure 500 {object} shared.ErrorResponse "Internal server error"
// @Router /products [post]
func (r *RegisterProductController) PostRegisterProduct(c *gin.Context) {
	var input portsusecase.RegisterProductInputDTO

	if err := c.BindJSON(&input); err != nil {
		errorResponse := shared.ErrorResponse{
			Message: "Invalid data",
			Code:    http.StatusBadRequest,
		}

		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Nesse cenário, o ID informado será ignorado e um novo ID será gerado
	fmt.Println("controller: Criando product: " + input.Name)
	output, err := r.usecase.RegisterProduct(r.ctx, &input)
	if err != nil {
		if err == ErrAlreadyExists {
			errorResponse := shared.ErrorResponse{
				Message: "product already exists",
				Code:    http.StatusConflict,
			}
			c.JSON(http.StatusConflict, errorResponse)
		} else {
			errorResponse := shared.ErrorResponse{
				Message: "error: " + err.Error(),
				Code:    http.StatusInternalServerError,
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
		}
		return
	}

	log.Println("Product created: ", output)

	c.JSON(http.StatusOK, output)
}
