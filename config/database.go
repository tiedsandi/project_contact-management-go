package config

import (
	"fmt"
	"log"
	"os"

	"github.com/tiedsandi/project_contact-management-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah instance global dari GORM DB yang akan digunakan di seluruh aplikasi
var DB *gorm.DB

// ConnectDB menghubungkan aplikasi ke database PostgreSQL menggunakan environment variable
func ConnectDB() {
	// Ambil nilai-nilai konfigurasi dari environment
	env := os.Getenv("APP_ENV")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Tentukan sslmode: non-production disable, production require
	sslmode := "disable"
	if env == "production" {
		sslmode = "require"
	}

	// Format Data Source Name (DSN) untuk koneksi PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	// Buka koneksi ke database menggunakan GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
	DB = db
}

// Migrate menjalankan migrasi otomatis ke database menggunakan model dari package `models`
// Fungsi ini aman untuk data existing karena bersifat non-destruktif
func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},    // Migrasi tabel users
		&models.Contact{}, // Migrasi tabel contacts
		&models.Address{}, // Migrasi tabel addresses
	)
	if err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	fmt.Println("✅ Database migrated successfully!")
}

// ========================================= ini yang beluim ada migrationnya tadi abis ubah modelnya jadi kalo mau gampang drop dbnya buat baru kalo mau yang langsung ada di diatas

// package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/tiedsandi/project_contact-management-go/models"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDB() {
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

// func Migrate() {
// 	err := DB.AutoMigrate(
// 		&models.User{},
// 		&models.Contact{},
// 		&models.Address{},
// 	)
// 	if err != nil {
// 		log.Fatal("❌ Failed to migrate database:", err)
// 	}
// 	fmt.Println("✅ Database migrated successfully!")
// }

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
