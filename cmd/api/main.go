package main

import (
	"github.com/halilylm/ticketing-ticket/internal/server"
	"github.com/halilylm/ticketing-ticket/pkg/utils"
	"log"
)

func main() {
	cfg, err := utils.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	// run the application
	srv := server.New(cfg)
	srv.Run()
}
