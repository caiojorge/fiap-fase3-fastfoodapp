package model

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateOrder(t *testing.T) {
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

	orderItem := OrderItem{
		ID:        "1",
		ProductID: "1",
		Product:   product,
		Quantity:  1,
		Price:     10.0,
		Status:    "pending",
	}

	status := "qq status"
	order := Order{
		ID: "1",
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

}

func TestCreateOrderWithoutCustomer(t *testing.T) {

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

	product := Product{
		ID:          "2",
		Name:        "Product 2",
		Description: "Description 2",
		Category:    "Category 2",
		Price:       10.0,
	}

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic("failed to load location")
	}

	orderItem := OrderItem{
		ID:        "2",
		ProductID: "2",
		Product:   product,
		Quantity:  1,
		Price:     10.0,
		Status:    "pending",
	}

	status := "qq status"

	order := Order{
		ID: "2",
		Items: []*OrderItem{
			&orderItem,
		},
		Total:       10.0,
		Status:      status,
		CustomerCPF: nil,
		Customer:    nil,
		CreatedAt:   time.Now().In(location),
	}

	log.Default().Println(order)

	result := db.Create(&order)
	assert.Nil(t, result.Error)

}

func TestCreateOrderWithExistingCustomer(t *testing.T) {
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

	status := "qq status"
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

}
