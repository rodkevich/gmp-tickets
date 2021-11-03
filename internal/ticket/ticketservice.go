package ticket

import (
	"context"

	"github.com/google/uuid"
)

type UsageSchema interface {
	Create(ctx context.Context, t *Ticket) (rtn *Ticket, err error)
	List(ctx context.Context, f *Filter) ([]*Ticket, error)
	Read(ctx context.Context, ticketID uuid.UUID) (*Ticket, error)
	Update(ctx context.Context, ticket *Ticket) (*Ticket, error)
	Delete(ctx context.Context, ticketID uuid.UUID) error
	Search(ctx context.Context, f *Filter) ([]*Ticket, error)
}
