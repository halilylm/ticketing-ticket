package domain

import (
	"context"
	"time"
)

// Ticket represents ticket domain
type Ticket struct {
	ID          int       `json:"-" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Price       int       `json:"price" db:"price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// TicketRepository is the contract for
// accessing the database layer
// for ticket domain
type TicketRepository interface {
	Insert(ctx context.Context, ticket *Ticket) (*Ticket, error)
	GetByID(ctx context.Context, id int) (*Ticket, error)
	Update(ctx context.Context, ticket *Ticket) (*Ticket, error)
	All(ctx context.Context, page, limit int) ([]*Ticket, error)
}
