package integration

import (
	"github.com/halilylm/ticketing-ticket/config"
	"github.com/halilylm/ticketing-ticket/internal/server"
	"net/http"
	"time"
)

// CreateServer creates a new web server
// for testing
func CreateServer() {
	srv := server.New(&config.Config{
		Server: config.Server{
			ServerHost:         "localhost",
			ServerPort:         5000,
			ServerReadTimeout:  5 * time.Second,
			ServerWriteTimeout: 5 * time.Second,
			ServerIdleTimeout:  5 * time.Second,
		},
	})
	go srv.Run()
	// wait until server boots up to run the tests
	count := 0
	for {
		_, err := http.Get("http://localhost:5000")
		if err == nil {
			break
		}
		time.Sleep(10 * time.Millisecond)
		count++
		if count > 20 {
			panic("failed to boots up the server")
		}
	}
}
