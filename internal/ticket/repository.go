package ticket

import (
	"context"
	"github.com/google/uuid"
)

// Repository ...
type Repository interface {
	Create(ctx context.Context, t *Ticket) (rtn uuid.UUID, err error)
	List(ctx context.Context, f *Filter) (rtn []*Ticket, err error)
	Read(ctx context.Context, id uuid.UUID) (rtn *Ticket, err error)
	Update(ctx context.Context, t *Ticket) (err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
	Search(ctx context.Context, f *Filter) (rtn []*Ticket, err error)
}
