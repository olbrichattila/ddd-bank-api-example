package services

import (
	"fmt"

	"eaglebank/cmd/eaglebank/repositories"
	"eaglebank/internal/application/account"
	"eaglebank/internal/application/transaction"
	"eaglebank/internal/application/user"
)

type Services struct {
	User        user.User
	Account     account.Account
	Transaction transaction.Transaction
}

func New(repositories *repositories.Repositories) (*Services, error) {
	userService, err := user.New(repositories.User)
	if err != nil {
		return nil, fmt.Errorf("Cannot create user service")
	}

	accountService, err := account.New(repositories.Account)
	if err != nil {
		return nil, fmt.Errorf("Cannot create account service")
	}

	transactionService, err := transaction.New(repositories.Transaction, repositories.Account, repositories.TransactionUow)
	if err != nil {
		return nil, fmt.Errorf("Cannot create account service")
	}

	return &Services{
		User:        userService,
		Account:     accountService,
		Transaction: transactionService,
	}, nil
}
