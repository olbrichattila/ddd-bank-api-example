// This work of unit is to pull the transaction insertion and the balance update
// into an ATOMIC transaction, as this is a critical operation, all or nothing.
// This is a challenge in Domain Driven design, as the tradeoff is this unit dependency is not
// loosely coupled
package workofunits

import (
	"database/sql"
	"fmt"

	"atybank/internal/domain/transaction"
	accountRepository "atybank/internal/infrastructure/persistence/account"
	transactionRepository "atybank/internal/infrastructure/persistence/transaction"
	domainWou "atybank/internal/infrastructure/workofunits"

	"github.com/shopspring/decimal"
)

func New(
	db *sql.DB,

) (domainWou.Transaction, error) {
	if db == nil {
		return nil, fmt.Errorf("transaction wou received nil *sql.DB")
	}

	return &transactionWou{
		db: db,
	}, nil
}

type transactionWou struct {
	db *sql.DB
}

func (t *transactionWou) Create(
	transactionEntity transaction.TransactionEntity,
	accountNumber string,
	correctedAmount decimal.Decimal,
) (err error) {
	tx, err := t.db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				err = txErr
			}
			return
		}
		txErr := tx.Commit()
		if txErr != nil {
			err = txErr
		}
	}()

	accountRepository, err := accountRepository.New(tx)
	transactionRepository, err := transactionRepository.New(tx)

	accountRepository.UpdateBalance(accountNumber, correctedAmount)
	err = transactionRepository.Create(transactionEntity)

	return
}
