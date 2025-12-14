package config

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

type Config struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

func NewDatabase(cfg Config) (*sql.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
    )

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("error opening database: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("error connecting to database: %w", err)
    }

    log.Println("Database connected successfully")
    return db, nil
}