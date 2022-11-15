package database

import (
	"fmt"
	"github.com/halilylm/ticketing-ticket/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	maxOpenConnections = 10
	connMaxLifetime    = 60 * time.Second
	maxIdleConnections = 10
	connMaxIdleTime    = 20 * time.Second
)

// NewPgSqlx creates a new instance of sqlx
// for postgresql configuration values
func NewPgSqlx(cfg *config.Config) (*sqlx.DB, error) {
	// creating dsn
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlDBName,
		cfg.Postgres.PostgresqlPassword,
	)
	// connect to database
	db, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}
	// set configurations
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(connMaxIdleTime)
	// ping the database
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
