package user

import (
	"fmt"

	"eaglebank/internal/shared/helpers"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	if !helpers.IsValidEmail(value) {
		return Email{}, fmt.Errorf("invalid email")
	}

	return Email{
		value: value,
	}, nil
}

func (u Email) AsString() string {
	return u.value
}
