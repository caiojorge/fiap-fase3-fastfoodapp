package controller

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/register"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPostRegisterCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocksrepository.NewMockCustomerRepository(ctrl)

	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Customer{})).
		Return(nil)

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := entity.NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	mockRepo.EXPECT().
		Find(gomock.Any(), customer.CPF.Value).
		Return(nil, nil).
		Times(1)

	//repo := NewMockCustomerRepository()
	//mock := NewMockRegisterCustomerUseCase(mockRepo)

	mock := usecase.NewCustomerRegister(mockRepo)

	controller := NewRegisterCustomerController(context.Background(), mock)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.POST("/register", controller.PostRegisterCustomer)

	// Create a JSON body
	requestBody := bytes.NewBuffer([]byte(`{"cpf":"123.456.789-09", "name":"John Doe","email":"email@email.com"}`))

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
	assert.Equal(t, http.StatusOK, w.Code, "Expected response code to be 200")
	//assert.Contains(t, w.Body.String(), "customer created John Doe", "Response body should contain correct customer name")
}
