package handlers

import (
	"eaglebank/cmd/eaglebank/services"
	"eaglebank/internal/api/v1/account"
	"eaglebank/internal/api/v1/transaction"
	"eaglebank/internal/api/v1/user"
)

type Handlers struct {
	User        *user.Handler
	Account     *account.Handler
	Transaction *transaction.Handler
}

func New(services *services.Services) (*Handlers, error) {
	userHandler, err := user.New(services.User)
	if err != nil {
		return nil, err
	}

	accountHandler, err := account.New(services.Account)
	if err != nil {
		return nil, err
	}

	transactionHandler, err := transaction.New(services.Transaction)
	if err != nil {
		return nil, err
	}

	return &Handlers{
		User:        userHandler,
		Account:     accountHandler,
		Transaction: transactionHandler,
	}, nil
}
