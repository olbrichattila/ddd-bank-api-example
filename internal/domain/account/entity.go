package account

import (
	"fmt"
	"time"

	"eaglebank/internal/domain/shared/helpers"

	"github.com/shopspring/decimal"
)

type Input struct {
	AccountNumber string
	UserId        string
	SortCode      string
	Name          string
	AccountType   string
	Balance       decimal.Decimal
	Currency      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func New(input Input) (AccountEntity, error) {
	// Domain level validation
	if !helpers.IsValidAccountNumber(input.AccountNumber) {
		return nil, fmt.Errorf("invalid account number format")
	}

	return &entity{
		accountNumber: input.AccountNumber,
		userId:        input.UserId,
		sortCode:      input.SortCode,
		name:          input.Name,
		accountType:   input.AccountType,
		balance:       input.Balance,
		currency:      input.Currency,
		createdAt:     input.CreatedAt,
		updatedAt:     input.UpdatedAt,
	}, nil
}

type entity struct {
	accountNumber string
	userId        string
	sortCode      string
	name          string
	accountType   string
	balance       decimal.Decimal
	currency      string
	createdAt     time.Time
	updatedAt     time.Time
}

func (e *entity) AccountNumber() string {
	return e.accountNumber
}

func (e *entity) UserId() string {
	return e.userId
}

func (e *entity) AccountType() string {
	return e.accountType
}

func (e *entity) Balance() decimal.Decimal {
	return e.balance
}

func (e *entity) Currency() string {
	return e.currency
}

func (e *entity) Name() string {
	return e.name
}

func (e *entity) SortCode() string {
	return e.sortCode
}

func (e *entity) CreatedAt() time.Time {
	return e.createdAt
}

func (e *entity) UpdatedAt() time.Time {
	return e.updatedAt
}

// Setters

func (e *entity) SetAccountType(accountType string) error {
	if !helpers.IsValidAccountType(accountType) {
		return fmt.Errorf("invalid account type")
	}

	e.accountType = accountType

	return nil
}

func (e *entity) SetName(name string) error {
	if e.name == "" {
		return fmt.Errorf("account name is required")
	}
	e.name = name

	return nil
}
