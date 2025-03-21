package client

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"parser/internal/config"
)

type PGClient struct {
	DB  *sql.DB
	ctx context.Context
}

func NewPGClient(ctx context.Context, cfg *config.DBConfig) (*PGClient, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	return &PGClient{
		DB:  db,
		ctx: ctx,
	}, nil
}

func (client *PGClient) QueryRow(query string, args ...interface{}) *sql.Row {
	return client.DB.QueryRowContext(client.ctx, query, args...)
}

func (client *PGClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return client.DB.QueryContext(client.ctx, query, args...)
}

func (client *PGClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return client.DB.ExecContext(client.ctx, query, args...)
}
