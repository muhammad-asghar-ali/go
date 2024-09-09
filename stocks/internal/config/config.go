package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		DBHost     string
		DBUser     string
		DBPassword string
		DBName     string
		DBPort     string
		DBTimezone string
		SSLMode    string
	}
)

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading from system environment")
	}

	return &Config{
		DBHost:     getEnv("DATABASE_HOST", "localhost"),
		DBUser:     getEnv("DATABASE_USER", "postgres"),
		DBPassword: getEnv("DATABASE_PASSWORD", ""),
		DBName:     getEnv("DATABASE_NAME", "portfolio"),
		DBPort:     getEnv("DATABASE_POST", "5432"),
		DBTimezone: getEnv("DATABASE_TIMEZONE", "UTC"),
		SSLMode:    getEnv("SSL_MODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
