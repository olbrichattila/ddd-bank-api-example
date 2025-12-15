package transaction

import (
	"fmt"

	domain "atybank/internal/domain/transaction"
	"atybank/internal/infrastructure/dbexecutor"
	"atybank/internal/infrastructure/implementations/database"
)

func New(db dbexecutor.DbExecutor) (domain.Transaction, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil in transaction repository")
	}

	return &transaction{
		db: db,
	}, nil
}

type transaction struct {
	db dbexecutor.DbExecutor
}

func (a *transaction) BelongToUser(accountNumber, transactionId string) (bool, error) {
	sql := `
		SELECT
			COUNT(*) as cnt
		FROM transactions t
		WHERE
			t.account_number = $1
			AND id = $2`

	var accountCount int64
	ok, err := database.FetchOneRow(a.db, sql, []any{&accountCount}, accountNumber, transactionId)
	if err != nil {
		return false, fmt.Errorf("cannot verify if transaction belongs to the account %w", err)
	}

	if ok {
		return accountCount > 0, nil
	}

	return false, fmt.Errorf("Cannot retrieve data if the transaction belongs to the user")

}

func (t *transaction) Create(entity domain.TransactionEntity) error {
	// TODO this should be in transaction block, as it is critical, if fails, balance should not be updated
	sql := `INSERT INTO transactions (
		id, account_number, user_id, amount, currency, type, reference
	) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := database.ExecuteSQL(
		t.db,
		sql,
		entity.Id(),
		entity.AccountNumber(),
		entity.UserId(),
		entity.Amount(),
		entity.Currency(),
		entity.Type(),
		entity.Reference(),
	)

	if err != nil {
		return fmt.Errorf("account repository, error saving user: %w", err)
	}

	return nil
}

func (t *transaction) Get(transactionId string) (domain.TransactionEntity, error) {
	sql := `
	SELECT 
		id, account_number, user_id, amount, currency, type, reference, created_at
	FROM
		transactions
	WHERE id = $1`

	var input domain.Input
	ok, err := database.FetchOneRow(t.db, sql, []any{
		&input.Id,
		&input.AccountNumber,
		&input.UserId,
		&input.Amount,
		&input.Currency,
		&input.Type,
		&input.Reference,
		&input.CreatedAt,
	},
		transactionId,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot get transaction %w", err)
	}

	if ok {
		return domain.New(input)
	}

	return nil, nil
}

func (t *transaction) List(accountNumber string) ([]domain.TransactionEntity, error) {
	sql := `
		SELECT 
			id, account_number, user_id, amount, currency, type, reference, created_at 
		FROM
			transactions where account_number = $1`

	return database.FetchRowsAnMapToEntities(
		t.db,
		sql,
		func(input *domain.Input) []any {
			return []any{
				&input.Id,
				&input.AccountNumber,
				&input.UserId,
				&input.Amount,
				&input.Currency,
				&input.Type,
				&input.Reference,
				&input.CreatedAt,
			}
		},
		func(input *domain.Input) (domain.TransactionEntity, error) {
			entity, err := domain.New(*input)
			if err != nil {
				return nil, err
			}
			return entity, nil
		},
		accountNumber,
	)
}
