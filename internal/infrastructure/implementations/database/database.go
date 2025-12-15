package database

import (
	"database/sql"
	"fmt"

	"atybank/internal/infrastructure/config"
	databaseInterface "atybank/internal/infrastructure/database"

	_ "github.com/lib/pq"
)

func New(cfg config.Config) (databaseInterface.Database, error) {
	if cfg == nil {
		return nil, fmt.Errorf("nil config in database creation")
	}

	return &dbase{
		cfg: cfg,
	}, nil
}

type dbase struct {
	cfg config.Config
}

func (d *dbase) Connect() (*sql.DB, error) {
	dsn := d.cfg.GetDBURL()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
