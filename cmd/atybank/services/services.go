package services

import (
	"fmt"

	"atybank/cmd/atybank/repositories"
	"atybank/internal/application/account"
	"atybank/internal/application/transaction"
	"atybank/internal/application/user"
	"atybank/internal/shared/services"
	serviceImplementations "atybank/internal/shared/services/implementation"
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
