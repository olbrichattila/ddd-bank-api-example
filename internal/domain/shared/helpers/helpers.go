package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	mathRand "math/rand"
	"regexp"
	"slices"
)

var (
	userIDValidator        = regexp.MustCompile("^usr-[A-Za-z0-9]+$")
	accountNumberValidator = regexp.MustCompile(`^01\d{6}$`)
	transactionIDRegex     = regexp.MustCompile(`^tan-[A-Za-z0-9]$`)
	accountTypes           = []string{"personal", "business", "savings", "corporate"}
	transactionTypes       = []string{"deposit", "withdrawal"}
	currencyTypes          = []string{"GBP"}
	transactionIdLen       = 6
)

func GenerateNewUserID() (string, error) {
	// TODO this would be better to be a database SEQUENCE, just for simplicity, like I did for accountNumber
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes for ID: %w", err)
	}

	suffix := hex.EncodeToString(randomBytes)

	return "usr-" + suffix, nil
}

func GenerateTransactionID() string {
	// TODO this would be better to be a database SEQUENCE, just for simplicity, like I did for accountNumber
	alphaNum := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	b := make([]byte, transactionIdLen)
	for i := range transactionIdLen {
		b[i] = alphaNum[mathRand.Intn(len(alphaNum))]
	}
	return "tan-" + string(b)
}

func IsValidUserID(id string) bool {
	if id == "" {
		return false
	}
	return userIDValidator.MatchString(id)
}

func IsValidAccountNumber(accountNumber string) bool {
	return accountNumberValidator.MatchString(accountNumber)
}

func IsValidAccountType(accountType string) bool {
	return slices.Contains(accountTypes, accountType)
}

func IsValidTransactionType(transactionType string) bool {
	return slices.Contains(transactionTypes, transactionType)
}

func IsValidCurrency(currency string) bool {
	return slices.Contains(currencyTypes, currency)
}

func IsValidTransactionID(t string) bool {
	return transactionIDRegex.MatchString(t)
}
