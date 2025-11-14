package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"electronic_shop/models"
	"electronic_shop/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func main() {
	var dsn string

	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		dsn = databaseURL
	} else {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")

		dsn = "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword +
		" dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Category{},
		&models.Supplier{},
		&models.Product{},
		&models.Position{},
		&models.Store{},
		&models.Employee{},
		&models.Client{},
		&models.Sale{},
		&models.SaleItem{},
		&models.Review{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	router := gin.Default()
	routes.SetupRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
