package account

import (
	"fmt"

	accountDomain "eaglebank/internal/domain/account"

	"github.com/shopspring/decimal"
)

const (
	eagleSortCode   = "10-10-10"
	defaultCurrency = "GBP"
)

func New(accountRepository accountDomain.Account) (Account, error) {
	if accountRepository == nil {
		return nil, fmt.Errorf("missing account repository from new account service")
	}

	return &account{
		accountRepository: accountRepository,
	}, nil
}

type account struct {
	accountRepository accountDomain.Account
}

func (a *account) BelongToUser(userId, accountNumber string) (bool, error) {
	return a.accountRepository.BelongToUser(userId, accountNumber)
}

func (a *account) Create(userID, name string, accountType string) error {
	accountNumber, err := a.accountRepository.NextAccountNumber()
	if err != nil {
		return fmt.Errorf("account number creation error: %w", err)
	}

	accountEntity, err := accountDomain.New(
		accountDomain.Input{
			AccountNumber: accountNumber,
			UserID:        userID,
			SortCode:      eagleSortCode,
			Name:          name,
			AccountType:   accountType,
			Balance:       decimal.New(0, 0), // initial balance
			Currency:      defaultCurrency,
		},
	)

	if err != nil {
		return fmt.Errorf("account creation error: %w", err)
	}

	return a.accountRepository.Create(accountEntity)
}

func (a *account) List(userID string) ([]accountDomain.AccountEntity, error) {
	return a.accountRepository.List(userID)
}

func (a *account) Get(accountNumber string) (accountDomain.AccountEntity, error) {
	return a.accountRepository.Get(accountNumber)
}

func (a *account) Delete(accountNumber string) (int64, error) {
	return a.accountRepository.Delete(accountNumber)
}

func (a *account) Update(accountNumber string, name string, accountType string) (accountDomain.AccountEntity, error) {
	accountEntity, err := a.accountRepository.Get(accountNumber)
	if err != nil {
		return nil, err
	}

	if accountEntity == nil {
		return nil, nil
	}

	err = accountEntity.SetName(name)
	if err != nil {
		return nil, err
	}

	err = accountEntity.SetAccountType(accountType)
	if err != nil {
		return nil, err
	}

	err = a.accountRepository.Update(accountEntity)
	if err != nil {
		return nil, err
	}

	return accountEntity, nil
}
