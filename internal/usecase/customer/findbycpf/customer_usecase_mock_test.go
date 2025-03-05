package usecase

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCustomerFindByCPF(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocksrepository.NewMockCustomerRepository(ctrl)

	// mockRepo.EXPECT().
	// 	Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Customer{})).
	// 	Return(nil)

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := entity.NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	// Configurando o mock para retornar `nil` na primeira chamada

	mockRepo.EXPECT().
		Find(gomock.Any(), gomock.Eq("123.456.789-09")).
		Return(&entity.Customer{
			CPF:   valueobject.CPF{Value: "123.456.789-09"},
			Name:  "John Doe",
			Email: "email@email.com",
		}, nil).
		Times(1)

	finder := NewCustomerFindByCPF(mockRepo)
	assert.NotNil(t, finder)

	customer2, err := finder.FindCustomerByCPF(context.Background(), "123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, customer2)
	assert.Equal(t, customer.CPF.Value, customer2.CPF)

}

// type MockCustomerRepository struct {
// 	mu        sync.Mutex
// 	customers map[string]*entity.Customer
// }

// // NewMockCustomerRepository cria uma nova instância de um MockCustomerRepository.
// func NewMockCustomerRepository() *MockCustomerRepository {
// 	return &MockCustomerRepository{
// 		customers: make(map[string]*entity.Customer),
// 	}
// }

// // Create simula a criação de um novo cliente no repositório.
// func (repo *MockCustomerRepository) Create(ctx context.Context, customer *entity.Customer) error {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	if _, exists := repo.customers[customer.GetCPF().Value]; exists {
// 		return errors.New("customer already exists")
// 	}

// 	repo.customers[customer.GetCPF().Value] = customer
// 	return nil
// }

// // Update simula a atualização de um cliente no repositório.
// func (repo *MockCustomerRepository) Update(ctx context.Context, customer *entity.Customer) error {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	if _, exists := repo.customers[customer.GetCPF().Value]; !exists {
// 		return errors.New("customer not found")
// 	}

// 	repo.customers[customer.GetCPF().Value] = customer
// 	return nil
// }

// // Find simula a recuperação de um cliente pelo ID.
// func (repo *MockCustomerRepository) Find(ctx context.Context, id string) (*entity.Customer, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	customer, exists := repo.customers[id]
// 	if !exists {
// 		return nil, errors.New("customer not found")
// 	}
// 	return customer, nil
// }

// // FindAll simula a recuperação de uma lista de clientes.
// func (repo *MockCustomerRepository) FindAll(ctx context.Context) ([]*entity.Customer, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	var customers []*entity.Customer
// 	for _, customer := range repo.customers {
// 		customers = append(customers, customer)
// 	}
// 	return customers, nil
// }
