package repositorygorm

import (
	"context"
	"errors"
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	sharedtime "github.com/caiojorge/fiap-challenge-ddd/internal/shared/time"
	"gorm.io/gorm"
)

type OrderRepositoryGorm struct {
	DB        *gorm.DB
	converter converter.Converter[entity.Order, model.Order]
}

func (r *OrderRepositoryGorm) getDB(ctx context.Context) *gorm.DB {
	// Verifica se existe uma transação ativa no contexto
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return r.DB
}

func NewOrderRepositoryGorm(db *gorm.DB, converter converter.Converter[entity.Order, model.Order]) *OrderRepositoryGorm {
	return &OrderRepositoryGorm{
		DB:        db,
		converter: converter,
	}
}

// Create creates a new product. It returns an error if something goes wrong.
func (r *OrderRepositoryGorm) Create(ctx context.Context, entity *entity.Order) error {

	db := r.getDB(ctx)

	model := parserOrder(entity)

	model.CreatedAt = sharedtime.GetBRTimeNow()

	if model.CustomerCPF != nil && *model.CustomerCPF == "" {
		model.CustomerCPF = nil
	}

	err := db.Create(model).Error
	if err != nil {
		fmt.Println("gorm: " + err.Error())
		return err
	}

	return nil
}

func parserOrder(entity *entity.Order) model.Order {
	modelItems := make([]*model.OrderItem, len(entity.Items))
	for i, item := range entity.Items {
		modelItems[i] = &model.OrderItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Status:    item.Status,
			OrderID:   entity.ID,
		}
	}

	model := model.Order{
		ID:             entity.ID,
		Items:          modelItems,
		Total:          entity.Total,
		CustomerCPF:    &entity.CustomerCPF,
		Status:         entity.Status.Name,
		DeliveryNumber: entity.DeliveryNumber,
	}

	if *model.CustomerCPF == "" {
		model.CustomerCPF = nil
	}
	return model
}

func (r *OrderRepositoryGorm) Update(ctx context.Context, entity *entity.Order) error {

	db := r.getDB(ctx)

	var retrievedOrder model.Order
	//err := db.Preload("Status").First(&retrievedOrder, "id = ?", entity.ID).Error
	err := db.First(&retrievedOrder, "id = ?", entity.ID).Error
	if err != nil {
		return err
	}

	model := parserOrder(entity)
	if *model.CustomerCPF == "" {
		model.CustomerCPF = nil
	}

	result := db.Save(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepositoryGorm) UpdateStatus(ctx context.Context, id string, status string) error {
	db := r.getDB(ctx)

	// Usando RAW SQL para atualizar o status da ordem
	result := db.Exec("UPDATE orders SET status = ? WHERE id = ?", status, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Find retrieves a product by its ID. It returns an error if something goes wrong.
func (r *OrderRepositoryGorm) Find(ctx context.Context, id string) (*entity.Order, error) {
	var orderModel model.Order
	result := r.DB.Preload("Items").Order("created_at desc").Find(&orderModel, "id = ?", id)
	//result := r.DB.Model(&model.Order{}).Where("id = ?", id).First(&orderModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	entity := r.converter.ToEntity(&orderModel)

	return entity, nil
}

func (r *OrderRepositoryGorm) FindAll(ctx context.Context) ([]*entity.Order, error) {
	var mOrders []model.Order

	result := r.DB.Preload("Items").Order("created_at desc").Find(&mOrders)
	if result.Error != nil {
		return nil, result.Error
	}

	var eOrders []*entity.Order

	for _, mOrder := range mOrders {
		eOrder := r.converter.ToEntity(&mOrder)
		eOrders = append(eOrders, eOrder)
	}

	return eOrders, nil
}

func (r *OrderRepositoryGorm) Delete(ctx context.Context, id string) error {
	var orderModel model.Order

	db := r.getDB(ctx)
	result := db.Model(&model.Order{}).Where("id = ?", id).First(&orderModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("repositorygorm: order not found")
			return nil
		}
		return result.Error
	}

	result = r.DB.Delete(&orderModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepositoryGorm) FindByParams(ctx context.Context, params map[string]interface{}) ([]*entity.Order, error) {

	var orders []*entity.Order
	var models []*model.Order

	query := r.DB.Preload("Items").Order("created_at desc")
	//query := r.DB.Model(&model.Order{})

	// Adiciona condições dinâmicas com base nos parâmetros
	if status, ok := params["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if customerCPF, ok := params["customer_cpf"]; ok {
		query = query.Where("customer_cpf = ?", customerCPF)
	}
	if startDate, ok := params["start_date"]; ok {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate, ok := params["end_date"]; ok {
		query = query.Where("created_at <= ?", endDate)
	}

	// Executa a consulta
	err := query.Find(&models).Error
	if err != nil {
		return nil, err
	}

	//copier.Copy(&orders, &models)

	for _, model := range models {
		order := r.converter.ToEntity(model)
		orders = append(orders, order)
	}

	return orders, err

}
