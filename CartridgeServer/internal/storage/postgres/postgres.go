package postgres

import (
	"CartridgeServer/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type Database struct {
	db *pgxpool.Pool
}

func New(cfg config.DatabaseConfig) *Database {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return &Database{
		db: db,
	}
}
