package account

import (
	"fmt"

	"atybank/internal/shared/helpers"
)

type AccountType struct {
	value string
}

func NewAccountType(value string) (AccountType, error) {
	if !helpers.IsValidAccountType(value) {
		return AccountType{}, fmt.Errorf("invalid account type")
	}

	return AccountType{
		value: value,
	}, nil
}

func (u AccountType) AsString() string {
	return u.value
}
