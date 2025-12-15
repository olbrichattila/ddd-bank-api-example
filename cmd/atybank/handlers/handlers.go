package handlers

import (
	"atybank/cmd/atybank/services"
	"atybank/internal/api/health"
	"atybank/internal/api/v1/account"
	"atybank/internal/api/v1/transaction"
	"atybank/internal/api/v1/user"
	configRepository "atybank/internal/infrastructure/config"
	"database/sql"
)

type Handlers struct {
	User        *user.Handler
	Account     *account.Handler
	Transaction *transaction.Handler
	Health      *health.Handler
}

func New(db *sql.DB, cfg configRepository.Config, services *services.Services) (*Handlers, error) {
	healthHandler, err := health.New(db)
	if err != nil {
		return nil, err
	}

	userHandler, err := user.New(cfg, services.Logger, services.User)
	if err != nil {
		return nil, err
	}

	accountHandler, err := account.New(services.Logger, services.Account)
	if err != nil {
		return nil, err
	}

	transactionHandler, err := transaction.New(services.Logger, services.Transaction)
	if err != nil {
		return nil, err
	}

	return &Handlers{
		User:        userHandler,
		Account:     accountHandler,
		Transaction: transactionHandler,
		Health:      healthHandler,
	}, nil
}
