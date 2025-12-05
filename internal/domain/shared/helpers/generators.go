package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	mathRand "math/rand"
)

const (
	transactionIdLen = 6
)

func GenerateNewUserId() (string, error) {
	// TODO this would be better to be a database SEQUENCE, just for simplicity, like I did for accountNumber
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes for I: %w", err)
	}

	suffix := hex.EncodeToString(randomBytes)

	return "usr-" + suffix, nil
}

func GenerateTransactionId() string {
	// TODO this would be better to be a database SEQUENCE, just for simplicity, like I did for accountNumber
	alphaNum := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	b := make([]byte, transactionIdLen)
	for i := range transactionIdLen {
		b[i] = alphaNum[mathRand.Intn(len(alphaNum))]
	}
	return "tan-" + string(b)
}
