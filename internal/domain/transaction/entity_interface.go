package transaction

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionEntity interface {
	Id() string
	AccountNumber() string
	UserId() string
	Amount() decimal.Decimal
	Currency() string
	Type() string
	Reference() *string
	CreatedAt() time.Time
}
