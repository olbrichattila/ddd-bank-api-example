package helpers

import (
	"regexp"
	"slices"
)

var (
	userIdValidator        = regexp.MustCompile("^usr-[A-Za-z0-9]+$")
	accountNumberValidator = regexp.MustCompile(`^01\d{6}$`)
	transactionIdRegex     = regexp.MustCompile(`^tan-[A-Za-z0-9]+$`)
	paymentAmount          = regexp.MustCompile(`^(?:10000|[0-9]{1,4})(?:\.[0-9]{1,2})?$`)
	emailRegex             = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	phoneRegex             = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

	accountTypes     = []string{"personal", "business", "savings", "corporate"}
	transactionTypes = []string{"deposit", "withdrawal"}
	currencyTypes    = []string{"GBP"}
)

func IsValidUserId(id string) bool {
	if id == "" {
		return false
	}
	return userIdValidator.MatchString(id)
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

func IsValidTransactionId(transactionId string) bool {
	return transactionIdRegex.MatchString(transactionId)
}

func IsValidPaymentAmount(amountAsStr string) bool {
	return paymentAmount.MatchString(amountAsStr)
}

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsValidPhone(phoneNumber string) bool {
	return phoneRegex.MatchString(phoneNumber)
}
