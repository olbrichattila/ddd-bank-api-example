package repositories

import (
	"database/sql"
	"fmt"

	accountDomain "eaglebank/internal/domain/account"
	transactionDomain "eaglebank/internal/domain/transaction"
	userDomain "eaglebank/internal/domain/user"

	"eaglebank/internal/infrastructure/persistence/account"
	"eaglebank/internal/infrastructure/persistence/transaction"
	"eaglebank/internal/infrastructure/persistence/user"
)

type Repositories struct {
	User        userDomain.User
	Account     accountDomain.Account
	Transaction transactionDomain.Transaction
}

func New(db *sql.DB) (*Repositories, error) {
	userRepository, err := user.New(db)
	if err != nil {
		return nil, fmt.Errorf("cannot create user repository %w", err)
	}

	accountRepository, err := account.New(db)
	if err != nil {
		return nil, fmt.Errorf("cannot create account repository %w", err)
	}

	transactionRepository, err := transaction.New(db)
	if err != nil {
		return nil, fmt.Errorf("cannot create transaction repository %w", err)
	}

	return &Repositories{
		User:        userRepository,
		Account:     accountRepository,
		Transaction: transactionRepository,
	}, nil
}
