package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	fmt.Printf("DB_HOST: %s, DB_PORT: %s, DB_USERNAME: %s, DB_PASSWORD: %s, DB_DB: %s, SSL_MODE: %s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DB"),
		os.Getenv("SSL_MODE"),
	)

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DB"),
		os.Getenv("SSL_MODE"),
	)

	retry := 0
	var err error
	for retry < 5 {
		dbConn, err := sqlx.Connect("postgres", connectionString)
		if err != nil {
			retry++
			time.Sleep(1 * time.Second)
			continue
		}

		return &Database{Client: dbConn}, nil
	}
	return &Database{}, fmt.Errorf("failed to connect to database: %v,", err)
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
