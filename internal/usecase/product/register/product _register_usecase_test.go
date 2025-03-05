package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	//mock "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/product/mock"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
)

func TestProductRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocksrepository.NewMockProductRepository(ctrl)

	mockProductRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Product{})).
		Return(nil)

	mockProductRepo.EXPECT().
		FindByName(gomock.Any(), "Lanche XPTO").
		Return(nil, nil)

	productDto := RegisterProductInputDTO{
		Name:        "Lanche XPTO",
		Description: "Pão queijo e carne",
		Category:    "Lanches",
		Price:       30.00,
	}

	assert.Equal(t, "Lanche XPTO", productDto.Name)

	register := NewProductRegister(mockProductRepo)
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
		FindAll(gomock.Any()).
		Return([]*entity.Product{product}, nil)

	mockProductRepo.EXPECT().
		Find(gomock.Any(), output.ID).
		Return(product, nil)

	product2, err := mockProductRepo.Find(context.Background(), output.ID)
	assert.Nil(t, err)
	assert.NotNil(t, product2)
	assert.Equal(t, output.ID, product2.ID)

	products, err := mockProductRepo.FindAll(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Len(t, products, 1)

}
