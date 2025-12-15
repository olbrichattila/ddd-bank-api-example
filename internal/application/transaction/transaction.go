package transaction

import (
	"fmt"

	"github.com/shopspring/decimal"

	accountDomain "atybank/internal/domain/account"
	domain "atybank/internal/domain/transaction"
	transactionDomain "atybank/internal/domain/transaction"
	transactionUoWDomain "atybank/internal/infrastructure/workofunits"
	"atybank/internal/shared/helpers"
)

func New(
	transactionRepository transactionDomain.Transaction,
	accountRepository accountDomain.Account,
	transactionWouService transactionUoWDomain.Transaction,

) (Transaction, error) {
	if transactionRepository == nil {
		return nil, fmt.Errorf("missing transaction repository from new transaction service")
	}

	if accountRepository == nil {
		return nil, fmt.Errorf("missing account repository from new transaction service")
	}

	if transactionWouService == nil {
		return nil, fmt.Errorf("missing transactionWouService from new transaction service")
	}

	return &transaction{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
		transactionWouService: transactionWouService,
	}, nil
}

type transaction struct {
	transactionRepository transactionDomain.Transaction
	accountRepository     accountDomain.Account
	transactionWouService transactionUoWDomain.Transaction
}

func (t *transaction) BelongToUser(accountNumber, transactionId string) (bool, error) {
	return t.transactionRepository.BelongToUser(accountNumber, transactionId)
}

func (t *transaction) Create(amount decimal.Decimal, userId, currency, transactionType, accountNumber string, reference *string) error {
	correctedAmount, err := t.getCorrectedAmount(amount, transactionType)
	if err != nil {
		return nil
	}

	isNegative, err := t.willBalanceBeNegative(accountNumber, correctedAmount)
	if err != nil {
		return err
	}

	if isNegative {
		return fmt.Errorf("Balance cannot be negative")
	}

	id := helpers.GenerateTransactionId()
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

	// This Wou puts the two repositories into a same atomic transaction
	return t.transactionWouService.Create(entity, accountNumber, correctedAmount)
}

func (t *transaction) Get(transactionNumber string) (domain.TransactionEntity, error) {
	return t.transactionRepository.Get(transactionNumber)
}

func (t *transaction) List(accountNumber string) ([]domain.TransactionEntity, error) {
	return t.transactionRepository.List(accountNumber)
}

func (t *transaction) getCorrectedAmount(amount decimal.Decimal, transactionType string) (decimal.Decimal, error) {
	switch transactionType {
	case transactionDomain.TransferTypeDeposit:
		return amount, nil
	case transactionDomain.TransferTypeWithdrawal:
		return amount.Neg(), nil
	default:
		return amount, fmt.Errorf("incorrect transaction type %s", transactionType)
	}
}

func (t *transaction) willBalanceBeNegative(accountNumber string, correctedAmount decimal.Decimal) (bool, error) {
	accountEntity, err := t.accountRepository.Get(accountNumber)
	if err != nil {
		return false, err
	}

	newBalance := accountEntity.Balance().Add(correctedAmount)
	return newBalance.IsNegative(), nil
}
