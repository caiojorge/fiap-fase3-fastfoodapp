package usecase

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	sharedgenerator "github.com/caiojorge/fiap-challenge-ddd/internal/shared/generator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
)

func TestProductRegisterAndUpdater(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocksrepository.NewMockProductRepository(ctrl)

	mockProductRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Product{})).
		Return(nil)

	product := &entity.Product{
		ID:          sharedgenerator.NewIDGenerator(),
		Name:        "Lanche XPTO 2",
		Description: "Pão queijo e carne",
		Category:    "Lanches",
		Price:       30.00,
	}

	mockProductRepo.Create(context.Background(), product)

	mockProductRepo.EXPECT().
		FindAll(gomock.Any()).
		Return([]*entity.Product{product}, nil)

	finderAll := NewProductFindAll(mockProductRepo)
	assert.NotNil(t, finderAll)

	outputs, err := finderAll.FindAllProducts(context.Background())
	assert.NotNil(t, outputs)
	assert.Len(t, outputs, 1)
	assert.Nil(t, err)

}
