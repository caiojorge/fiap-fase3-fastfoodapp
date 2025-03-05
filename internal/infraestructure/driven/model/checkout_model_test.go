package model

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateCheckoutWithExistingCustomer(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//db := setupMysql()

	// Migrar o esquema
	err = db.AutoMigrate(&Customer{}, &Product{}, &Order{}, &OrderItem{})
	if err != nil {
		panic("failed to migrate database")
	}

	// customer
	customer := Customer{
		CPF:   "75419654059", //75419654059
		Name:  "John Doe",
		Email: "email@email.com",
	}

	product := Product{
		ID:          "3",
		Name:        "Product 3",
		Description: "Description 3",
		Category:    "Category 3",
		Price:       10.0,
	}

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic("failed to load location")
	}

	orderItem := OrderItem{
		ID:        "3",
		ProductID: "3",
		Product:   product,
		Quantity:  1,
		Price:     10.0,
		Status:    "pending",
	}

	status := "pending"
	order := Order{
		ID: "3",
		Items: []*OrderItem{
			&orderItem,
		},
		Total:       10.0,
		Status:      status,
		CustomerCPF: &customer.CPF,
		Customer:    &customer,
		CreatedAt:   time.Now().In(location),
	}

	log.Default().Println(order)

	result := db.Create(&order)
	assert.Nil(t, result.Error)

	var order2 Order
	db.Find(&order2, "id = ?", "3")

	assert.Equal(t, order.ID, order2.ID)

	// var retrievedOrder Order
	// err = db.Preload("Status").Preload("Customer").Preload("Items.Product").First(&retrievedOrder, "id = ?", "3").Error
	// if err != nil {
	// 	t.Fatalf("failed to retrieve order: %v", err)
	// }

}
