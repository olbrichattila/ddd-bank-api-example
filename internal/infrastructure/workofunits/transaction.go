package workofunits

import (
	"eaglebank/internal/domain/transaction"

	"github.com/shopspring/decimal"
)

type Transaction interface {
	Create(
		transactionEntity transaction.TransactionEntity,
		accountNumber string,
		correctedAmount decimal.Decimal,
	) error
}
