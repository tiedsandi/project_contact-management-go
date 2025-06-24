package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")

	envFile := ".env"
	if env == "production" {
		envFile = ".env.production"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("⚠️  No %s file found, continuing...", envFile)
	}
}
