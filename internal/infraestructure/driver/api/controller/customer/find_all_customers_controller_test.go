package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	//"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/findall"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestGetCustomerByCPF tests the GetCustomerByCPF handler for both valid and invalid requests.
func TestGetAllCustomers(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocksrepository.NewMockCustomerRepository(ctrl)

	cpf1, _ := valueobject.NewCPF("400.228.165-50")
	cpf2, _ := valueobject.NewCPF("364.584.534-85")

	mockRepo.EXPECT().
		FindAll(gomock.Any()).
		Return([]*entity.Customer{
			{CPF: *cpf1, Name: "John Doe", Email: "1email@email.com"},
			{CPF: *cpf2, Name: "Jane Doe", Email: "2email@email.com"},
		}, nil)

	mock := usecase.NewCustomerFindAll(mockRepo)

	controller := NewFindAllCustomersController(context.Background(), mock)

	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register the handler
	r.GET("/customer", controller.GetAllCustomers)

	req, _ := http.NewRequest(http.MethodGet, "/customer", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Now parse the JSON body into the slice of Customer structs
	var customers []*usecase.CustomerFindAllOutputDTO
	err := json.Unmarshal(resp.Body.Bytes(), &customers)
	assert.NoError(t, err, "Should decode response without error")

	// Check the length of the returned slice to ensure you received two customers
	assert.Len(t, customers, 2, "Expected 2 customers in response")

	assert.Equal(t, "400.228.165-50", customers[0].CPF)
	assert.Equal(t, "John Doe", customers[0].Name)

	assert.Equal(t, "364.584.534-85", customers[1].CPF)
	assert.Equal(t, "Jane Doe", customers[1].Name)

}

// type MockFindAllCustomersUseCase struct {
// 	repository portsrepository.CustomerRepository
// }

// func NewMockFindAllCustomersUseCase(repository portsrepository.CustomerRepository) *MockFindAllCustomersUseCase {
// 	return &MockFindAllCustomersUseCase{
// 		repository: repository,
// 	}
// }

// func (m *MockFindAllCustomersUseCase) FindAllCustomers(ctx context.Context) ([]*entity.Customer, error) {

// 	cpf1, _ := valueobject.NewCPF("400.228.165-50")
// 	cpf2, _ := valueobject.NewCPF("364.584.534-85")

// 	expectedCustomers := []*entity.Customer{
// 		{CPF: *cpf1, Name: "John Doe", Email: "john@example.com"},
// 		{CPF: *cpf2, Name: "Jane Doe", Email: "jane@example.com"},
// 	}

// 	return expectedCustomers, nil
// }
