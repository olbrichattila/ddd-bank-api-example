package account

import "github.com/shopspring/decimal"

type Account interface {
	Create(entity AccountEntity) error
	List(userID string) ([]AccountEntity, error)
	Get(accountNumber string) (AccountEntity, error)
	Update(entity AccountEntity) error
	Delete(accountNumber string) (int64, error)
	NextAccountNumber() (string, error)
	UpdateBalance(accountNumber string, amount decimal.Decimal) error
}
