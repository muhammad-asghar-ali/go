package main

import (
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"stocks/internal/config"
)

var (
	//go:embed migrations/*.sql
	src embed.FS
)

func main() {
	cfg := config.LoadConfig()

	config.InitDB(cfg)

	db := config.DB()

	dir, err := iofs.New(src, "migrations")
	if err != nil {
		log.Fatal("error creating IOFS migration source:", err)
		return
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatal("error creating Postgres driver for migration:", err)
		return
	}

	migrations, err := migrate.NewWithInstance("iofs", dir, "postgres", driver)
	if err != nil {
		log.Fatal("unable to create migration instance:", err)
		return
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("error applying migrations:", err)
		return
	}

	log.Println("Migrations applied successfully.")
}
