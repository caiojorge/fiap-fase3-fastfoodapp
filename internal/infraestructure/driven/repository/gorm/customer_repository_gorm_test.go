package repositorygorm

import (
	"context"
	"testing"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

// create a test function
func TestNewCustomer(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&model.Customer{})
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	repo := NewCustomerRepositoryGorm(db)

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := entity.NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	err = repo.Create(ctx, customer)
	assert.Nil(t, err)

	customers, err := repo.FindAll(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, customers)
	assert.Len(t, customers, 1)

	customer2, err := repo.Find(ctx, "12345678909")
	assert.Nil(t, err)
	assert.NotNil(t, customer2)
	assert.Equal(t, customer, customer2)

	customer.Name = "John Doe 2"
	err = repo.Update(ctx, customer)
	assert.Nil(t, err)

	customer3, err := repo.Find(ctx, "12345678909")
	assert.Nil(t, err)
	assert.NotNil(t, customer3)
	assert.Equal(t, customer.Name, customer3.GetName())
}
