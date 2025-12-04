package transaction

import (
	"fmt"

	"github.com/shopspring/decimal"

	accountDomain "eaglebank/internal/domain/account"
	"eaglebank/internal/domain/shared/helpers"
	domain "eaglebank/internal/domain/transaction"
	transactionDomain "eaglebank/internal/domain/transaction"
)

func New(
	transactionRepository transactionDomain.Transaction,
	accountRepository accountDomain.Account,
) (Transaction, error) {
	if transactionRepository == nil {
		return nil, fmt.Errorf("missing transaction repository from new transaction service")
	}

	if accountRepository == nil {
		return nil, fmt.Errorf("missing account repository from new transaction service")
	}

	return &transaction{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}, nil
}

type transaction struct {
	transactionRepository transactionDomain.Transaction
	accountRepository     accountDomain.Account
}

func (t *transaction) Create(amount decimal.Decimal, userId, currency, transactionType, accountNumber string, reference *string) error {
	var correctedAmount decimal.Decimal
	switch transactionType {
	case transactionDomain.TransferTypeDeposit:
		correctedAmount = amount
	case transactionDomain.TransferTypeWithdrawal:
		correctedAmount = amount.Neg()
	default:
		return fmt.Errorf("incorrect transaction type %s", transactionType)
	}

	isNegative, err := t.willBalanceBeNegative(accountNumber, correctedAmount)
	if err != nil {
		return err
	}

	if isNegative {
		return fmt.Errorf("Balance cannot be negative")
	}

	id := helpers.GenerateTransactionID()
	entity, err := transactionDomain.New(
		transactionDomain.Input{
			Id:            id,
			AccountNumber: accountNumber,
			UserId:        userId,
			Amount:        amount,
			Currency:      currency,
			Type:          transactionType,
			Reference:     reference,
		},
	)

	if err != nil {
		return err
	}

	// TODO: update rolling balance, ideally this should be in an ATOMIC transaction, rollback
	// should create UoW Unit of Work pattern instead
	t.accountRepository.UpdateBalance(accountNumber, correctedAmount)

	return t.transactionRepository.Create(entity)
}

func (t *transaction) Get(transactionNumber string) (domain.TransactionEntity, error) {
	return t.transactionRepository.Get(transactionNumber)
}

func (t *transaction) List(accountNumber string) ([]domain.TransactionEntity, error) {
	return t.transactionRepository.List(accountNumber)
}

func (t *transaction) willBalanceBeNegative(accountNumber string, correctedAmount decimal.Decimal) (bool, error) {
	accountEntity, err := t.accountRepository.Get(accountNumber)
	if err != nil {
		return false, err
	}

	newBalance := accountEntity.Balance().Sub(correctedAmount)
	return newBalance.IsNegative(), nil
}
