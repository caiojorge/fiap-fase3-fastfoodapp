package repositorygorm

import (
	"context"
	"testing"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

// create a test function
func TestProdcut(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&model.Product{})
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	converter := converter.NewProductConverter()

	repo := NewProductRepositoryGorm(db, converter)

	id := sharedgenerator.NewIDGenerator()
	assert.NotEmpty(t, id)

	product, err := entity.ConvertProduct(id, "Lanche XPTO", "Pão queijo e carne", "Lanches", 30.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.GetID())

	err = repo.Create(ctx, product)

	assert.Nil(t, err)

	product2, err := repo.Find(ctx, product.GetID())
	assert.Nil(t, err)
	assert.NotNil(t, product2)
	assert.Equal(t, product.ID, product2.ID)

	products, err := repo.FindAll(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Len(t, products, 1)

	product.Name = "Lanche XPTO 2"
	err = repo.Update(ctx, product)
	assert.Nil(t, err)

	product2, err = repo.Find(ctx, product.GetID())
	assert.Nil(t, err)
	assert.NotNil(t, product2)
	assert.Equal(t, product, product2)

	product2, err = repo.FindByName(ctx, product.GetName())
	assert.Nil(t, err)
	assert.NotNil(t, product2)
	assert.Equal(t, product, product2)

	err = repo.Delete(ctx, product.GetID())
	assert.Nil(t, err)

	product2, err = repo.Find(ctx, product.GetID())
	assert.Nil(t, err)
	assert.Nil(t, product2)

}
