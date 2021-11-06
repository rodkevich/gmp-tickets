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

func (d datasource) String() string {
	return "Ticket Postgres"
}

const InsertNewTicket = `
INSERT INTO tickets
(owner_id, name_short, name_ext, description, amount, price, currency, active, perk)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9::enum_tickets_perk_type)
RETURNING id`

func (d *datasource) Create(ctx context.Context, t *ticket.Ticket) (rtn uuid.UUID, err error) {

	tx, err := d.db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer tx.Rollback(ctx)
	err = tx.QueryRow(
		ctx, InsertNewTicket,
		t.OwnerID,
		t.NameShort,
		t.NameExt,
		t.Description,
		t.Amount,
		t.Price,
		t.Currency,
		t.Active,
		t.Perk).Scan(&rtn)
	err = tx.Commit(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

const ReadTicket = `
SELECT
id, owner_id, name_short, name_ext,
description, amount, price, currency, active, ` + ` published_at, created_at, updated_at, deleted_at
FROM tickets
WHERE id = $1
LIMIT 1;
`

// plus = perk::varchar/text/ etc..,
func (d *datasource) Read(ctx context.Context, id uuid.UUID) (*ticket.Ticket, error) {
	var t ticket.Ticket
	err := d.db.QueryRow(
		ctx, ReadTicket,
		id,
	).Scan(
		&t.ID, &t.OwnerID, &t.NameShort, &t.NameExt, &t.Description,
		&t.Amount, &t.Price, &t.Currency, &t.Active,
		// &t.Perk, TODO: find how to implement encode / decode `pgtypes`
		&t.PublishedAt, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (d datasource) List(ctx context.Context, f *ticket.Filter) (rtn []*ticket.Ticket, err error) {
	panic("implement me")
}

func (d datasource) Update(ctx context.Context, t *ticket.Ticket) (err error) {
	panic("implement me")
}

func (d datasource) Delete(ctx context.Context, id uuid.UUID) (err error) {
	panic("implement me")
}

func (d datasource) Search(ctx context.Context, f *ticket.Filter) (rtn []*ticket.Ticket, err error) {
	panic("implement me")
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
