package config

import "time"

type Config struct {
	Server
	Postgres
}

type Server struct {
	ServerHost         string        `env:"SERVER_HOST,required"`
	ServerPort         int           `env:"SERVER_PORT,required"`
	ServerReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT,required"`
	ServerWriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT,required"`
	ServerIdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT,required"`
}

type Postgres struct {
	PostgresqlHost     string `env:"POSTGRES_HOST,required"`
	PostgresqlPort     int    `env:"POSTGRES_PORT,required"`
	PostgresqlUser     string `env:"POSTGRES_USER,required"`
	PostgresqlDBName   string `env:"POSTGRES_DB_NAME,required"`
	PostgresqlPassword string `env:"POSTGRES_PASSWORD,required"`
}
