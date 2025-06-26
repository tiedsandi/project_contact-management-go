package config

import (
	"fmt"
	"log"
	"os"

	"github.com/tiedsandi/project_contact-management-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	env := os.Getenv("APP_ENV")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	sslmode := "disable"
	if env == "production" {
		sslmode = "require"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Contact{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}
	fmt.Println("✅ Database migrated successfully!")
}

// ======================== ini yang sebelum di buat env.go ===================
// package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDB() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("⚠️  No .env file found (if you're in production, this is expected)")
// 	}

// 	env := os.Getenv("APP_ENV")
// 	host := os.Getenv("DB_HOST")
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbname := os.Getenv("DB_NAME")
// 	port := os.Getenv("DB_PORT")

// 	sslmode := "disable"
// 	if env == "production" {
// 		sslmode = "require"
// 	}

// 	dsn := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
// 		host, user, password, dbname, port, sslmode,
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("❌ Failed to connect to database:", err)
// 	}

// 	fmt.Println("✅ Connected to PostgreSQL!")
// 	DB = db
// }

// ======================== ini yang awal ===================
// package config

// import (
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDB() {
// 	dsn := "host=localhost user=postgres password=123 dbname=contact_db port=5432 sslmode=disable"
// 	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}

// 	DB = database
// }
