package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	SlackBotToken string
	ChannelIDs    []string
}

var (
	instance *Config
	once     sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		instance = &Config{
			SlackBotToken: get("SLACK_BOT_TOKEN"),
			ChannelIDs:    parse(get("CHANNEL_ID")),
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

func parse(channels string) []string {
	if channels == "" {
		return []string{}
	}

	return strings.Split(channels, ",")
}
