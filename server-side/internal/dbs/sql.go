package dbs

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
)

func ConnectPostgres() (*sql.DB, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	url := cfg.NeonDBURL
	if url == "" {
		return nil, fmt.Errorf("NEON__DB__URL is not configured")
	}

	pgxConfig, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("invalid Neon database URL: %w", err)
	}
	pgxConfig.DialFunc = func(ctx context.Context, _, address string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "tcp4", address)
	}

	db := sql.OpenDB(stdlib.GetConnector(*pgxConfig))

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to Neon: %w", err)
	}
	fmt.Println("Neon connected successfully!")

	return db, nil
}
