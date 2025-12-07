package transaction

import (
	"time"

	"github.com/shopspring/decimal"
)

// TODO: Add value objects here as well, instead of primitives for the getters
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
