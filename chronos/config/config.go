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
		SlackBotToken string
		SlackAppToken string
	}
)

func LoadConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		instance = &Config{
			SlackBotToken: get("SLACK_BOT_TOKEN"),
			SlackAppToken: get("SLACK_APP_TOKEN"),
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
