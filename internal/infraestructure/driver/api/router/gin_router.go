package router

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/di"
	"github.com/gin-gonic/gin"
)

func SetupRouter(container *di.Container) *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())

	// Customer routes
	customerGroup := r.Group("/kitchencontrol/api/v1/customers")
	{
		customerGroup.POST("/", container.RegisterCustomerController.PostRegisterCustomer)
		customerGroup.PUT("/:cpf", container.UpdateCustomerController.PutUpdateCustomer)
		customerGroup.GET("/:cpf", container.FindCustomerByCPFController.GetCustomerByCPF)
		customerGroup.GET("/", container.FindAllCustomersController.GetAllCustomers)
	}

	// Product routes
	productGroup := r.Group("/kitchencontrol/api/v1/products")
	{
		productGroup.POST("/", container.RegisterProductController.PostRegisterProduct)
		productGroup.GET("/", container.FindAllProductController.GetAllProducts)
		productGroup.GET("/:id", container.FindProductByIDController.GetProductByID)
		productGroup.GET("/category/:id", container.FindProductByCategoryController.GetProductByCategory)
		productGroup.PUT("/:id", container.UpdateProductController.PutUpdateProduct)
		productGroup.DELETE("/:id", container.DeleteProductController.DeleteProduct)
	}

	// Order routes
	orderGroup := r.Group("/kitchencontrol/api/v1/orders")
	{
		orderGroup.POST("/", container.CreateOrderController.PostCreateOrder)
		orderGroup.GET("/", container.FindAllOrdersController.GetAllOrders)
		orderGroup.GET("/:id", container.FindOrderByIDController.GetOrderByID)
		orderGroup.GET("/confirmed", container.FindByParamsOrdersConfirmedController.GetOrdersConfirmed)            // listagem: pedido confimado, mas ainda não pago
		orderGroup.GET("/pending", container.FindByParamsOrdersNotConfirmedController.GetOrdersNotConfirmed)        // listagem: ordens aguardando o pagamento
		orderGroup.GET("/paid", container.FindByParamsOrdersPaymentApprovedController.GetOrdersWithPaymentApproved) // listagem: ordens pagas
	}

	// Checkout routes
	checkoutGroup := r.Group("/kitchencontrol/api/v1/checkouts")
	{
		checkoutGroup.POST("/", container.CreateCheckoutController.PostCreateCheckout)
		checkoutGroup.GET("/:id/check/payment", container.CheckoutCheckController.GetCheckPaymentCheckout)            // verifica se determinada ordem foi paga
		checkoutGroup.PUT("/confirmation/payment", container.WebhookCheckoutController.PutConfirmPayment)             // verifica se determinada ordem foi paga
		checkoutGroup.POST("/reprocessing/payment", container.CheckoutReprocessingController.PostReprocessingPayment) // verifica se determinada ordem foi paga
	}

	// Kitchen routes
	kitchenGroup := r.Group("/kitchencontrol/api/v1/kitchens")
	{
		kitchenGroup.GET("/orders/flow", container.FindKitchenAllController.GetAllOrdersInKitchen)
		kitchenGroup.POST("/orders/notifier", container.NotifierKitchenController.PostNotifyKitchen)
		kitchenGroup.POST("/orders/cooking", container.CookingKitchenController.PostCookingKitchen)
		kitchenGroup.GET("/orders/monitor", container.MonitorKitchenController.GetMonitorKitchen)
		kitchenGroup.POST("/orders/delivery", container.DeliveryKitchenController.PostDeliveryKitchen)
	}

	return r
}

// corsMiddleware - no K8s, precisamos tratar o cors
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, Referer, User-Agent")
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
