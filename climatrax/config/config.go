package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	instance *Config
	once     sync.Once
)

type (
	Config struct {
		ApiKey string
	}
)

func LoadConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		instance = &Config{
			ApiKey: get("API_KEY"),
		}
	})

	return instance
}

func get(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}

	return value
}

func GetConfig() *Config {
	return LoadConfig()
}

func (c *Config) GetApiKey() string {
	return c.ApiKey
}
