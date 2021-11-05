package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rodkevich/gmp-tickets/internal/configs"
)

func NewConnectionPool(ctx context.Context, cfg *configs.Configs) (pool *pgxpool.Pool, err error) {

	var connString string
	connString = fmt.Sprintf(
		"%s://%s/%s?sslmode=%s&user=%s&password=%s&pool_max_conns=%v&pool_min_conns=%v",
		cfg.Database.Driver,
		cfg.Database.Host,
		cfg.Database.Name,
		cfg.Database.SslMode,
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.MaxPoolConnections,
		cfg.Database.MinPoolConnections,
	)
	pool, err = pgxpool.Connect(ctx, connString)

	if err != nil {
		log.Printf("Unable to connect database: %v\n", err)
		return nil, err
	}
	log.Printf("New PG datasource connected to: %v", connString)
	return pool, nil
}
