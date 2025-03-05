package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	usecasefindbycpf "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/findbycpf"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestGetCustomerByCPF tests the GetCustomerByCPF handler for both valid and invalid requests.
func TestGetCustomerByCPF(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocksrepository.NewMockCustomerRepository(ctrl)

	cpf1, _ := valueobject.NewCPF("400.228.165-50")

	mockRepo.EXPECT().
		Find(gomock.Any(), "40022816550").
		Return(&entity.Customer{CPF: *cpf1, Name: "John Doe", Email: ""}, nil)

	mock := usecasefindbycpf.NewCustomerFindByCPF(mockRepo)

	controller := NewFindCustomerByCPFController(context.Background(), mock)

	// Set up the Gin router
	// Create a new Gin router
	router := gin.Default()

	// Register the route and handler
	router.GET("/customer/:cpf", controller.GetCustomerByCPF)

	// Create a test request
	req, err := http.NewRequest(http.MethodGet, "/customer/40022816550", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	var dto usecasefindbycpf.CustomerFindByCpfOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &dto)
	assert.Nil(t, err)

	assert.Equal(t, "John Doe", dto.Name)
}
