package test_containers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	DB        *sqlx.DB
	Container *postgres.PostgresContainer
}

func NewPostgresTestContainer(ctx context.Context) *PostgresContainer {
	fmt.Println("Starting Postgres container...")
	postgresContainer, err := postgres.Run(ctx,
		"postgres",
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start postgres container: %w", err)
	}

	// Create a connection string
	dbUrl, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("failed to open database connection: %w", err)
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("failed to open database connection: %w", err)
	}

	//Run goose migrations
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to set goose dialect: %w", err)
	}

	fmt.Println("Running migrations...")
	if err := goose.Up(
		db,
		"../../migrations",
	); err != nil {
		log.Fatalf("Failed to run migrations", err)
	}

	return &PostgresContainer{
		DB:        sqlx.NewDb(db, "postgres"),
		Container: postgresContainer,
	}
}
