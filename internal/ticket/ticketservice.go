package ticket

import (
	"context"

	"github.com/google/uuid"
)

type CreationRequest struct {
}

type Fields struct {
}

type UsageSchema interface {
	Create(ctx context.Context, ticket CreationRequest) (*Ticket, error)
	List(ctx context.Context, f *Fields) ([]*Ticket, error)
	Read(ctx context.Context, ticketID uuid.UUID) (*Ticket, error)
	Update(ctx context.Context, ticket *Ticket) (*Ticket, error)
	Delete(ctx context.Context, ticketID uuid.UUID) error
	Search(ctx context.Context, f *Fields) ([]*Ticket, error)
}
