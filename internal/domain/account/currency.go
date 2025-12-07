package account

import (
	"fmt"

	"eaglebank/internal/shared/helpers"
)

type Currency struct {
	value string
}

func NewCurrency(value string) (Currency, error) {
	if !helpers.IsValidCurrency(value) {
		return Currency{}, fmt.Errorf("invalid account type")
	}

	return Currency{
		value: value,
	}, nil
}

func (u Currency) AsString() string {
	return u.value
}
