package account

import (
	"time"

	"github.com/shopspring/decimal"
)

type AccountEntity interface {
	AccountNumber() string
	UserId() string
	SortCode() string
	Name() string
	AccountType() string
	Balance() decimal.Decimal
	Currency() string
	CreatedAt() time.Time
	UpdatedAt() time.Time

	// Setters
	SetName(name string) error
	SetAccountType(accountType string) error
}
