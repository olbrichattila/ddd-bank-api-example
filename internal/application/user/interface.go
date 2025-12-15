package user

import (
	userDomain "atybank/internal/domain/user"
)

type User interface {
	GetByEmail(email string) (userDomain.UserEntity, error)
	Create(
		name string,
		line1 string,
		line2 *string,
		line3 *string,
		town string,
		county string,
		postcode string,
		phoneNumber string,
		email string,
	) (string, error)

	Get(userId string) (userDomain.UserEntity, error)
	Update(
		id string,
		name string,
		line1 string,
		line2 *string,
		line3 *string,
		town string,
		county string,
		postcode string,
		phoneNumber string,
		email string,
	) (userDomain.UserEntity, error)
	Delete(userId string) (int64, error)
}
