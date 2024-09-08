package config

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func InitDB(cfg *Config) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode, cfg.DBTimezone)

	var err error
	db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Database connection successfully established")
}

func DB() *sqlx.DB {
	return db
}
