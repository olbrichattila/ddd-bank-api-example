package account

import (
	"database/sql"
	"fmt"

	domain "eaglebank/internal/domain/account"
	"eaglebank/internal/infrastructure/implementations/database"

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
	sql := "SELECT nextval('account_number_seq') as next_account_number"

	var accountNumber int64
	ok, err := database.FetchOneRow(a.db, sql, []any{&accountNumber})
	if err != nil {
		return "", fmt.Errorf("cannot generate account number ")
	}

	if ok {
		return fmt.Sprintf("%08d", accountNumber), nil
	}

	return "", fmt.Errorf("cannot generate account number ")
}

func (a *account) BelongToUser(userId, accountNumber string) (bool, error) {
	sql := `
	SELECT
		COUNT(*) AS account_count
	FROM
		accounts
	WHERE
		account_number = $1
		AND user_id = $2`

	var accountCount int64
	ok, err := database.FetchOneRow(a.db, sql, []any{&accountCount}, accountNumber, userId)
	if err != nil {
		return false, fmt.Errorf("error fetching account count %w", err)
	}

	if ok {
		return accountCount > 0, nil
	}

	return false, fmt.Errorf("cannot fetch account count")
}

func (a *account) Create(entity domain.AccountEntity) error {
	sql := `
	INSERT INTO accounts (
		account_number,
		user_id,
		sort_code,
		name,
		account_type,
		balance,
		currency
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := database.ExecuteSQL(
		a.db,
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

	result, err := database.ExecuteSQL(
		a.db,
		sql,
		accountNumber,
	)

	if err != nil {
		return 0, fmt.Errorf("user repository, error deleting user: %w", err)
	}

	return result.RowsAffected()
}

func (a *account) Get(accountNumber string) (domain.AccountEntity, error) {
	sql := `
	SELECT 
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

	var input domain.Input
	ok, err := database.FetchOneRow(
		a.db,
		sql,
		[]any{
			&input.AccountNumber,
			&input.UserId,
			&input.SortCode,
			&input.Name,
			&input.AccountType,
			&input.Balance,
			&input.Currency,
			&input.CreatedAt,
			&input.UpdatedAt,
		},
		accountNumber,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot fetch account %w", err)
	}

	if ok {
		return domain.New(input)
	}

	return nil, nil
}

func (a *account) List(userId string) ([]domain.AccountEntity, error) {
	sql := `
	SELECT 
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

	return database.FetchRowsAnMapToEntities(
		a.db,
		sql,
		func(input *domain.Input) []any {
			return []any{
				&input.AccountNumber,
				&input.UserId,
				&input.SortCode,
				&input.Name,
				&input.AccountType,
				&input.Balance,
				&input.Currency,
				&input.CreatedAt,
				&input.UpdatedAt,
			}
		},
		func(input *domain.Input) (domain.AccountEntity, error) {
			entity, err := domain.New(*input)
			if err != nil {
				return nil, err
			}
			return entity, nil
		},
		userId,
	)
}

func (a *account) Update(entity domain.AccountEntity) error {
	sql := `UPDATE accounts SET
		name = $1,
		account_type = $2
	WHERE 
		account_number = $3`

	_, err := database.ExecuteSQL(
		a.db,
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

	_, err := database.ExecuteSQL(
		a.db,
		sql,
		amount.StringFixed(2),
		accountNumber,
	)

	if err != nil {
		return fmt.Errorf("account repository, error updating balance: %w", err)
	}

	return nil
}
