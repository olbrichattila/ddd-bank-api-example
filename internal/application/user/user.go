// Package user represents user application level abstraction
package user

import (
	"fmt"

	"eaglebank/internal/domain/shared/helpers"
	userDomain "eaglebank/internal/domain/user"
)

func New(userRepo userDomain.User) (User, error) {
	if userRepo == nil {
		return nil, fmt.Errorf("nil value passed to user repository new function")
	}
	return &userService{
		userRepo: userRepo,
	}, nil
}

type userService struct {
	userRepo userDomain.User
}

func (u *userService) GetByEmail(email string) (userDomain.UserEntity, error) {
	return u.userRepo.GetByEmail(email)
}

func (u *userService) Create(
	name string,
	line1 string,
	line2 *string,
	line3 *string,
	town string,
	county string,
	postcode string,
	phoneNumber string,
	email string,
) (string, error) {
	existingUserEntity, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", fmt.Errorf("could not create new user id %w", err)
	}

	if existingUserEntity != nil {
		return "", fmt.Errorf("User already exists")
	}

	newUserId, err := helpers.GenerateNewUserId()
	if err != nil {
		return "", fmt.Errorf("could not create new user id %w", err)
	}

	userEntity, err := userDomain.New(
		userDomain.UserInput{
			Id:          newUserId,
			Name:        name,
			Line1:       line1,
			Line2:       line2,
			Line3:       line3,
			Town:        town,
			County:      county,
			Postcode:    postcode,
			PhoneNumber: phoneNumber,
			Email:       email,
		},
	)

	if err != nil {
		return "", err
	}

	return u.userRepo.Create(userEntity)
}

func (u *userService) Get(userId string) (userDomain.UserEntity, error) {
	return u.userRepo.Get(userId)
}

func (u *userService) Update(
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
) (userDomain.UserEntity, error) {

	userEntity, err := userDomain.New(
		userDomain.UserInput{
			Id:          id,
			Name:        name,
			Line1:       line1,
			Line2:       line2,
			Line3:       line3,
			Town:        town,
			County:      county,
			Postcode:    postcode,
			PhoneNumber: phoneNumber,
			Email:       email,
		},
	)

	if err != nil {
		return nil, err
	}

	err = u.userRepo.Update(userEntity)
	if err != nil {
		return nil, err
	}

	return u.userRepo.Get(id)
}

func (u *userService) Delete(userId string) (int64, error) {
	return u.userRepo.Delete(userId)
}
