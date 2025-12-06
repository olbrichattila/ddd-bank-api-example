package user

import (
	"fmt"

	"eaglebank/internal/shared/helpers"
)

type Phone struct {
	value string
}

func NewPhoneNumber(value string) (Phone, error) {
	if !helpers.IsValidPhone(value) {
		return Phone{}, fmt.Errorf("invalid phone")
	}

	return Phone{
		value: value,
	}, nil
}

func (u Phone) AsString() string {
	return u.value
}
