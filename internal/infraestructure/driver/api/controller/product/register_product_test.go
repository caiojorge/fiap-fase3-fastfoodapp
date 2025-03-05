package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	usecasefindall "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findall"
	usecasefindbyid "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findbyid"
	usecaseregister "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/register"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
)

func TestRegisterProductController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocksrepository.NewMockProductRepository(ctrl)

	mockProductRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Product{})).
		Return(nil)

	mockProductRepo.EXPECT().
		FindByName(gomock.Any(), "Lanche XPTO").
		Return(nil, nil)

	//repo := NewMockProductRepository()
	mock := usecaseregister.NewProductRegister(mockProductRepo)
	controller := NewRegisterProductController(context.Background(), mock)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.POST("/register", controller.PostRegisterProduct)

	// Create a JSON body
	requestBody := bytes.NewBuffer([]byte(`{"name":"Lanche XPTO","description":"Pão, carne, queijo e presunto","category":"Lanches","price": 100}`))

	// Create the HTTP request with JSON body
	req, err := http.NewRequest("POST", "/register", requestBody)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println("Body:", w.Body.String())

	var dto usecaseregister.RegisterProductOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &dto)
	assert.Nil(t, err)

	assert.Equal(t, "Lanche XPTO", dto.Name)
}

func TestFindAllProductsController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := &entity.Product{
		Name:        "Lanche XPTO",
		Description: "Pão, carne, queijo e presunto",
		Category:    "Lanches",
		Price:       100,
	}

	mockProductRepo := mocksrepository.NewMockProductRepository(ctrl)

	mockProductRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Product{})).
		Return(nil)

	mockProductRepo.EXPECT().
		FindByName(gomock.Any(), "Lanche XPTO").
		Return(nil, nil)

	mockProductRepo.EXPECT().
		FindAll(gomock.Any()).
		Return([]*entity.Product{product}, nil)

	//repo := NewMockProductRepository()
	usecase := usecasefindall.NewProductFindAll(mockProductRepo)
	controller := NewFindAllProductController(context.Background(), usecase)

	register := usecaseregister.NewProductRegister(mockProductRepo)
	productDto := usecaseregister.RegisterProductInputDTO{
		Name:        "Lanche XPTO",
		Description: "Pão, carne, queijo e presunto",
		Category:    "Lanches",
		Price:       100,
	}

	_, err := register.RegisterProduct(context.Background(), &productDto)
	assert.Nil(t, err)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.GET("/products", controller.GetAllProducts)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)

	var dto []usecasefindall.FindAllProductOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &dto)
	assert.Nil(t, err)

	assert.Equal(t, "Lanche XPTO", dto[0].Name)

}

func TestFindProductByIDController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocksrepository.NewMockProductRepository(ctrl)

	mockProductRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Product{})).
		Return(nil)

	mockProductRepo.EXPECT().
		FindByName(gomock.Any(), "Lanche XPTO").
		Return(nil, nil)

	//repo := NewMockProductRepository()
	mock := usecasefindbyid.NewProductFindByID(mockProductRepo)
	controller := NewFindProductByIDController(context.Background(), mock)

	register := usecaseregister.NewProductRegister(mockProductRepo)
	productDto := usecaseregister.RegisterProductInputDTO{
		Name:        "Lanche XPTO",
		Description: "Pão, carne, queijo e presunto",
		Category:    "Lanches",
		Price:       100,
	}

	output, err := register.RegisterProduct(context.Background(), &productDto)
	assert.Nil(t, err)

	product := &entity.Product{
		ID:          output.ID,
		Name:        "Lanche XPTO",
		Description: "Pão, carne, queijo e presunto",
		Category:    "Lanches",
		Price:       100,
	}

	mockProductRepo.EXPECT().
		Find(gomock.Any(), output.ID).
		Return(product, nil)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.GET("/products/:id", controller.GetProductByID)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/products/"+output.ID, nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)

}
