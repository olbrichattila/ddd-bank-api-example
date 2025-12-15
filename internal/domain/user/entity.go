// Package user entity + constructor + validator as aggregate root
package user

import (
	"fmt"
	"strings"
	"time"

	"atybank/internal/domain/valueobjects"
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
	userId, err := valueobjects.NewUserId(input.Id)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(input.Name) == "" {
		return nil, fmt.Errorf("name is required")
	}

	email, err := NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	phone, err := NewPhoneNumber(input.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &user{
		id:          userId,
		name:        input.Name,
		line1:       input.Line1,
		line2:       input.Line2,
		line3:       input.Line3,
		town:        input.Town,
		county:      input.County,
		postcode:    input.Postcode,
		phoneNumber: Phone(phone),
		email:       email,
		createdAt:   input.CreatedAt,
		updatedAt:   input.UpdatedAt,
	}, nil
}

type user struct {
	id          valueobjects.UserId
	name        string
	line1       string
	line2       *string
	line3       *string
	town        string
	county      string
	postcode    string
	phoneNumber Phone
	email       Email
	createdAt   time.Time
	updatedAt   time.Time
}

func (u *user) Id() valueobjects.UserId { return u.id }
func (u *user) Name() string            { return u.name }
func (u *user) Line1() string           { return u.line1 }
func (u *user) Line2() *string          { return u.line2 }
func (u *user) Line3() *string          { return u.line3 }
func (u *user) Town() string            { return u.town }
func (u *user) County() string          { return u.county }
func (u *user) Postcode() string        { return u.postcode }
func (u *user) PhoneNumber() Phone      { return u.phoneNumber }
func (u *user) Email() Email            { return u.email }
func (u *user) CreatedAt() time.Time    { return u.createdAt }
func (u *user) UpdatedAt() time.Time    { return u.updatedAt }
