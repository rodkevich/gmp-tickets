package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rodkevich/gmp-tickets/internal/user"
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
