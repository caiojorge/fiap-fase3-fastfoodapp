package usecase

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	portsservice "github.com/caiojorge/fiap-challenge-ddd/internal/domain/gateway"
	repository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
	sharedurl "github.com/caiojorge/fiap-challenge-ddd/internal/shared/url"
	"github.com/joho/godotenv"
)

// CheckoutCreateUseCase é a implementação da interface CheckoutCreateUseCase.
// Nesse momento, não iremos implementar a integração com o gateway de pagamento.
type CheckoutCreateUseCase struct {
	orderRepository    repository.OrderRepository
	checkoutRepository repository.CheckoutRepository
	gatewayService     portsservice.GatewayTransactionService
	kitchenRepository  repository.KitchenRepository
	productRepository  repository.ProductRepository
}

func NewCheckoutCreate(orderRepository repository.OrderRepository,
	checkoutRepository repository.CheckoutRepository,
	gatewayService portsservice.GatewayTransactionService,
	kitchenRepository repository.KitchenRepository,
	productRepository repository.ProductRepository) *CheckoutCreateUseCase {
	return &CheckoutCreateUseCase{
		orderRepository:    orderRepository,
		checkoutRepository: checkoutRepository,
		gatewayService:     gatewayService,
		productRepository:  productRepository,
		kitchenRepository:  kitchenRepository,
	}
}

// CreateCheckout registra o checkout de um pedido.
// Requisito: Checkout Pedido que deverá receber os produtos solicitados e retornar à identificação do pedido.
// o checkout recebe alguns parametros, sendo um deles, o id da ordem. Esse id, é usado para retornar a ordem com seus items, que por sua vez, tem os produtos.
// É dessa forma que penso em atender o requisito indicado acima
func (cr *CheckoutCreateUseCase) CreateCheckout(ctx context.Context, checkoutDTO *CheckoutInputDTO) (*CheckoutOutputDTO, error) {

	log.Default().Println("CreateCheckout")

	if checkoutDTO.NotificationURL == "" {
		_ = godotenv.Load() // Carrega o .env se não estiver definido em variáveis de ambiente

		hostname := os.Getenv("HOST_NAME")
		hostport := os.Getenv("HOST_PORT_K8S")

		u := sharedurl.NewURL(hostname, hostport)
		url := u.GetWebhookURL()
		checkoutDTO.NotificationURL = url
	}

	log.Default().Println("validou url")

	err := cr.validateDuplicatedCheckout(ctx, checkoutDTO)
	if err != nil {
		log.Default().Println("erro:" + err.Error())
		return nil, err
	}
	log.Default().Println("validou checkout duplicado")

	// Order- o pedido deve existir e não pode estar pago
	order, err := cr.validateAndReturnOrder(ctx, checkoutDTO)
	if err != nil {
		log.Default().Println("erro:" + err.Error())
		return nil, err
	}
	log.Default().Println("validou ordem")

	// via orderitems podemos pegar os produtos do pedido
	productList, err := cr.getProductList(ctx, order)
	if err != nil {
		log.Default().Println("erro:" + err.Error())
		return nil, errors.New("no products to send to checkout")
	}
	log.Default().Println("buscou produtos")

	// converte dto para entidade
	checkout := checkoutDTO.ToEntity()
	log.Default().Println("to entity...")

	// integração com o gateway de pagamento e confirmação do pagamento na ordem
	payment, err := cr.handlePayment(ctx, checkout, order, productList, checkoutDTO.NotificationURL, checkoutDTO.SponsorID, checkoutDTO.DiscontCoupon)
	if err != nil {
		log.Default().Println("erro:" + err.Error())
		return nil, err
	}
	log.Default().Println("integração com gateway")

	// confirma a transação e grava o checkout
	err = cr.handleCheckout(ctx, checkout, order.Total, payment.InStoreOrderID, payment.QrData)
	if err != nil {
		log.Default().Println("erro:" + err.Error())
		return nil, err
	}
	log.Default().Println("grava o checkout")

	// muda o status para checkout-confirmado
	err = cr.handleOrder(ctx, order)
	if err != nil {
		log.Default().Println("erro:" + err.Error())
		return nil, err
	}
	log.Default().Println("muda o status da ordem")

	output := &CheckoutOutputDTO{
		ID:                   checkout.ID,
		GatewayTransactionID: payment.InStoreOrderID,
		OrderID:              order.ID,
	}

	return output, nil
}

// Checkout - o cliente não pode fazer checkout duas vezes
// se existir um checkout para o pedido, não pode fazer outro
func (cr *CheckoutCreateUseCase) validateDuplicatedCheckout(ctx context.Context, checkout *CheckoutInputDTO) error {
	ch, _ := cr.checkoutRepository.FindbyOrderID(ctx, checkout.OrderID)
	if ch != nil {
		return errors.New("you can not checkout twice")
	}

	return nil
}

// Order- o pedido deve existir e não pode estar pago
func (cr *CheckoutCreateUseCase) validateAndReturnOrder(ctx context.Context, checkout *CheckoutInputDTO) (*entity.Order, error) {
	// Order- o pedido deve existir e não pode estar pago
	order, err := cr.orderRepository.Find(ctx, checkout.OrderID)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}
	if order.IsPaymentApproved() {
		return nil, errors.New("order already paid")
	}

	return order, nil
}

// getProductList retorna a lista de produtos do pedido
// TODO: esse método pode ser movido para um repositório de produtos
func (cr *CheckoutCreateUseCase) getProductList(ctx context.Context, order *entity.Order) ([]*entity.Product, error) {
	var products []*entity.Product

	for _, item := range order.Items {
		product, err := cr.productRepository.Find(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// handlePayment integração com o gateway de pagamento e confirmação do pagamento na ordem
func (cr *CheckoutCreateUseCase) handlePayment(ctx context.Context, checkout *entity.Checkout, order *entity.Order, productList []*entity.Product, notificationURL string, sponsorID int, cupon float64) (*entity.Payment, error) {

	payment, err := cr.gatewayService.ConfirmPayment(ctx, checkout, order, productList, notificationURL, sponsorID)
	if err != nil {
		order.InformPaymentNotApproval()
		_ = cr.orderRepository.Update(ctx, order)

		//checkout.InformPaymentNotApproval()
		return nil, err
	}

	if payment == nil {
		order.InformPaymentNotApproval()
		_ = cr.orderRepository.Update(ctx, order)
		return nil, errors.New("failed to create transaction on gateway")
	}

	return payment, nil
}

// handleCheckout confirma a transação e grava o checkout
func (cr *CheckoutCreateUseCase) handleCheckout(ctx context.Context, checkout *entity.Checkout, orderTotal float64, paymentID string, qrcode string) error {

	// confirma o pagamento de forma fake na própria entidade (muda status na entidade)
	// se houver algum cupom de desconto, ele será aplicado na ordem no handlePayment
	err := checkout.ConfirmTransaction(paymentID, orderTotal, qrcode)
	if err != nil {
		cr.gatewayService.CancelPayment(ctx, paymentID)
		return err
	}

	// grava o checkout
	err = cr.checkoutRepository.Create(ctx, checkout)
	if err != nil {
		cr.gatewayService.CancelPayment(ctx, paymentID)
		return err
	}

	return nil
}

func (cr *CheckoutCreateUseCase) handleOrder(ctx context.Context, order *entity.Order) error {
	order.ConfirmCheckout()
	//err := cr.orderRepository.Update(ctx, order)
	err := cr.orderRepository.UpdateStatus(ctx, order.ID, order.Status.Name)
	if err != nil {
		return err
	}
	return nil
}
