package user

import "time"

type UserEntity interface {
	ID() string
	Name() string
	Line1() string
	Line2() *string
	Line3() *string
	Town() string
	County() string
	Postcode() string
	PhoneNumber() string
	Email() string
	CreatedAt() time.Time
	UpdatedAt() time.Time
}
