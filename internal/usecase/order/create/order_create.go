package usecase

import (
	"context"
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	domainRepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"

	//portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"

	"github.com/jinzhu/copier"
)

// OrderCreateUseCase é a implementação de OrderCreateUseCase.
// Um caso de uso é uma estrutura que contém todas as regras de negócio para uma determinada funcionalidade.
// Nesse cenário, vamos precisar acessar 3 agregados e seus repositorios.
type OrderCreateUseCase struct {
	orderRepository    domainRepository.OrderRepository
	customerRepository domainRepository.CustomerRepository
	productRepository  domainRepository.ProductRepository
}

func NewOrderCreate(orderRepository domainRepository.OrderRepository,
	customerRepository domainRepository.CustomerRepository,
	productRepository domainRepository.ProductRepository) *OrderCreateUseCase {
	return &OrderCreateUseCase{
		orderRepository:    orderRepository,
		customerRepository: customerRepository,
		productRepository:  productRepository,
	}
}

// CreateOrder registra um novo pedido.
func (cr *OrderCreateUseCase) CreateOrder(ctx context.Context, input *OrderCreateInputDTO) (*OrderCreateOutputDTO, error) {

	if input.Items == nil || len(input.Items) == 0 {
		return nil, fmt.Errorf("order items not informed")
	}

	for _, item := range input.Items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product %s", item.ProductID)
		}
	}

	wrapper := &CreateOrderWrapper{
		dto: input,
	}

	// converte o DTO de input para a entidade Order
	order, err := wrapper.ToEntity()
	if err != nil {
		return nil, err
	}

	// handle customer
	_, err = cr.handleCustomer(ctx, order)
	if err != nil {
		return nil, err
	}

	fmt.Println("usecase: Criando Order: " + order.CustomerCPF)

	products, err := cr.getProductList(ctx, order)
	if err != nil {
		return nil, err
	}

	order.ConfirmItemsPrice(products)

	fmt.Println("usecase: Criando Order: ConfirmOrder: " + order.CustomerCPF)

	order.Confirm()

	if err := cr.orderRepository.Create(ctx, order); err != nil {
		fmt.Println("usecase: repo create: " + err.Error())
		return nil, err
	}

	// copia os dados da ordem para o output
	var output OrderCreateOutputDTO
	err = copier.Copy(&output, &order)
	if err != nil {
		return nil, err
	}

	output.Status = order.Status.Name

	return &output, nil
}

func (cr *OrderCreateUseCase) handleCustomer(ctx context.Context, order *entity.Order) (*entity.Customer, error) {
	// handle customer
	// se o cpf for empty, indica que o cliente não quis se identificar, e isso esta ok, segundo as regras de negócio
	if !order.IsCustomerInformed() {
		return nil, nil
	}
	// busca o cliente pelo cpf
	customer, err := cr.customerRepository.Find(ctx, order.CustomerCPF)
	if err != nil {
		return nil, err
	}

	// se o cliente for informado, temos q validar o cpf, e cadastra-lo caso não exista
	// só entra aqui se o cliente não existir na base de dados
	// se o cliente não for nulo, ele já existe o cpf é considerado válido
	// em teoria, não existe ordem duplicada. o mesmo cliente pode comprar várias vezes. (não vou validar isso aqui)
	if customer == nil {
		// o cliente não é obrigatório, mas se for informado, ele precisa ser válido.
		// apenas nesse caso, se o cliente não existir, ele será persistido. (apenas o cpf)
		// identifica o cliente pelo cpf
		newCustomer, err := entity.NewCustomerWithCPFInformed(order.CustomerCPF)
		if err != nil {
			return nil, err
		}

		// cria o cliente sem o nome e email
		if err := cr.customerRepository.Create(ctx, newCustomer); err != nil {
			return nil, err
		}
	}

	return customer, nil

}

// getProductList retorna a lista de produtos do pedido
// TODO: esse método pode ser movido para um repositório de produtos
func (cr *OrderCreateUseCase) getProductList(ctx context.Context, order *entity.Order) ([]*entity.Product, error) {
	var products []*entity.Product

	for _, item := range order.Items {
		product, err := cr.productRepository.Find(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}

		if product == nil {
			return nil, fmt.Errorf("product not found: %s", item.ProductID)
		}

		products = append(products, product)
	}

	return products, nil
}
