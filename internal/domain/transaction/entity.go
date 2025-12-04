package transaction

import (
	"eaglebank/internal/domain/shared/helpers"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	minAllowedAmount = 0
	maxAllowedAmount = 10000
)

type Input struct {
	Id            string
	AccountNumber string
	UserId        string
	Amount        decimal.Decimal
	Currency      string
	Type          string
	Reference     *string
	CreatedAt     time.Time
}

func New(input Input) (TransactionEntity, error) {
	if input.Id != "" && helpers.IsValidTransactionID(input.Id) {
		return nil, fmt.Errorf("invalid transaction ID")
	}

	if !helpers.IsValidUserID(input.UserId) {
		return nil, fmt.Errorf("invalid user ID")
	}

	if !helpers.IsValidAccountNumber(input.AccountNumber) {
		return nil, fmt.Errorf("invalid account number")
	}

	if input.Amount.LessThan(decimal.New(minAllowedAmount, 0)) || input.Amount.GreaterThan(decimal.New(maxAllowedAmount, 0)) {
		return nil, fmt.Errorf("amount cannot must be between %d and %d", minAllowedAmount, maxAllowedAmount)
	}

	return &entity{
		id:              input.Id,
		accountNumber:   input.AccountNumber,
		userId:          input.UserId,
		amount:          input.Amount,
		currency:        input.Currency,
		transactionType: input.Type,
		reference:       input.Reference,
		createdAt:       input.CreatedAt,
	}, nil
}

type entity struct {
	id              string
	accountNumber   string
	userId          string
	amount          decimal.Decimal
	currency        string
	transactionType string
	reference       *string
	createdAt       time.Time
}

func (e *entity) Id() string {
	return e.id
}

func (e *entity) UserId() string {
	return e.userId
}

func (e *entity) AccountNumber() string {
	return e.accountNumber
}

func (e *entity) Amount() decimal.Decimal {
	return e.amount
}

func (e *entity) Currency() string {
	return e.currency
}

func (e *entity) Reference() *string {
	return e.reference
}

func (e *entity) Type() string {
	return e.transactionType
}

func (e *entity) CreatedAt() time.Time {
	return e.createdAt
}
