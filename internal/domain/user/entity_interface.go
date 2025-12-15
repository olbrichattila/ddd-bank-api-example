package user

import (
	"atybank/internal/domain/valueobjects"
	"time"
)

type UserEntity interface {
	Id() valueobjects.UserId
	Name() string
	Line1() string
	Line2() *string
	Line3() *string
	Town() string
	County() string
	Postcode() string
	PhoneNumber() Phone
	Email() Email
	CreatedAt() time.Time
	UpdatedAt() time.Time
}
