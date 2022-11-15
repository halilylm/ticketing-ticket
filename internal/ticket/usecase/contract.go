package usecase

import (
	"context"
	"github.com/halilylm/ticketing-ticket/internal/domain"
)

// TicketUsecase is the contract for
// usecase layer
type TicketUsecase interface {
	NewTicket(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error)
	FindTicketByID(ctx context.Context, id int) (*domain.Ticket, error)
	Tickets(ctx context.Context, page, limit int) ([]*domain.Ticket, error)
	Update(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error)
}
