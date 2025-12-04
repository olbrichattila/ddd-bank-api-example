package transaction

import (
	domain "eaglebank/internal/domain/transaction"

	"github.com/shopspring/decimal"
)

type Transaction interface {
	Create(amount decimal.Decimal, userId, currency, transactionType, accountNumber string, reference *string) error
	List(accountNumber string) ([]domain.TransactionEntity, error)
	Get(transactionNumber string) (domain.TransactionEntity, error)
}
