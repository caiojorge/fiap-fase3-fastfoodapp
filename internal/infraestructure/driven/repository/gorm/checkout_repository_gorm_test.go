package repositorygorm

import (
	"context"
	"testing"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateCheckout(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//db := setupMysql()

	// Migrar o esquema
	err = db.AutoMigrate(&model.Customer{}, &model.Product{}, &model.Order{}, &model.OrderItem{}, &model.Checkout{})
	if err != nil {
		panic("failed to migrate database")
	}

	converter := converter.NewOrderConverter()
	repo := NewOrderRepositoryGorm(db, converter)

	// customer
	customer := model.Customer{
		CPF:   "75419654059", //75419654059
		Name:  "John Doe",
		Email: "email@email.com",
	}

	product := model.Product{
		ID:          "1",
		Name:        "Product 1",
		Description: "Description 1",
		Category:    "Category 1",
		Price:       10.0,
	}

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic("failed to load location")
	}

	orderItem := model.OrderItem{
		ID:        "111",
		ProductID: "1",
		Product:   product,
		Quantity:  1,
		Price:     10.0,
		Status:    "pending",
	}

	order := model.Order{
		ID: "111",
		Items: []*model.OrderItem{
			&orderItem,
		},
		Total:       10.0,
		Status:      "pending",
		CustomerCPF: &customer.CPF,
		Customer:    &customer,
		CreatedAt:   time.Now().In(location),
	}

	entity := converter.ToEntity(&order)
	err = repo.Create(context.Background(), entity)
	assert.Nil(t, err)

	checkout := model.Checkout{
		ID:                   "111",
		OrderID:              order.ID,
		GatewayName:          "pagseguro",
		GatewayToken:         "token123",
		GatewayTransactionID: "trans123",
		Total:                order.Total,
		CreatedAt:            time.Now().In(location),
	}

	checkoutRepo := NewCheckoutRepositoryGorm(db)
	checkoutEntity := checkout.ToEntity()
	err = checkoutRepo.Create(context.Background(), checkoutEntity)
	assert.Nil(t, err)
	assert.NotNil(t, checkoutEntity.ID)
	assert.NotNil(t, checkoutEntity.OrderID)
	assert.Equal(t, order.ID, checkoutEntity.OrderID)

}
