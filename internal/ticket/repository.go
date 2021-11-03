package ticket

import (
	"context"
	"github.com/google/uuid"
)

// Repository ...
type Repository interface {
	Create(ctx context.Context, arg *Ticket) (ticketID uuid.UUID, err error)
	List(ctx context.Context, f *Filter) ([]*Ticket, error)
	Read(ctx context.Context, ticketID uuid.UUID) (*Ticket, error)
	Update(ctx context.Context, ticket *Ticket) error
	Delete(ctx context.Context, ticketID uuid.UUID) error
	Search(ctx context.Context, req *Filter) ([]*Ticket, error)
}
