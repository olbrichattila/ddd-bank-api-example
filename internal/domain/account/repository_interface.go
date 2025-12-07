package account

import "github.com/shopspring/decimal"

//go:generate mockgen -destination=../../infrastructure/persistence/account/mock/account-mock.go -package=mock . Account
type Account interface {
	BelongToUser(userId, accountNumber string) (bool, error)
	Create(entity AccountEntity) error
	List(userId string) ([]AccountEntity, error)
	Get(accountNumber string) (AccountEntity, error)
	Update(entity AccountEntity) error
	Delete(accountNumber string) (int64, error)
	NextAccountNumber() (string, error)
	UpdateBalance(accountNumber string, amount decimal.Decimal) error
}
