package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// mockControllerGetCustomerByCPF is a mock controller function
func mockControllerGetCustomerByCPF(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"cpf": "12345678901"})
}

// Setup the server and define URLs with mock controller
func setupTestServer() *GinServer {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//db := setupMysql()

	// Migrar o esquema
	err = db.AutoMigrate(&model.Customer{}, &model.Product{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		panic("failed to migrate database")
	}
	gin.SetMode(gin.TestMode)
	logger, _ := zap.NewProduction()
	server := NewServer(db, logger)
	server.router.GET("/customer", mockControllerGetCustomerByCPF)
	return server
}

// TestServerRoutes tests whether the routing is properly configured
func TestServerRoutes(t *testing.T) {
	server := setupTestServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customer", nil)
	server.GetRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "12345678901")
}
