package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rodkevich/gmp-tickets/internal/configs"
	"github.com/rodkevich/gmp-tickets/internal/ticket"
	"log"
	"math/rand"
	"time"
)

type datasource struct {
	db *pgxpool.Pool
}

const (
	Select     = "SELECT () FROM tickets"
	SelectByID = "SELECT () FROM tickets"
	Update     = "UPDATE"
	Delete     = "DELETE"
	Search     = "SELECT () FROM tickets"
)

// INSERT INTO public.tickets (id, name, full_name, description, status, owner_id, amount,
// price, currency, created_at, updated_at, deleted_at,
// published_at)
// VALUES (DEFAULT, 'test_ticket', 'test_ticket_full_name', 'test_ticket_description',
// 'active', 1, 22, 333.33, 251, DEFAULT, null, null, null);

const InsertNewPhoto = `
INSERT INTO public.photo (id, ticket_id, is_main, presented, mime_type, size_kb)
VALUES (
DEFAULT, $1, DEFAULT, DEFAULT, $2, $3);
`

const InsertNewTicket = `
INSERT INTO ticket
(name, full_name, description, status, owner_id, amount, price, currency)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;
`

func (d *datasource) Create(ctx context.Context, arg *ticket.Ticket) (id uuid.UUID, err error) {

	tx, err := d.db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer tx.Rollback(ctx)
	err = tx.QueryRow(
		ctx, InsertNewTicket,
		arg.Name,
		arg.FullName,
		arg.Description,
		arg.Status,
		arg.OwnerID,
		arg.Amount,
		arg.Price,
		arg.Currency).Scan(&id)
	err = tx.Commit(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (d *datasource) List(ctx context.Context, f *ticket.Filter) ([]*ticket.Ticket, error) {
	panic("implement me")
}

func (d *datasource) Read(ctx context.Context, ticketID uuid.UUID) (*ticket.Ticket, error) {
	panic("implement me")
}

func (d *datasource) Update(ctx context.Context, ticket *ticket.Ticket) error {
	panic("implement me")
}

func (d *datasource) Delete(ctx context.Context, ticketID uuid.UUID) error {
	panic("implement me")
}

func (d *datasource) Search(ctx context.Context, req *ticket.Filter) ([]*ticket.Ticket, error) {
	panic("implement me")
}

func (d datasource) String() string {
	return "Ticket Postgres"
}

func NewDatasource(cfg *configs.Configs, pool *pgxpool.Pool) (ticket.Repository, error) {

	for {
		_, err := pool.Exec(context.Background(), "SELECT '"+cfg.Database.Schema+".tickets'::regclass")
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

			_, err := pool.Exec(context.Background(), "SELECT '"+cfg.Database.Schema+".tickets'::regclass")
			if err == nil {
				log.Println("Database is ready")
				return &datasource{pool}, nil
			}
		}
	}
}
