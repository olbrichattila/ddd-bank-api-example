package user

import (
	userDomain "eaglebank/internal/domain/user"
)

type User interface {
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
	Delete(userID string) (int64, error)
}
