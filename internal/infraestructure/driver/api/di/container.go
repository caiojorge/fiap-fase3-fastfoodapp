package di

import (
	"context"
	"errors"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	service "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/gateway"

	repositorygorm "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/repository/gorm"
	controllercheckout "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/checkout"
	controllercustomer "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/customer"
	controllerkitchen "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/kitchen"
	controllerorder "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/order"
	controllerproduct "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/product"

	usecasecheckoutcheck "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/checkpayment"
	usecasewebhookcheckout "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/confirmation"
	usecasecheckout "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/create"
	usecaseReprocessing "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/reprocessing"

	usecasecustomerfindall "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/findall"
	usecasecustomerfindbycpf "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/findbycpf"
	usecasecustomerregister "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/register"
	usecasecustomerupdate "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/update"

	usecasekitchenCook "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/cooking"
	usecasekitchenDelivery "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/delivery"
	usecasekitchen "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/findall"
	usecasekitchenMonitor "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/monitor"
	usecasekitchenNotifier "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/notifier"

	usecaseordercreate "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/create"
	usecaseorderfindall "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findall"
	usecaseorderfindbyid "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findbyid"
	usecaseorderfindbyparam "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findbyparam"

	usecaseproductdelete "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/delete"
	usecaseproductfindall "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findall"
	usecaseproductfindbycategory "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findbycategory"
	usecaseproductfindbyid "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findbyid"
	usecaseproductregister "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/register"
	usecaseproductupdater "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/update"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	DB     *gorm.DB
	Logger *zap.Logger

	/* repositories */
	// ProductRepo      repositorygorm.ProductRepositoryGorm
	// OrderRepo        repositorygorm.OrderRepositoryGorm
	// CustomerRepo     repositorygorm.CustomerRepositoryGorm
	// CheckoutRepo     repositorygorm.CheckoutRepositoryGorm
	// KitchenRepo      repositorygorm.KitchenRepositoryGorm
	// GatewayService   service.PaymentGateway
	// ProductConverter converter.ProductConverter
	// OrderConverter   converter.OrderConverter

	// Customer
	RegisterCustomerController  controllercustomer.RegisterCustomerController
	UpdateCustomerController    controllercustomer.UpdateCustomerController
	FindCustomerByCPFController controllercustomer.FindCustomerByCPFController
	FindAllCustomersController  controllercustomer.FindAllCustomersController
	// Customer
	// CustomerRegisterUseCase  usecasecustomerregister.CustomerRegisterUseCase
	// CustomerUpdateUseCase    usecasecustomerupdate.CustomerUpdateUseCase
	// CustomerFindByCPFUseCase usecasecustomerfindbycpf.CustomerFindByCPFUseCase
	// CustomerFindAllUseCase   usecasecustomerfindall.CustomerFindAllUseCase

	// Product
	RegisterProductController       controllerproduct.RegisterProductController
	UpdateProductController         controllerproduct.UpdateProductController
	FindProductByIDController       controllerproduct.FindProductByIDController
	FindAllProductController        controllerproduct.FindAllProductController
	FindProductByCategoryController controllerproduct.FindProductByCategoryController
	DeleteProductController         controllerproduct.DeleteProductController
	// Product
	// ProductRegisterUseCase       usecaseproductregister.ProductRegisterUseCase
	// ProductUpdateUseCase         usecaseproductupdater.ProductUpdateUseCase
	// ProductFindByIDUseCase       usecaseproductfindbyid.ProductFindByIDUseCase
	// ProductFindAllUseCase        usecaseproductfindall.ProductFindAllUseCase
	// ProductFindByCategoryUseCase usecaseproductfindbycategory.ProductFindByCategoryUseCase
	// ProductDeleteUseCase         usecaseproductdelete.ProductDeleteUseCase

	// Order
	CreateOrderController                       controllerorder.CreateOrderController
	FindAllOrdersController                     controllerorder.FindAllController
	FindOrderByIDController                     controllerorder.FindOrderByIDController
	FindByParamsOrdersNotConfirmedController    controllerorder.FindByParamsNotConfirmedController
	FindByParamsOrdersConfirmedController       controllerorder.FindByParamsConfirmedController
	FindByParamsOrdersPaymentApprovedController controllerorder.FindByParamsPaymentApprovedController
	// Order
	// OrderCreateUseCase       usecaseordercreate.OrderCreateUseCase
	// OrderFindAllUseCase      usecaseorderfindall.OrderFindAllUseCase
	// OrderFindByIDUseCase     usecaseorderfindbyid.OrderFindByIDUseCase
	// OrderFindByParamsUseCase usecaseorderfindbyparam.OrderFindByParamsUseCase

	// Checkout
	CreateCheckoutController       controllercheckout.CreateCheckoutController
	CheckoutCheckController        controllercheckout.CheckPaymentCheckoutController
	WebhookCheckoutController      controllercheckout.WebhookCheckoutController
	CheckoutReprocessingController controllercheckout.ReprocessingPaymentController
	// Checkout
	// CheckoutCreateUseCase    usecasecheckout.CheckoutCreateUseCase
	// CheckoutCheckUseCase     usecasecheckoutcheck.CheckPaymentUseCase
	// ReprocessingCheckUseCase usecaseReprocessing.ICheckoutReprocessingUseCase

	// Kitchen
	FindKitchenAllController  controllerkitchen.FindKitchenAllController
	NotifierKitchenController controllerkitchen.NotifyKitchenController
	CookingKitchenController  controllerkitchen.CookingKitchenController
	MonitorKitchenController  controllerkitchen.MonitorKitchenController
	DeliveryKitchenController controllerkitchen.DeliveryKitchenController
	// Kitchen
	// KitchenFindAllUseCase  usecasekitchen.KitchenFindAllUseCase
	NotifierKitchenUseCase usecasekitchenNotifier.KitchenNotifierUseCase
}

func NewContainer(db *gorm.DB, logger *zap.Logger) *Container {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Initialize converters
	productConverter := converter.NewProductConverter()
	orderConverter := converter.NewOrderConverter()

	// Initialize repositories
	customerRepo := repositorygorm.NewCustomerRepositoryGorm(db)
	productRepo := repositorygorm.NewProductRepositoryGorm(db, productConverter)
	orderRepo := repositorygorm.NewOrderRepositoryGorm(db, orderConverter)
	checkoutRepo := repositorygorm.NewCheckoutRepositoryGorm(db)
	kitchenRepo := repositorygorm.NewKitchenRepositoryGorm(db)
	gatewayService := service.NewPaymentGateway(service.NewPaymentService())
	transactionManager := repositorygorm.NewGormTransactionManager(db)

	// Customer
	customerRegisterUseCase := usecasecustomerregister.NewCustomerRegister(customerRepo)
	customerUpdateUseCase := usecasecustomerupdate.NewCustomerUpdate(customerRepo)
	customerFindByCPFUseCase := usecasecustomerfindbycpf.NewCustomerFindByCPF(customerRepo)
	customerFindAllUseCase := usecasecustomerfindall.NewCustomerFindAll(customerRepo)
	registerCustomerController := controllercustomer.NewRegisterCustomerController(ctx, customerRegisterUseCase)
	updateCustomerController := controllercustomer.NewUpdateCustomerController(ctx, customerUpdateUseCase)
	findCustomerByCPFController := controllercustomer.NewFindCustomerByCPFController(ctx, customerFindByCPFUseCase)
	findAllCustomersController := controllercustomer.NewFindAllCustomersController(ctx, customerFindAllUseCase)

	// Product
	productRegisterUseCase := usecaseproductregister.NewProductRegister(productRepo)
	productUpdateUseCase := usecaseproductupdater.NewProductUpdate(productRepo)
	productFindByIDUseCase := usecaseproductfindbyid.NewProductFindByID(productRepo)
	productFindAllUseCase := usecaseproductfindall.NewProductFindAll(productRepo)
	productFindByCategoryUseCase := usecaseproductfindbycategory.NewProductFindByCategory(productRepo)
	productDeleteUseCase := usecaseproductdelete.NewProductDelete(productRepo)
	registerProductController := controllerproduct.NewRegisterProductController(ctx, productRegisterUseCase)
	updateProductController := controllerproduct.NewUpdateProductController(ctx, productUpdateUseCase)
	findProductByIDController := controllerproduct.NewFindProductByIDController(ctx, productFindByIDUseCase)
	findAllProductController := controllerproduct.NewFindAllProductController(ctx, productFindAllUseCase)
	findProductByCategoryController := controllerproduct.NewFindProductByCategoryController(ctx, productFindByCategoryUseCase)
	deleteProductController := controllerproduct.NewDeleteProductController(ctx, productDeleteUseCase)

	// Order
	orderCreateUseCase := usecaseordercreate.NewOrderCreate(orderRepo, customerRepo, productRepo)
	orderFindAllUseCase := usecaseorderfindall.NewOrderFindAll(orderRepo)
	orderFindByIDUseCase := usecaseorderfindbyid.NewOrderFindByID(orderRepo)
	orderFindByParamsUseCase := usecaseorderfindbyparam.NewOrderFindByParams(orderRepo)
	createOrderController := controllerorder.NewCreateOrderController(ctx, orderCreateUseCase)
	findAllOrdersController := controllerorder.NewFindAllController(ctx, orderFindAllUseCase)
	findOrderByIDController := controllerorder.NewFindOrderByIDController(ctx, orderFindByIDUseCase)
	findByParamsOrdersNotConfirmedController := controllerorder.NewFindByParamsNotConfirmedController(ctx, orderFindByParamsUseCase)
	findByParamsOrdersConfirmedController := controllerorder.NewFindByParamsConfirmedController(ctx, orderFindByParamsUseCase)
	findByParamsOrdersPaymentApprovedController := controllerorder.NewFindByParamsPaymentApprovedController(ctx, orderFindByParamsUseCase)

	// Checkout
	checkoutCreateUseCase := usecasecheckout.NewCheckoutCreate(orderRepo, checkoutRepo, gatewayService, kitchenRepo, productRepo)
	checkoutCheckUseCase := usecasecheckoutcheck.NewCheckPaymentUseCase(checkoutRepo, orderRepo)
	reprocessingPaymentUseCase := usecaseReprocessing.NewCheckoutReprocessing(orderRepo, checkoutRepo, productRepo, gatewayService, transactionManager, logger)
	webhookCheckoutUseCase := usecasewebhookcheckout.NewCheckoutConfirmation(orderRepo, checkoutRepo, transactionManager, logger)
	createCheckoutController := controllercheckout.NewCreateCheckoutController(ctx, checkoutCreateUseCase)
	checkCheckoutController := controllercheckout.NewCheckPaymentCheckoutController(ctx, checkoutCheckUseCase)
	webhookCheckoutController := controllercheckout.NewWebhookCheckoutController(ctx, webhookCheckoutUseCase, logger)
	reprocessingCheckoutController := controllercheckout.NewReprocessingPaymentController(ctx, reprocessingPaymentUseCase)

	// Kitchen
	kitchenFindAllUseCase := usecasekitchen.NewKitchenFindAll(kitchenRepo)
	notifierkitchenUseCase := usecasekitchenNotifier.NewKitchenNotifierUseCase(kitchenRepo, orderRepo, productRepo)
	cookkitchenUseCase := usecasekitchenCook.NewKitchenCookingUseCase(kitchenRepo, orderRepo, productRepo)
	monitorkitchenUseCase := usecasekitchenMonitor.NewMonitorKitchenUseCase(kitchenRepo, orderRepo)
	deliverykitchenUseCase := usecasekitchenDelivery.NewKitchenDeliveryUseCase(kitchenRepo, orderRepo, productRepo)
	findKitchenAllController := controllerkitchen.NewFindKitchenAllController(ctx, kitchenFindAllUseCase)
	notifierKitchenController := controllerkitchen.NewNotifyKitchenController(ctx, notifierkitchenUseCase)
	cookKitchenController := controllerkitchen.NewCookingKitchenController(ctx, cookkitchenUseCase)
	monitorKitchenController := controllerkitchen.NewMonitorKitchenController(ctx, monitorkitchenUseCase)
	deliveryKitchenController := controllerkitchen.NewDeliveryKitchenController(ctx, deliverykitchenUseCase)

	return &Container{
		DB:     db,
		Logger: logger,

		// ProductRepo:      *productRepo,
		// OrderRepo:        *orderRepo,
		// CustomerRepo:     *customerRepo,
		// CheckoutRepo:     *checkoutRepo,
		// KitchenRepo:      *kitchenRepo,
		// GatewayService:   *gatewayService,
		// ProductConverter: *productConverter,
		// OrderConverter:   *orderConverter,

		// CustomerRegisterUseCase:     *customerRegisterUseCase,
		// CustomerUpdateUseCase:       *customerUpdateUseCase,
		// CustomerFindByCPFUseCase:    *customerFindByCPFUseCase,
		// CustomerFindAllUseCase:      *customerFindAllUseCase,
		RegisterCustomerController:  *registerCustomerController,
		UpdateCustomerController:    *updateCustomerController,
		FindCustomerByCPFController: *findCustomerByCPFController,
		FindAllCustomersController:  *findAllCustomersController,

		// ProductRegisterUseCase:          *productRegisterUseCase,
		// ProductUpdateUseCase:            *productUpdateUseCase,
		// ProductFindByIDUseCase:          *productFindByIDUseCase,
		// ProductFindAllUseCase:           *productFindAllUseCase,
		// ProductFindByCategoryUseCase:    *productFindByCategoryUseCase,
		// ProductDeleteUseCase:            *productDeleteUseCase,
		RegisterProductController:       *registerProductController,
		UpdateProductController:         *updateProductController,
		FindProductByIDController:       *findProductByIDController,
		FindAllProductController:        *findAllProductController,
		FindProductByCategoryController: *findProductByCategoryController,
		DeleteProductController:         *deleteProductController,

		// OrderCreateUseCase:                          *orderCreateUseCase,
		// OrderFindAllUseCase:                         *orderFindAllUseCase,
		// OrderFindByIDUseCase:                        *orderFindByIDUseCase,
		// OrderFindByParamsUseCase:                    *orderFindByParamsUseCase,
		CreateOrderController:                       *createOrderController,
		FindAllOrdersController:                     *findAllOrdersController,
		FindOrderByIDController:                     *findOrderByIDController,
		FindByParamsOrdersNotConfirmedController:    *findByParamsOrdersNotConfirmedController,
		FindByParamsOrdersConfirmedController:       *findByParamsOrdersConfirmedController,
		FindByParamsOrdersPaymentApprovedController: *findByParamsOrdersPaymentApprovedController,

		//CheckoutCreateUseCase:          *checkoutCreateUseCase,
		CreateCheckoutController:       *createCheckoutController,
		CheckoutCheckController:        *checkCheckoutController,
		CheckoutReprocessingController: *reprocessingCheckoutController,
		WebhookCheckoutController:      *webhookCheckoutController,

		//KitchenFindAllUseCase: *kitchenFindAllUseCase,
		NotifierKitchenUseCase: *notifierkitchenUseCase,

		FindKitchenAllController:  *findKitchenAllController,
		NotifierKitchenController: *notifierKitchenController,
		CookingKitchenController:  *cookKitchenController,
		MonitorKitchenController:  *monitorKitchenController,
		DeliveryKitchenController: *deliveryKitchenController,
	}
}

func (c *Container) Validate() error {
	// check if all dependencies are set
	if c.DB == nil {
		c.Logger.Error("DB is not set")
		return errors.New("db is not set")
	}
	if c.Logger == nil {
		//c.Logger.Error("Logger is not set")
		return errors.New("logger is not set")
	}

	return nil

}
