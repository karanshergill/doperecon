package database

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    dbHost := os.Getenv("POSTGRES_HOST")
    dbPort := os.Getenv("POSTGRES_PORT")
    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")

    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    config, err := pgxpool.ParseConfig(dsn)
    if err != nil {
        log.Fatalf("Unable to parse connection string: %v", err)
    }

    config.MaxConns = 10
    config.MinConns = 1
    config.HealthCheckPeriod = 2 * time.Minute

    Pool, err = pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        log.Fatalf("Unable to create connection pool: %v", err)
    }

    if err := Pool.Ping(context.Background()); err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    fmt.Println("Database connected successfully")
}

func CloseDB() {
    if Pool != nil {
        Pool.Close()
        fmt.Println("Database connection closed")
    }
}