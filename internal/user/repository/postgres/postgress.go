package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rodkevich/gmp-tickets/internal/user"
	"log"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/rodkevich/gmp-tickets/internal/configs"
)

type datasource struct {
	db *pgxpool.Pool
}

func (d datasource) Create(ctx context.Context, arg *user.User) (id string, err error) {
	panic("implement me")
}

func (d datasource) List(ctx context.Context, f *user.Filter) ([]*user.User, error) {
	panic("implement me")
}

func (d datasource) Read(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	panic("implement me")
}

func (d datasource) Update(ctx context.Context, user *user.User) error {
	panic("implement me")
}

func (d datasource) Delete(ctx context.Context, userID uuid.UUID) error {
	panic("implement me")
}

func (d datasource) Search(ctx context.Context, req *user.Filter) ([]*user.User, error) {
	panic("implement me")
}

const (
	Select     = "SELECT () FROM tickets"
	SelectByID = "SELECT () FROM tickets"
	Update     = "UPDATE"
	Delete     = "DELETE"
	Search     = "SELECT () FROM tickets"
)

func (d datasource) String() string {
	return "User Postgres"
}

func NewDatasource(cfg *configs.Configs) (user.Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var connString string
	connString = fmt.Sprintf(
		"%s://%s/%s?sslmode=%s&user=%s&password=%s&pool_max_conns=%v",
		cfg.Database.Driver,
		cfg.Database.Host,
		cfg.Database.Name,
		cfg.Database.SslMode,
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.MaxPoolConnections,
	)

	pool, err := pgxpool.Connect(ctx, connString)

	if err != nil {
		log.Printf("Unable to connect database: %v\n", err)
		return nil, err
	}
	log.Printf("New PG datasource connected to: %v", connString)
	for {
		_, err := pool.Exec(context.Background(), "SELECT '"+cfg.Api.Name+".user'::regclass")
		if err == nil {
			log.Println("Database is ready")
			return &datasource{pool}, nil
		}

		base, plug := time.Second, time.Minute
		for backoff := base; err != nil; backoff <<= 1 {
			if backoff > plug {
				backoff = plug
			}
			log.Println("Test query: retrying...")

			/* #nosec */
			step := rand.Int63n(int64(backoff * 3))
			sleep := base + time.Duration(step)
			time.Sleep(sleep)

			_, err := pool.Exec(context.Background(), "SELECT '"+cfg.Api.Name+".user'::regclass")
			if err == nil {
				log.Println("Database is ready")
				return nil, err
			}
		}
	}
}
