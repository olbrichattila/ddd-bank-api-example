package workofunits

import (
	"eaglebank/internal/domain/transaction"

	"github.com/shopspring/decimal"
)

//go:generate mockgen -destination=mock/transaction-wou-mock.go -package=mock . Transaction
type Transaction interface {
	Create(
		transactionEntity transaction.TransactionEntity,
		accountNumber string,
		correctedAmount decimal.Decimal,
	) error
}
