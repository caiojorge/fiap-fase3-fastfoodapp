package connection

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	db       *gorm.DB
}

func NewDB(host, port, user, password, dbName string) *DB {
	return &DB{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}
}

func (d *DB) GetConnection(dbType string) *gorm.DB {

	if dbType == "sqlite" {
		d.db = d.setupSQLite()
	}

	if dbType == "mysql" {
		d.db = d.setupMysql()
	}

	if d.db == nil {
		log.Fatalf("Database type %s not supported", dbType)
	}

	return d.db
}

func (d *DB) setupSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Customer{})
	return db
}

func (d *DB) setupMysql() *gorm.DB {

	//dbPassword := os.Getenv("DB_PASSWORD")

	dbHost := getEnv("DB_HOST", d.Host)
	dbPort := getEnv("DB_PORT", d.Port)
	dbUser := getEnv("DB_USER", d.User)
	dbPass := getEnv("DB_PASS", d.Password)
	dbName := getEnv("DB_NAME", d.DBName)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var db *gorm.DB
	var err error

	for i := 0; i < 20; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to database. Retrying in 5 seconds... (%d/%d)\n", i+1, 10)
		fmt.Printf("host: %s, port: %s, user: %s, pass: %s, name: %s\n", dbHost, dbPort, dbUser, dbPass, dbName)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic("failed to connect database after multiple attempts")
	}

	fmt.Println("Successfully connected to the database")

	return db
}

// getEnv retorna o valor da variável de ambiente ou um valor default se não estiver definido
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
