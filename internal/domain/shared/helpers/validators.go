package helpers

import (
	"regexp"
	"slices"
)

var (
	userIDValidator        = regexp.MustCompile("^usr-[A-Za-z0-9]+$")
	accountNumberValidator = regexp.MustCompile(`^01\d{6}$`)
	transactionIDRegex     = regexp.MustCompile(`^tan-[A-Za-z0-9]+$`)
	paymentAmount          = regexp.MustCompile(`^(?:10000|[0-9]{1,4})(?:\.[0-9]{1,2})?$`)
	emailRegex             = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	accountTypes     = []string{"personal", "business", "savings", "corporate"}
	transactionTypes = []string{"deposit", "withdrawal"}
	currencyTypes    = []string{"GBP"}
)

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

func IsValidTransactionID(transactionId string) bool {
	return transactionIDRegex.MatchString(transactionId)
}

func IsValidPaymentAmount(amountAsStr string) bool {
	return paymentAmount.MatchString(amountAsStr)
}

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
