package server

import (
	"github.com/halilylm/ticketing-ticket/internal/ticket/delivery/http"
	"github.com/halilylm/ticketing-ticket/internal/ticket/repository/postgresql"
	"github.com/halilylm/ticketing-ticket/internal/ticket/usecase"
	"github.com/halilylm/ticketing-ticket/pkg/database"
	"log"
)

func (s *Server) setUpRoutes() {
	// routing definitions
	api := s.app.Group("/api")
	ticket := api.Group("/ticket")
	v1 := ticket.Group("/v1")

	// initialize repositories
	db, err := database.NewPgSqlx(s.cfg)
	if err != nil {
		log.Fatalln(err)
	}
	ticketRepo := postgresql.NewTicketRepository(db)
	// initialize use cases
	ticketUC := usecase.NewTicketUsecase(ticketRepo)

	// initialize handlers
	http.NewTicketHandler(v1, ticketUC)
}
