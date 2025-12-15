package account

import (
	"fmt"
	"time"

	"atybank/internal/domain/valueobjects"

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
	accountNumber, err := NewAccountNumber(input.AccountNumber)
	if err != nil {
		return nil, fmt.Errorf("creating account, invalid account number %w", err)
	}

	userId, err := valueobjects.NewUserId(input.UserId)
	if err != nil {
		return nil, fmt.Errorf("creating account, invalid user id number %w", err)
	}

	accountType, err := NewAccountType(input.AccountType)
	if err != nil {
		return nil, fmt.Errorf("creating account, invalid account type %w", err)
	}

	currency, err := NewCurrency(input.Currency)
	if err != nil {
		return nil, fmt.Errorf("creating account, invalid currency %w", err)
	}

	return &entity{
		accountNumber: accountNumber,
		userId:        userId,
		sortCode:      input.SortCode,
		name:          input.Name,
		accountType:   accountType,
		balance:       input.Balance,
		currency:      currency,
		createdAt:     input.CreatedAt,
		updatedAt:     input.UpdatedAt,
	}, nil
}

type entity struct {
	accountNumber AccountNumber
	userId        valueobjects.UserId
	sortCode      string
	name          string
	accountType   AccountType
	balance       decimal.Decimal
	currency      Currency
	createdAt     time.Time
	updatedAt     time.Time
}

func (e *entity) AccountNumber() AccountNumber {
	return e.accountNumber
}

func (e *entity) UserId() valueobjects.UserId {
	return e.userId
}

func (e *entity) AccountType() AccountType {
	return e.accountType
}

func (e *entity) Balance() decimal.Decimal {
	return e.balance
}

func (e *entity) Currency() Currency {
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
func (e *entity) SetAccountType(accountType AccountType) error {
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
