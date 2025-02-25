package main

import (
	"database/sql"
	"fmt"
	"embed"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed sql/schema
var migrations embed.FS

func runMigrations(dbURL string) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	goose.SetBaseFS(migrations)

	if err := goose.Up(db, "sql/schema"); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	fmt.Println("Database migrations applied successfully.")
	return nil
}

