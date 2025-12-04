package transaction

import (
	"database/sql"
	"fmt"

	domain "eaglebank/internal/domain/transaction"
)

func New(db *sql.DB) (domain.Transaction, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil in transaction repository")
	}

	return &transaction{
		db: db,
	}, nil
}

type transaction struct {
	db *sql.DB
}

func (t *transaction) Create(entity domain.TransactionEntity) error {
	// TODO this should be in transaction block, as it is critical, if fails, balance should not be updated
	sql := `INSERT INTO transactions (
		id, account_number, user_id, amount, currency, type, reference
	) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := t.db.Exec(
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
	sql := `SELECT 
		id, account_number, user_id, amount, currency, type, reference, created_at
	FROM
		transactions
	WHERE id = $1`

	rows, err := t.db.Query(sql, transactionId)
	if err != nil {
		return nil, fmt.Errorf("query execution error %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var input domain.Input
		err := rows.Scan(
			&input.Id,
			&input.AccountNumber,
			&input.UserId,
			&input.Amount,
			&input.Currency,
			&input.Type,
			&input.Reference,
			&input.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		return domain.New(input)
	}

	return nil, nil
}

func (t *transaction) List(accountNumber string) ([]domain.TransactionEntity, error) {
	sql := `SELECT 
			id, account_number, user_id, amount, currency, type, reference, created_at 
		FROM
			transactions where account_number = $1`

	rows, err := t.db.Query(sql, accountNumber)
	if err != nil {
		return nil, fmt.Errorf("query execution error %w", err)
	}
	defer rows.Close()

	var transactions []domain.TransactionEntity

	for rows.Next() {
		var input domain.Input
		err := rows.Scan(
			&input.Id,
			&input.AccountNumber,
			&input.UserId,
			&input.Amount,
			&input.Currency,
			&input.Type,
			&input.Reference,
			&input.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		transaction, err := domain.New(input)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)

	}

	return transactions, nil
}
