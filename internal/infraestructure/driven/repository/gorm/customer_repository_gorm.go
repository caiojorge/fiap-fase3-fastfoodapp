package repositorygorm

import (
	"context"
	"errors"
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
	sharedDate "github.com/caiojorge/fiap-challenge-ddd/internal/shared/time"
	"gorm.io/gorm"
)

type CustomerRepositoryGorm struct {
	DB *gorm.DB
}

func NewCustomerRepositoryGorm(db *gorm.DB) *CustomerRepositoryGorm {
	return &CustomerRepositoryGorm{
		DB: db,
	}
}

func (r *CustomerRepositoryGorm) Create(ctx context.Context, entity *entity.Customer) error {
	fmt.Println("repositorygorm: Criando cliente: " + entity.GetCPF().Value)
	model := converter.FromEntity(entity)
	return r.DB.Create(model).Error
}

func (r *CustomerRepositoryGorm) Update(ctx context.Context, entity *entity.Customer) error {

	model := model.Customer{
		CPF:       formatter.RemoveMaskFromCPF(entity.GetCPF().Value),
		Name:      entity.GetName(),
		Email:     entity.GetEmail(),
		CreatedAt: sharedDate.GetBRTimeNow(),
	}

	if model.CPF == "" {
		return fmt.Errorf("cpf is required")
	}

	result := r.DB.Save(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Find busca um cliente pelo CPF (sem máscara!)
func (r *CustomerRepositoryGorm) Find(ctx context.Context, id string) (*entity.Customer, error) {
	var customerModel model.Customer
	// sempre removo a mascara do cpf para buscar no banco
	result := r.DB.Model(&model.Customer{}).Where("cpf = ?", formatter.RemoveMaskFromCPF(id)).First(&customerModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := converter.ToEntity(&customerModel)

	return entity, nil
}

func (r *CustomerRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Customer, error) {
	var mCustomers []model.Customer
	result := r.DB.Find(&mCustomers)
	if result.Error != nil {
		return nil, result.Error
	}

	var eCustomers []*entity.Customer

	for _, mCustomer := range mCustomers {
		eCustomer := converter.ToEntity(&mCustomer)
		eCustomers = append(eCustomers, eCustomer)
	}

	return eCustomers, nil
}
