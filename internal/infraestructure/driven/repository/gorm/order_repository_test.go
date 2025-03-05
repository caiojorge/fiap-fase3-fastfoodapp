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

func TestCreateOrder(t *testing.T) {

	t.Run("CreateOrder", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrar o esquema
		err = db.AutoMigrate(&model.Customer{}, &model.Product{}, &model.Order{}, &model.OrderItem{})
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

		status := "qq status"
		order := model.Order{
			ID: "111",
			Items: []*model.OrderItem{
				&orderItem,
			},
			Total:       10.0,
			Status:      status,
			CustomerCPF: &customer.CPF,
			Customer:    &customer,
			CreatedAt:   time.Now().In(location),
		}

		entity := converter.ToEntity(&order)
		err = repo.Create(context.Background(), entity)
		assert.Nil(t, err)
	})

	t.Run("UpdateOrder and UpdateStatus", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrar o esquema
		err = db.AutoMigrate(&model.Customer{}, &model.Product{}, &model.Order{}, &model.OrderItem{})
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

		status := "approved"
		order := model.Order{
			ID: "111",
			Items: []*model.OrderItem{
				&orderItem,
			},
			Total:       10.0,
			Status:      status,
			CustomerCPF: &customer.CPF,
			Customer:    &customer,
			CreatedAt:   time.Now().In(location),
		}

		entity := converter.ToEntity(&order)
		err = repo.Create(context.Background(), entity)
		assert.Nil(t, err)

		//order.Status.Name = "approved"
		entity = converter.ToEntity(&order)
		err = repo.Update(context.Background(), entity)
		assert.Nil(t, err)
		assert.Equal(t, "approved", entity.Status.Name)

		canceledStatus := "canceled"
		order.Status = canceledStatus
		entity = converter.ToEntity(&order)
		err = repo.UpdateStatus(context.Background(), entity.ID, entity.Status.Name)
		assert.Nil(t, err)
		assert.Equal(t, "canceled", entity.Status.Name)

	})
}
