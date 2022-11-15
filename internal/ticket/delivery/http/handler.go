package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/halilylm/ticketing-common/http/response"
	"github.com/halilylm/ticketing-ticket/internal/domain"
	"github.com/halilylm/ticketing-ticket/internal/ticket/usecase"
	"net/http"
)

type TicketHandler struct {
	ticketUC usecase.TicketUsecase
}

func NewTicketHandler(f fiber.Router, ticketUC usecase.TicketUsecase) {
	handler := TicketHandler{ticketUC: ticketUC}
	f.Post("/", handler.NewTicket)
}

// NewTicket is the responsible for creating new tickets
func (t *TicketHandler) NewTicket(ctx *fiber.Ctx) error {
	var ticket domain.Ticket
	var httpErr response.HTTPError
	// bind body to ticket object
	if err := ctx.BodyParser(&ticket); err != nil {
		httpErr = response.NewBadRequestError(err.Error(), err)
		return ctx.Status(httpErr.Status).JSON(httpErr)
	}
	// call usecase to create a new ticket
	createdTicket, err := t.ticketUC.NewTicket(ctx.Context(), &ticket)
	if err != nil {
		httpErr = response.ParseHTTPError(err)
		return ctx.Status(httpErr.Status).JSON(httpErr)
	}
	return ctx.Status(http.StatusCreated).JSON(createdTicket)
}
