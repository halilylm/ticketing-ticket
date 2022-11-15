package postgresql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/halilylm/ticketing-ticket/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicketRepository_Insert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	defer db.Close()
	dbx := sqlx.NewDb(db, "sqlmock")
	defer dbx.Close()
	ticketRepo := NewTicketRepository(dbx)
	t.Run("insert a new ticket", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"title", "description", "price"}).
			AddRow("liverpool - city", "liverpool - city ticket for next game", 30)
		ticket := domain.Ticket{
			Title:       "liverpool - city",
			Description: "liverpool - city ticket for next game",
			Price:       30,
		}
		mock.ExpectQuery(insertQuery).WithArgs(ticket.Title, ticket.Description, ticket.Price).WillReturnRows(rows)
		createdTicket, err := ticketRepo.Insert(context.Background(), &ticket)
		assert.NoError(t, err)
		assert.NotNil(t, createdTicket)
	})
}

func TestTicketRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	defer db.Close()
	dbx := sqlx.NewDb(db, "sqlmock")
	defer dbx.Close()
	ticketRepo := NewTicketRepository(dbx)
	t.Run("get ticket by id", func(t *testing.T) {
		id := 5
		rows := sqlmock.NewRows([]string{"id", "title", "description", "price"}).
			AddRow(id, "the weeknd concert", "the weeknd concert in istanbul", 250)
		ticket := domain.Ticket{
			ID:          id,
			Title:       "the weeknd concert",
			Description: "the weeknd concert in istanbul",
			Price:       250,
		}
		mock.ExpectQuery(getByIDQuery).WithArgs(id).WillReturnRows(rows)
		foundTicket, err := ticketRepo.GetByID(context.Background(), id)
		assert.NoError(t, err)
		assert.NotNil(t, foundTicket)
		assert.NotNil(t, ticket.Price, foundTicket.Price)
	})
}

func TestTicketRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	defer db.Close()
	dbx := sqlx.NewDb(db, "sqlmock")
	defer dbx.Close()
	ticketRepo := NewTicketRepository(dbx)
	t.Run("updates the ticket", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"title", "description", "price"}).
			AddRow("the weeknd concert", "the weeknd concert in istanbul", 300)
		ticket := domain.Ticket{
			Title:       "the weeknd concert",
			Description: "the weeknd concert in istanbul",
			Price:       300,
		}
		mock.ExpectQuery(updateQuery).WithArgs(ticket.Title, ticket.Description, ticket.Price, ticket.ID).WillReturnRows(rows)
		updatedTicket, err := ticketRepo.Update(context.Background(), &ticket)
		assert.NoError(t, err)
		assert.NotNil(t, updatedTicket)
		assert.Equal(t, ticket.Title, updatedTicket.Title)
	})
}
