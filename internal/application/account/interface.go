package account

import accountDomain "atybank/internal/domain/account"

type Account interface {
	BelongToUser(userId, accountNumber string) (bool, error)
	Create(userId, name string, accountType string) error
	List(userId string) ([]accountDomain.AccountEntity, error)
	Get(accountNumber string) (accountDomain.AccountEntity, error)
	Delete(accountNumber string) (int64, error)
	Update(accountNumber, name, accountType string) (accountDomain.AccountEntity, error)
}
