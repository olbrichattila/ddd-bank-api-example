package account

import (
	"atybank/internal/domain/valueobjects"
	"time"

	"github.com/shopspring/decimal"
)

type AccountEntity interface {
	AccountNumber() AccountNumber
	UserId() valueobjects.UserId
	SortCode() string
	Name() string
	AccountType() AccountType
	Balance() decimal.Decimal
	Currency() Currency
	CreatedAt() time.Time
	UpdatedAt() time.Time

	// Setters
	SetName(name string) error
	SetAccountType(accountType AccountType) error
}
