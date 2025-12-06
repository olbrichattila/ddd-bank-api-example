// Package user
package user

import (
	"database/sql"
	"fmt"

	domain "eaglebank/internal/domain/user"
	"eaglebank/internal/infrastructure/implementations/database"
)

func New(db *sql.DB) (domain.User, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil in user repository")
	}
	return &user{
		db: db,
	}, nil
}

type user struct {
	db *sql.DB
}

func (u *user) Create(entity domain.UserEntity) (string, error) {
	sql := `INSERT INTO users (id, name, line1, line2, line3, town, county, postcode, phone_number, email)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := database.ExecuteSQL(
		u.db,
		sql,
		entity.Id(),
		entity.Name(),
		entity.Line1(),
		entity.Line2(),
		entity.Line3(),
		entity.Town(),
		entity.County(),
		entity.Postcode(),
		entity.PhoneNumber(),
		entity.Email(),
	)

	if err != nil {
		return "", fmt.Errorf("user repository, error saving user: %w", err)
	}

	return entity.Id(), nil
}

func (u *user) GetByEmail(email string) (domain.UserEntity, error) {
	return u.getByIdOrEmail(false, email)
}

func (u *user) Get(userId string) (domain.UserEntity, error) {
	return u.getByIdOrEmail(true, userId)
}

func (u *user) getByIdOrEmail(isId bool, value string) (domain.UserEntity, error) {
	sql := `SELECT 
		id, name, line1, line2, line3, town, county, postcode, phone_number, email, created_at, updated_at 
	FROM
		users
	WHERE`

	if isId {
		sql += ` id = $1`
	} else {
		sql += ` email = $1`
	}

	var input domain.UserInput
	ok, err := database.FetchOneRow(
		u.db,
		sql,
		[]any{
			&input.Id,
			&input.Name,
			&input.Line1,
			&input.Line2,
			&input.Line3,
			&input.Town,
			&input.County,
			&input.Postcode,
			&input.PhoneNumber,
			&input.Email,
			&input.CreatedAt,
			&input.UpdatedAt,
		},
		value,
	)
	if err != nil {
		return nil, fmt.Errorf("query execution error %w", err)
	}

	if ok {
		return domain.New(input)
	}

	return nil, nil
}

func (u *user) Update(entity domain.UserEntity) error {
	sql := `UPDATE users 
	SET name = $1,
		line1 = $2,
		line2 = $3,
		line3 =$4,
		town = $5,
		county =$6,
		postcode =$7,
		phone_number = $8,
		email = $9
	WHERE id = $10`

	_, err := database.ExecuteSQL(
		u.db,
		sql,
		entity.Name(),
		entity.Line1(),
		entity.Line2(),
		entity.Line3(),
		entity.Town(),
		entity.County(),
		entity.Postcode(),
		entity.PhoneNumber(),
		entity.Email(),
		entity.Id(),
	)

	if err != nil {
		return fmt.Errorf("user repository, error updating user: %w", err)
	}

	return nil
}

func (u *user) Delete(userId string) (int64, error) {
	sql := `DELETE FROM users WHERE id = $1`

	result, err := database.ExecuteSQL(
		u.db,
		sql,
		userId,
	)

	if err != nil {
		return 0, fmt.Errorf("user repository, error deleting user: %w", err)
	}

	return result.RowsAffected()
}
