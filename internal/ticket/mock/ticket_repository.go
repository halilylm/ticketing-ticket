package mock

import (
	"context"
	"database/sql"
	"github.com/halilylm/ticketing-ticket/internal/domain"
)

type TicketRepository struct {
	tickets []*domain.Ticket
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{tickets: []*domain.Ticket{}}
}

func (m *TicketRepository) Insert(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	m.tickets = append(m.tickets, ticket)
	return ticket, nil
}

func (m *TicketRepository) GetByID(ctx context.Context, id int) (*domain.Ticket, error) {
	for _, ticket := range m.tickets {
		if ticket.ID == id {
			return ticket, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (m *TicketRepository) Update(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	for i, foundTicket := range m.tickets {
		if foundTicket.ID == ticket.ID {
			m.tickets[i] = ticket
			return ticket, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (m *TicketRepository) All(ctx context.Context, page, limit int) ([]*domain.Ticket, error) {
	offset := limit * (page - 1)
	tickets := m.tickets[offset:(offset + limit)]
	if len(tickets) == 0 {
		return nil, sql.ErrNoRows
	}
	return tickets, nil
}
