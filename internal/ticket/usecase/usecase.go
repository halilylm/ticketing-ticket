package usecase

import (
	"context"
	"database/sql"
	"github.com/halilylm/ticketing-common/http/response"
	"github.com/halilylm/ticketing-ticket/internal/domain"
)

type ticketUsecase struct {
	ticketRepository domain.TicketRepository
}

// A NewTicketUsecase is the factory function
// for ticket usecase layer
func NewTicketUsecase(ticketRepo domain.TicketRepository) TicketUsecase {
	return &ticketUsecase{ticketRepo}
}

// NewTicket implements the business logic for adding
// new tickets to the database
func (t *ticketUsecase) NewTicket(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	createdTicket, err := t.ticketRepository.Insert(ctx, ticket)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}
	return createdTicket, nil
}

// FindTicketByID finds the ticket by ids id
func (t *ticketUsecase) FindTicketByID(ctx context.Context, id int) (*domain.Ticket, error) {
	ticket, err := t.ticketRepository.GetByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NewNotFoundError(err)
		}
		return nil, response.NewInternalServerError(err)
	}
	return ticket, nil
}

// Tickets bring all the tickets from database
func (t *ticketUsecase) Tickets(ctx context.Context, page, limit int) ([]*domain.Ticket, error) {
	tickets, err := t.ticketRepository.All(ctx, page, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NewNotFoundError(err)
		}
		return nil, response.NewInternalServerError(err)
	}
	return tickets, nil
}

// Update the ticket in the database
func (t *ticketUsecase) Update(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	foundTicket, err := t.ticketRepository.Update(ctx, ticket)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NewNotFoundError(err)
		}
		return nil, response.NewInternalServerError(err)
	}
	return foundTicket, nil
}
