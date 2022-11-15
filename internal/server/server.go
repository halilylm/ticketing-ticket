package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/halilylm/ticketing-ticket/config"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
)

type Server struct {
	app     *fiber.App
	cfg     *config.Config
	address string
}

// New creates a new server
func New(cfg *config.Config) *Server {
	address := net.JoinHostPort(cfg.ServerHost, strconv.Itoa(cfg.ServerPort))
	app := fiber.New(fiber.Config{
		ReadTimeout:           cfg.ServerReadTimeout,
		WriteTimeout:          cfg.ServerWriteTimeout,
		IdleTimeout:           cfg.ServerIdleTimeout,
		DisableStartupMessage: false,
	})
	return &Server{
		app:     app,
		cfg:     cfg,
		address: address,
	}
}

// Run will start the webserver then block
// until signal interrupt signal received
// then gracefully shutdown the server
// stop new incoming requests
// will stay alive until requests
// being served
func (s *Server) Run() {
	// registering the routes
	log.Println("registering routes")
	s.setUpRoutes()
	log.Println("starting the server!")
	// start the webserver
	go func() {
		if err := s.app.Listen(s.address); err != nil {
			log.Fatalln(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("gracefully shutting down the app...")
	if err := s.app.Shutdown(); err != nil {
		log.Fatalln(err)
	}
	log.Println("shut down the app!")
}
