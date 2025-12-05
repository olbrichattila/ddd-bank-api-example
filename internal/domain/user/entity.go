// Package user entity + constructor + validator as aggregate root
package user

import (
	"eaglebank/internal/domain/shared/helpers"
	"fmt"
	"strings"
	"time"
)

type UserInput struct {
	Id          string
	Name        string
	Line1       string
	Line2       *string
	Line3       *string
	Town        string
	County      string
	Postcode    string
	PhoneNumber string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(input UserInput) (UserEntity, error) {
	if input.Id != "" && !helpers.IsValidUserId(input.Id) {
		return nil, fmt.Errorf("invalid user id")
	}

	if strings.TrimSpace(input.Name) == "" {
		return nil, fmt.Errorf("name is required")
	}

	// TODO add further domain level validation rules

	return &user{
		id:          input.Id,
		name:        input.Name,
		line1:       input.Line1,
		line2:       input.Line2,
		line3:       input.Line3,
		town:        input.Town,
		county:      input.County,
		postcode:    input.Postcode,
		phoneNumber: input.PhoneNumber,
		email:       input.Email,
		createdAt:   input.CreatedAt,
		updatedAt:   input.UpdatedAt,
	}, nil
}

type user struct {
	id          string
	name        string
	line1       string
	line2       *string
	line3       *string
	town        string
	county      string
	postcode    string
	phoneNumber string
	email       string
	createdAt   time.Time
	updatedAt   time.Time
}

func (u *user) Id() string           { return u.id }
func (u *user) Name() string         { return u.name }
func (u *user) Line1() string        { return u.line1 }
func (u *user) Line2() *string       { return u.line2 }
func (u *user) Line3() *string       { return u.line3 }
func (u *user) Town() string         { return u.town }
func (u *user) County() string       { return u.county }
func (u *user) Postcode() string     { return u.postcode }
func (u *user) PhoneNumber() string  { return u.phoneNumber }
func (u *user) Email() string        { return u.email }
func (u *user) CreatedAt() time.Time { return u.createdAt }
func (u *user) UpdatedAt() time.Time { return u.updatedAt }
