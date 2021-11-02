package rest

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodkevich/gmp-tickets/internal/ticket"
)

type ticketService struct {
	repo ticket.Repository
}

func NewTicketService(repo ticket.Repository) *ticketService {
	return &ticketService{repo: repo}
}

func (t ticketService) Create(ctx context.Context, ticket ticket.CreationRequest) (*ticket.Ticket, error) {
	panic("implement me")
}

func (t ticketService) List(ctx context.Context, opt *ticket.Fields) ([]*ticket.Ticket, error) {
	panic("implement me")
}

func (t ticketService) Read(ctx context.Context, ticketID uuid.UUID) (*ticket.Ticket, error) {
	panic("implement me")
}

func (t ticketService) Update(ctx context.Context, ticket *ticket.Ticket) (*ticket.Ticket, error) {
	panic("implement me")
}

func (t ticketService) Delete(ctx context.Context, ticketID uuid.UUID) error {
	panic("implement me")
}

func (t ticketService) Search(ctx context.Context, opt *ticket.Fields) ([]*ticket.Ticket, error) {
	panic("implement me")
}
