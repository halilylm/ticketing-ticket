package postgresql

import (
	"context"
	"github.com/halilylm/ticketing-ticket/internal/domain"
	"github.com/jmoiron/sqlx"
)

const (
	insertQuery  = `INSERT INTO tickets (title, description, price, created_at, updated_at) VALUES ($1, $2, $3, now(), now())`
	getByIDQuery = `SELECT * FROM tickets WHERE id=$1`
	updateQuery  = `UPDATE tickets SET title=$1, description=$2, price=$3 WHERE id = $4`
	getAllQuery  = `SELECT * FROM tickets LIMIT $1 OFFSET $2`
)

type ticketRepository struct {
	db *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) domain.TicketRepository {
	return &ticketRepository{db}
}

// Insert saves a new ticket to the postgresql database
func (t *ticketRepository) Insert(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	var createdTicket domain.Ticket
	if err := t.db.QueryRowxContext(ctx, insertQuery, ticket.Title, ticket.Description, ticket.Price).
		StructScan(&createdTicket); err != nil {
		return nil, err
	}
	return &createdTicket, nil
}

// GetByID gets ticket by ids id
func (t *ticketRepository) GetByID(ctx context.Context, id int) (*domain.Ticket, error) {
	var foundTicket domain.Ticket
	if err := t.db.QueryRowxContext(ctx, getByIDQuery, id).StructScan(&foundTicket); err != nil {
		return nil, err
	}
	return &foundTicket, nil
}

// Update the tickets
func (t *ticketRepository) Update(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	var updatedTicket domain.Ticket
	if err := t.db.QueryRowxContext(ctx, updateQuery, ticket.Title, ticket.Description, ticket.Price, ticket.ID).
		StructScan(&updatedTicket); err != nil {
		return nil, err
	}
	return &updatedTicket, nil
}

// All gets all the tickets from the database
// it also provides pagination feature
func (t *ticketRepository) All(ctx context.Context, page, limit int) ([]*domain.Ticket, error) {
	tickets := make([]*domain.Ticket, 0)
	offset := limit * (page - 1)
	rows, err := t.db.QueryxContext(ctx, getAllQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	var foundTicket domain.Ticket
	for rows.Next() {
		if err := rows.StructScan(&foundTicket); err != nil {
			return nil, err
		}
		tickets = append(tickets, &foundTicket)
	}
	return tickets, nil
}
