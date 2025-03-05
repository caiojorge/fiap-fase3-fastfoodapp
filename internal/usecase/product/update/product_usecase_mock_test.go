package usecase

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	registerusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/register"
)

func TestProductRegisterAndUpdater(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocksrepository.NewMockProductRepository(ctrl)

	mockProductRepo.EXPECT().
		FindByName(gomock.Any(), "Lanche XPTO").
		Return(nil, nil) // Simula que o produto não existe ainda, permitindo o registro

	mockProductRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Product{})).
		Return(nil)

	productDto := registerusecase.RegisterProductInputDTO{
		Name:        "Lanche XPTO",
		Description: "Pão queijo e carne",
		Category:    "Lanches",
		Price:       30.00,
	}

	register := registerusecase.NewProductRegister(mockProductRepo)
	assert.NotNil(t, register)

	output, err := register.RegisterProduct(context.Background(), &productDto)
	assert.Nil(t, err)

	product := &entity.Product{
		ID:          output.ID,
		Name:        "Lanche XPTO",
		Description: "Pão, carne, queijo e presunto",
		Category:    "Lanches",
		Price:       100,
	}

	mockProductRepo.EXPECT().
		Find(gomock.Any(), output.ID).
		Return(product, nil).
		Times(2)

	product2, err := mockProductRepo.Find(context.Background(), output.ID)
	assert.Nil(t, err)
	assert.NotNil(t, product2)
	assert.Equal(t, output.ID, product2.ID)

	updater := NewProductUpdate(mockProductRepo)
	assert.NotNil(t, updater)

	input := UpdateProductInputDTO{
		ID:          output.ID,
		Name:        "Lanche XPTO 2",
		Description: "Pão queijo e carne",
		Category:    "Lanches",
		Price:       30.00,
	}

	mockProductRepo.EXPECT().
		Update(gomock.Any(), gomock.AssignableToTypeOf(&entity.Product{})).
		Return(nil)

	outputUpdate, err := updater.UpdateProduct(context.Background(), input)
	assert.Nil(t, err)
	assert.NotNil(t, outputUpdate)
	assert.NotNil(t, outputUpdate.Description)

}
