package repositorygorm

import (
	"context"
	"errors"
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	sharedDate "github.com/caiojorge/fiap-challenge-ddd/internal/shared/time"
	"gorm.io/gorm"
)

type ProductRepositoryGorm struct {
	DB        *gorm.DB
	converter converter.Converter[entity.Product, model.Product]
}

func NewProductRepositoryGorm(db *gorm.DB, converter converter.Converter[entity.Product, model.Product]) *ProductRepositoryGorm {
	return &ProductRepositoryGorm{
		DB:        db,
		converter: converter,
	}
}

// Create creates a new product. It returns an error if something goes wrong.
func (r *ProductRepositoryGorm) Create(ctx context.Context, entity *entity.Product) error {
	fmt.Println("repositorygorm: Criando produto: " + entity.GetID())
	model := r.converter.FromEntity(entity)

	model.CreatedAt = sharedDate.GetBRTimeNow()

	return r.DB.Create(model).Error
}

func (r *ProductRepositoryGorm) Update(ctx context.Context, entity *entity.Product) error {

	var productModel model.Product
	result := r.DB.Model(&model.Product{}).Where("id = ?", entity.ID).First(&productModel)
	if result.Error != nil {
		return result.Error
	}

	productModel.Category = entity.Category
	productModel.Description = entity.Description
	productModel.Name = entity.Name
	productModel.Price = entity.Price

	result = r.DB.Save(&productModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepositoryGorm) Find(ctx context.Context, id string) (*entity.Product, error) {
	var productModel model.Product
	result := r.DB.Model(&model.Product{}).Where("id = ?", id).First(&productModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := r.converter.ToEntity(&productModel)

	return entity, nil
}

func (r *ProductRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Product, error) {
	var mProducts []model.Product
	result := r.DB.Find(&mProducts)
	if result.Error != nil {
		return nil, result.Error
	}

	var eProducts []*entity.Product

	for _, mProduct := range mProducts {
		eProduct := r.converter.ToEntity(&mProduct)
		eProducts = append(eProducts, eProduct)
	}

	return eProducts, nil
}

func (r *ProductRepositoryGorm) Delete(ctx context.Context, id string) error {
	var productModel model.Product
	result := r.DB.Model(&model.Product{}).Where("id = ?", id).First(&productModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("repositorygorm: product not found")
			return nil
		}
		return result.Error
	}

	result = r.DB.Delete(&productModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepositoryGorm) FindByName(ctx context.Context, name string) (*entity.Product, error) {
	var productModel model.Product
	result := r.DB.Model(&model.Product{}).Where("name = ?", name).First(&productModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := r.converter.ToEntity(&productModel)

	return entity, nil
}

func (r *ProductRepositoryGorm) FindByCategory(ctx context.Context, category string) ([]*entity.Product, error) {
	var productModel []model.Product
	result := r.DB.Model(&model.Product{}).Where("category = ?", category).Find(&productModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	var entities []*entity.Product
	for _, product := range productModel {
		entity := r.converter.ToEntity(&product)
		entities = append(entities, entity)
	}

	return entities, nil
}
