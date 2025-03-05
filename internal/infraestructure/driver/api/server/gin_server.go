package server

import (
	"log"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/di"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GinServer struct {
	router    *gin.Engine
	db        *gorm.DB
	logger    *zap.Logger
	container *di.Container
}

func NewServer(db *gorm.DB, logger *zap.Logger) *GinServer {
	container := di.NewContainer(db, logger)
	err := container.Validate()
	if err != nil {
		log.Fatal("Failed to create container: ", err)
	}

	r := router.SetupRouter(container)

	// Configurar CORS
	//r.Use(corsMiddleware())

	return &GinServer{
		router:    r,
		db:        db,
		logger:    logger,
		container: container,
	}
}

/*
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Responder diretamente às requisições OPTIONS (pré-voo)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
*/

func (s *GinServer) GetDB() *gorm.DB {
	return s.db
}

/*
func (s *GinServer) Initialization() *GinServer {

	//db := setupSQLite()
	//s.db = setupDB()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	productConverter := converter.NewProductConverter()
	orderConverter := converter.NewOrderConverter()
	customerRepo := repositorygorm.NewCustomerRepositoryGorm(s.db)
	productRepo := repositorygorm.NewProductRepositoryGorm(s.db, productConverter)
	orderRepo := repositorygorm.NewOrderRepositoryGorm(s.db, orderConverter)
	checkoutRepo := repositorygorm.NewCheckoutRepositoryGorm(s.db)
	kitchenRepo := repositorygorm.NewKitchenRepositoryGorm(s.db)

	gatewayService := service.NewFakePaymentService()

	g := s.router.Group("/kitchencontrol/api/v1/customers")
	{
		registerController := controllercustomer.NewRegisterCustomerController(ctx, usecasecustomerregister.NewCustomerRegister(customerRepo))
		g.POST("/", registerController.PostRegisterCustomer)

		updateController := controllercustomer.NewUpdateCustomerController(ctx, usecasecustomerupdate.NewCustomerUpdate(customerRepo))
		g.PUT("/:cpf", updateController.PutUpdateCustomer)

		findByCPFController := controllercustomer.NewFindCustomerByCPFController(ctx, usecasecustomerfindbycpf.NewCustomerFindByCPF(customerRepo))
		g.GET("/:cpf", findByCPFController.GetCustomerByCPF)

		findAllController := controllercustomer.NewFindAllCustomersController(ctx, usecasecustomerfindall.NewCustomerFindAll(customerRepo))
		g.GET("/", findAllController.GetAllCustomers)
	}

	p := s.router.Group("/kitchencontrol/api/v1/products")
	{
		registerController := controllerproduct.NewRegisterProductController(ctx, usecaseproductregister.NewProductRegister(productRepo))
		p.POST("/", registerController.PostRegisterProduct)

		findAllController := controllerproduct.NewFindAllProductController(ctx, usecaseproductfindall.NewProductFindAll(productRepo))
		p.GET("/", findAllController.GetAllProducts)

		findByIDController := controllerproduct.NewFindProductByIDController(ctx, usecaseproductfindbyid.NewProductFindByID(productRepo))
		p.GET("/:id", findByIDController.GetProductByID)

		findByCategoryController := controllerproduct.NewFindProductByCategoryController(ctx, usecaseproductfindbycategory.NewProductFindByCategory(productRepo))
		p.GET("/category/:id", findByCategoryController.GetProductByCategory)

		updateController := controllerproduct.NewUpdateProductController(ctx, usecaseproductupdater.NewProductUpdate(productRepo))
		p.PUT("/:id", updateController.PutUpdateProduct)

		deleteController := controllerproduct.NewDeleteProductController(ctx, usecaseproductdelete.NewProductDelete(productRepo))
		p.DELETE("/:id", deleteController.DeleteProduct)

	}

	o := s.router.Group("/kitchencontrol/api/v1/orders")
	{

		orderController := controllerorder.NewCreateOrderController(ctx, usecaseordercreate.NewOrderCreate(orderRepo, customerRepo, productRepo))
		o.POST("/", orderController.PostCreateOrder)

		findAllOrdersController := controllerorder.NewFindAllController(ctx, usecaseorderfindall.NewOrderFindAll(orderRepo))
		o.GET("/", findAllOrdersController.GetAllOrders)

		findByIDController := controllerorder.NewFindOrderByIDController(ctx, usecaseorderfindbyid.NewOrderFindByID(orderRepo))
		o.GET("/:id", findByIDController.GetOrderByID)

		findConfirmedOrdersController := controllerorder.NewFindByParamsConfirmedController(ctx, usecaseorderfindbyparam.NewOrderFindByParams(orderRepo))
		o.GET("/confirmed", findConfirmedOrdersController.GetOrdersConfirmed)
		findNotConfirmedOrdersController := controllerorder.NewFindByParamsController(ctx, usecaseorderfindbyparam.NewOrderFindByParams(orderRepo))
		o.GET("/pending", findNotConfirmedOrdersController.GetOrdersNotConfirmed)
		findPaymentApprovedOrdersController := controllerorder.NewFindByParamsPaymentApprovedController(ctx, usecaseorderfindbyparam.NewOrderFindByParams(orderRepo))
		o.GET("/paid", findPaymentApprovedOrdersController.GetOrdersWithPaymentApproved)

	}

	c := s.router.Group("/kitchencontrol/api/v1/checkouts")
	{
		checkoutController := controllercheckout.NewCreateCheckoutController(ctx, usecasecheckout.NewCheckoutCreate(orderRepo, checkoutRepo, gatewayService, kitchenRepo, productRepo))
		c.POST("/", checkoutController.PostCreateCheckout)
	}

	k := s.router.Group("/kitchencontrol/api/v1/kitchens")
	{
		ktController := controllerkitchen.NewFindKitchenAllController(ctx, usecasekitchen.NewKitchenFindAll(kitchenRepo))
		k.GET("/orders", ktController.GetAllOrdersInKitchen)
	}

	return s
}
*/

func (s *GinServer) Run(port string) {
	if err := s.router.Run(port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func (s *GinServer) GetRouter() *gin.Engine {
	return s.router
}

func (s *GinServer) GetContainer() *di.Container {
	return s.container
}
