package services

import (
	"fmt"

	"eaglebank/cmd/eaglebank/repositories"
	"eaglebank/internal/application/account"
	"eaglebank/internal/application/transaction"
	"eaglebank/internal/application/user"
	"eaglebank/internal/shared/services"
	serviceImplementations "eaglebank/internal/shared/services/implementation"
)

type Services struct {
	User        user.User
	Account     account.Account
	Transaction transaction.Transaction
	Logger      services.Logger
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

	logger := serviceImplementations.NewScreenLogger()

	return &Services{
		User:        userService,
		Account:     accountService,
		Transaction: transactionService,
		Logger:      logger,
	}, nil
}
