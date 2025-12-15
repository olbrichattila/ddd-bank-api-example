package repositories

import (
	"database/sql"
	"fmt"

	accountDomain "atybank/internal/domain/account"
	transactionDomain "atybank/internal/domain/transaction"
	userDomain "atybank/internal/domain/user"
	transactionUoWDomain "atybank/internal/infrastructure/workofunits"

	workOfUnits "atybank/internal/infrastructure/implementations/transactionwou"
	"atybank/internal/infrastructure/persistence/account"
	"atybank/internal/infrastructure/persistence/transaction"
	"atybank/internal/infrastructure/persistence/user"
)

type Repositories struct {
	User           userDomain.User
	Account        accountDomain.Account
	Transaction    transactionDomain.Transaction
	TransactionUow transactionUoWDomain.Transaction
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

	transactionWou, err := workOfUnits.New(db)
	if err != nil {
		return nil, fmt.Errorf("cannot create transaction work of unit %w", err)
	}

	return &Repositories{
		User:           userRepository,
		Account:        accountRepository,
		Transaction:    transactionRepository,
		TransactionUow: transactionWou,
	}, nil
}
