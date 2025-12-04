package account

import accountDomain "eaglebank/internal/domain/account"

type Account interface {
	Create(userID, name string, accountType string) error
	List(userID string) ([]accountDomain.AccountEntity, error)
	Get(accountNumber string) (accountDomain.AccountEntity, error)
	Delete(accountNumber string) (int64, error)
	Update(accountNumber, name, accountType string) (accountDomain.AccountEntity, error)
}
