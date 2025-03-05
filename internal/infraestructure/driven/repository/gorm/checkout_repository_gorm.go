package repositorygorm

import (
	"context"
	"errors"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	sharedDate "github.com/caiojorge/fiap-challenge-ddd/internal/shared/time"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CheckoutRepositoryGorm struct {
	DB *gorm.DB
}

func (r *CheckoutRepositoryGorm) getDB(ctx context.Context) *gorm.DB {
	// Verifica se existe uma transação ativa no contexto
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return r.DB
}

func NewCheckoutRepositoryGorm(db *gorm.DB) *CheckoutRepositoryGorm {
	return &CheckoutRepositoryGorm{
		DB: db,
	}
}

// Create creates a new checkcout. It returns an error if something goes wrong.
func (r *CheckoutRepositoryGorm) Create(ctx context.Context, entity *entity.Checkout) error {

	model := model.Checkout{
		ID:                   entity.ID,
		OrderID:              entity.OrderID,
		GatewayName:          entity.Gateway.GatewayName,
		GatewayToken:         entity.Gateway.GatewayToken,
		GatewayTransactionID: entity.Gateway.GatewayTransactionID,
		Total:                entity.Total,
		CreatedAt:            sharedDate.GetBRTimeNow(),
		QRCode:               entity.QRCode,
	}

	db := r.getDB(ctx)
	if err := db.Create(&model).Error; err != nil {
		return err
	}

	return nil
}

// Update updates the checkout. It returns an error if something goes wrong.
func (r *CheckoutRepositoryGorm) Update(ctx context.Context, entity *entity.Checkout) error {

	//var model model.Checkout
	//copier.Copy(&model, entity)
	model := model.Checkout{
		ID:                   entity.ID,
		OrderID:              entity.OrderID,
		GatewayName:          entity.Gateway.GatewayName,
		GatewayToken:         entity.Gateway.GatewayToken,
		GatewayTransactionID: entity.Gateway.GatewayTransactionID,
		Total:                entity.Total,
		CreatedAt:            entity.CreatedAt,
		QRCode:               entity.QRCode,
	}
	db := r.getDB(ctx)
	return db.Save(model).Error
}

func (r *CheckoutRepositoryGorm) UpdateStatus(ctx context.Context, id string, status string) error {
	db := r.getDB(ctx)
	// Usando RAW SQL para atualizar o status da ordem
	result := db.Exec("UPDATE checkouts SET status = ? WHERE id = ?", status, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Find checkout by id
func (r *CheckoutRepositoryGorm) Find(ctx context.Context, id string) (*entity.Checkout, error) {
	var orderModel model.Checkout
	result := r.DB.Find(&orderModel, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	var entity *entity.Checkout
	err := copier.Copy(&entity, &orderModel)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// FindAll not implemented
func (r *CheckoutRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Checkout, error) {
	// var mOrders []model.Order

	// result := r.DB.Preload("Items").Order("created_at desc").Find(&mOrders)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// var eOrders []*entity.Order

	// for _, mOrder := range mOrders {
	// 	eOrder := r.converter.ToEntity(&mOrder)
	// 	eOrders = append(eOrders, eOrder)
	// }

	// return eOrders, nil
	return nil, nil
}

// Delete not implemented
func (r *CheckoutRepositoryGorm) Delete(ctx context.Context, id string) error {
	// var orderModel model.Order
	// result := r.DB.Model(&model.Order{}).Where("id = ?", id).First(&orderModel)
	// if result.Error != nil {
	// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 		fmt.Println("repositorygorm: order not found")
	// 		return nil
	// 	}
	// 	return result.Error
	// }

	// result = r.DB.Delete(&orderModel)
	// if result.Error != nil {
	// 	return result.Error
	// }

	return nil
}

func (r *CheckoutRepositoryGorm) FindbyOrderID(ctx context.Context, id string) (*entity.Checkout, error) {
	var orderModel model.Checkout
	result := r.DB.Find(&orderModel, "order_id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	if orderModel.ID == "" {
		return nil, errors.New("checkout not found")
	}

	entity := &entity.Checkout{
		ID:        orderModel.ID,
		OrderID:   orderModel.OrderID,
		Gateway:   valueobject.Gateway{GatewayName: orderModel.GatewayName, GatewayToken: orderModel.GatewayToken, GatewayTransactionID: orderModel.GatewayTransactionID},
		Total:     orderModel.Total,
		CreatedAt: orderModel.CreatedAt,
	}

	return entity, nil
}
