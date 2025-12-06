package account

import (
	"fmt"

	"eaglebank/internal/shared/helpers"
)

type AccountNumber struct {
	value string
}

func NewAccountNumber(value string) (AccountNumber, error) {
	if !helpers.IsValidAccountNumber(value) {
		return AccountNumber{}, fmt.Errorf("invalid account number")
	}

	return AccountNumber{
		value: value,
	}, nil
}

func (u AccountNumber) AsString() string {
	return u.value
}
