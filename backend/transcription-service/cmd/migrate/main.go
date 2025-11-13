package main

import (
	"fmt"
	"log"
	"transcription-service/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	log.Println("Loading configuration...")
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	log.Printf("Connecting to database at %s:%d/%s...", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database)
	log.Printf("Using migrations from: %s", cfg.Postgres.MigrationsPath)

	m, err := migrate.New(
		fmt.Sprintf("file://%s", cfg.Postgres.MigrationsPath),
		dsn,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}
	if sErr, err := m.Close(); err != nil {
		log.Printf("Error closing migration instance: %v", err)
	} else if sErr != nil {
		log.Printf("Migration close error: %v", sErr)
	}

	err = m.Up()
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Migration successful!")
}
