package database

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	fmt.Println("Migrating database...")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
	if err := d.Ping(context.Background()); err != nil {
		fmt.Println("Failed to ping database")
		return err
	}
		return fmt.Errorf("count not create the postgres driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", 
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not run up migration: %w", err)
	}
	fmt.Println("successfully migrated database")

	return nil
}
