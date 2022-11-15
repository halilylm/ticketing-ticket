package utils

import (
	"github.com/caarlos0/env/v6"
	"github.com/halilylm/ticketing-ticket/config"
	"github.com/joho/godotenv"
)

// ReadConfig Read env variables and
// bind them into config struct
func ReadConfig() (*config.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
