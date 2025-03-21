package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/caiojorge/fiap-challenge-ddd/docs"
	connection "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/db/connection"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/db/migration"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/robot"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/server"
	payment "github.com/caiojorge/fiap-challenge-ddd/internal/shared/fake"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Fiap Fase 3 Challenge AWS - 9SOAT
// @version 1.0
// @description This is fiap fase 3 challenge project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /kitchencontrol/api/v1
func main() {

	_ = godotenv.Load() // Carrega o .env se não estiver definido em variáveis de ambiente

	hostname := os.Getenv("HOST_NAME")
	hostport := os.Getenv("HOST_PORT_K8S")
	hostportContainer := os.Getenv("HOST_PORT_CONTAINER")

	gin.SetMode(gin.ReleaseMode)

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any

	conn := connection.NewConnection("mysql", logger)
	db, err := conn.Execute()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	server := server.NewServer(db, logger)

	logger.Info("Server Initialized")

	// Migrate the schema
	logger.Info("Startin Migration")

	m := migration.NewMigration(db)
	if err := m.Execute(); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	logger.Info("Migration executed successfully")

	fmt.Println("hostname: ", hostname)
	fmt.Println("hostport: ", hostport)
	fmt.Println("hostportContainer: ", hostportContainer)

	setupSwagger(hostname, hostport, server, logger)

	// Iniciar o "cron"
	c := cron.New()

	robot := robot.NewNotifierRobot(server, logger)

	// Simulador do sistema da cozinha que puxa as ordens pagas para inicio do preparo
	c.AddFunc("@every 30s", func() {
		logger.Info("Notify kitchen")
		//server.GetContainer().NotifierKitchenUseCase.Notify(context.Background())
		err := robot.Notify(context.Background())
		if err != nil {
			logger.Error("Failed to notify kitchen", zap.Error(err))
		}
	})

	// Start do cron em background
	c.Start()

	server.GetRouter().POST("/instore/orders/qr/seller/collectors/:collectorID/pos/:posID/qrs", payment.PostPaymentFake)

	server.Run(":" + hostportContainer)

}

func setupSwagger(hostname string, hostport string, server *server.GinServer, logger *zap.Logger) {

	_hostname := hostname
	_hostname = strings.TrimPrefix(_hostname, "http://")
	_hostname = strings.TrimPrefix(_hostname, "https://")

	//docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", _hostname, hostport)
	docs.SwaggerInfo.Host = _hostname
	docs.SwaggerInfo.BasePath = "/kitchencontrol/api/v1"

	//swaggerURL := fmt.Sprintf("http://%s:%s/kitchencontrol/api/v1/docs/doc.json", hostname, hostport)
	//swaggerURL := fmt.Sprintf("%s:%s/kitchencontrol/api/v1/docs/doc.json", _hostname, hostport)
	swaggerURL := fmt.Sprintf("%s/kitchencontrol/api/v1/docs/doc.json", _hostname)

	fmt.Println("swaggerURL: ", swaggerURL)

	server.GetRouter().GET("/kitchencontrol/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	logger.Info("Server running on " + hostname + ":" + hostport)
	logger.Info("swagger running on " + swaggerURL)
}
