package config

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(cfg *Config) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode, cfg.DBTimezone)

	var err error
	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB.SetConnMaxLifetime(time.Hour)
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Database connection successfully established")
}

func GetDB() (*sqlx.DB, error) {
	if DB == nil {
		return nil, fmt.Errorf("database connection has not been initialized")
	}

	return DB, nil
}
