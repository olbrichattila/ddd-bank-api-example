package account

import (
	"database/sql"
	"fmt"

	domain "eaglebank/internal/domain/account"

	"github.com/shopspring/decimal"
)

func New(db *sql.DB) (domain.Account, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil in account repository")
	}
	return &account{
		db: db,
	}, nil
}

type account struct {
	db *sql.DB
}

func (a *account) NextAccountNumber() (string, error) {
	sql := "SELECT nextval('account_number_seq')"

	rows, err := a.db.Query(sql)
	if err != nil {
		return "", fmt.Errorf("cannot get account number %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var accountNumber int64
		err := rows.Scan(&accountNumber)

		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%08d", accountNumber), nil
	}

	if rows.Err() != nil {
		return "", fmt.Errorf("sql fetch error %w", err)
	}

	return "", fmt.Errorf("cannot generate account number ")
}

func (a *account) BelongToUser(userId, accountNumber string) (bool, error) {
	sql := `SELECT COUNT(*) AS account_count
	FROM
		accounts
	WHERE
		account_number = $1
		AND user_id = $2`

	rows, err := a.db.Query(sql, accountNumber, userId)
	if err != nil {
		return false, fmt.Errorf("cannot verify if account belongs to user %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var accountCount int64
		err := rows.Scan(&accountCount)

		if err != nil {
			return false, err
		}

		return accountCount > 0, nil
	}

	return false, rows.Err()
}

func (a *account) Create(entity domain.AccountEntity) error {
	sql := `INSERT INTO accounts (
		account_number,
		user_id,
		sort_code,
		name,
		account_type,
		balance,
		currency
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := a.db.Exec(
		sql,
		entity.AccountNumber(),
		entity.UserId(),
		entity.SortCode(),
		entity.Name(),
		entity.AccountType(),
		entity.Balance().String(),
		entity.Currency(),
	)

	if err != nil {
		return fmt.Errorf("account repository, error saving user: %w", err)
	}

	return nil
}

func (a *account) Delete(accountNumber string) (int64, error) {
	sql := `DELETE FROM accounts WHERE account_number = $1`

	result, err := a.db.Exec(
		sql,
		accountNumber,
	)

	if err != nil {
		return 0, fmt.Errorf("user repository, error deleting user: %w", err)
	}

	return result.RowsAffected()
}

func (a *account) Get(accountNumber string) (domain.AccountEntity, error) {
	sql := `SELECT 
		account_number,
		user_id,
		sort_code,
		name,
		account_type,
		balance,
		currency,
		created_at,
		updated_at 
	FROM
		accounts where account_number = $1`
	rows, err := a.db.Query(sql, accountNumber)
	if err != nil {
		return nil, fmt.Errorf("query execution error %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var input domain.Input
		err := rows.Scan(
			&input.AccountNumber,
			&input.UserId,
			&input.SortCode,
			&input.Name,
			&input.AccountType,
			&input.Balance,
			&input.Currency,
			&input.CreatedAt,
			&input.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		return domain.New(input)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("sql fetch error %w", err)
	}

	return nil, nil
}

func (a *account) List(userId string) ([]domain.AccountEntity, error) {
	var accountEntities []domain.AccountEntity
	sql := `SELECT 
		account_number,
		user_id,
		sort_code,
		name,
		account_type,
		balance,
		currency,
		created_at,
		updated_at 
	FROM
		accounts where user_id = $1`
	rows, err := a.db.Query(sql, userId)
	if err != nil {
		return nil, fmt.Errorf("query execution error %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var input domain.Input
		err := rows.Scan(
			&input.AccountNumber,
			&input.UserId,
			&input.SortCode,
			&input.Name,
			&input.AccountType,
			&input.Balance,
			&input.Currency,
			&input.CreatedAt,
			&input.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		accountEntity, err := domain.New(input)
		if err != nil {
			return nil, err
		}

		accountEntities = append(accountEntities, accountEntity)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("sql fetch error %w", err)
	}

	return accountEntities, nil
}

func (a *account) Update(entity domain.AccountEntity) error {
	sql := `UPDATE accounts SET
		name = $1,
		account_type = $2
	WHERE 
		account_number = $3`

	_, err := a.db.Exec(
		sql,
		entity.Name(),
		entity.AccountType(),
		entity.AccountNumber(),
	)

	if err != nil {
		return fmt.Errorf("account repository, error updating account: %w", err)
	}

	return nil
}

func (a *account) UpdateBalance(accountNumber string, amount decimal.Decimal) error {
	// this should be in db transaction
	sql := `UPDATE accounts
	SET
		balance = balance + $1
	WHERE 
		account_number = $2`

	_, err := a.db.Exec(
		sql,
		amount.StringFixed(2),
		accountNumber,
	)

	if err != nil {
		return fmt.Errorf("account repository, error updating balance: %w", err)
	}

	return nil
}
