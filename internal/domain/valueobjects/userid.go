package valueobjects

import (
	"fmt"

	"atybank/internal/shared/helpers"
)

type UserId struct {
	value string
}

func NewUserId(value string) (UserId, error) {
	if value != "" && !helpers.IsValidUserId(value) {
		return UserId{}, fmt.Errorf("invalid user id")
	}

	return UserId{
		value: value,
	}, nil
}

func (u UserId) AsString() string {
	return u.value
}
