package rest

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/rodkevich/gmp-tickets/internal/ticket"
)

type ticketService struct {
	repo ticket.Repository
}

func NewTicketService(repo ticket.Repository) *ticketService {
	return &ticketService{repo: repo}
}

func (ts ticketService) Create(ctx context.Context, t *ticket.Ticket) (rtn *ticket.Ticket, err error) {
	ID, err := ts.repo.Create(ctx, t)
	if err != nil {
		log.Println(err)
		return
	}
	rtn, err = ts.Read(ctx, ID)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (ts ticketService) List(ctx context.Context, f *ticket.Filter) ([]*ticket.Ticket, error) {
	panic("implement me")
}

func (ts ticketService) Read(ctx context.Context, ticketID uuid.UUID) (*ticket.Ticket, error) {
	panic("implement me")
}

func (ts ticketService) Update(ctx context.Context, ticket *ticket.Ticket) (*ticket.Ticket, error) {
	panic("implement me")
}

func (ts ticketService) Delete(ctx context.Context, ticketID uuid.UUID) error {
	panic("implement me")
}

func (ts ticketService) Search(ctx context.Context, f *ticket.Filter) ([]*ticket.Ticket, error) {
	panic("implement me")
}
