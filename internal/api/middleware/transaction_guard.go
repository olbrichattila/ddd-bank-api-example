package middleware

import (
	"eaglebank/internal/application/transaction"
	"eaglebank/internal/domain/shared/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	transactionIdURLParam = "transactionNumber"
)

func TransactionGuard(transactionService transaction.Transaction) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			transactionId := chi.URLParam(r, string(transactionIdURLParam))
			if !helpers.IsValidTransactionId(transactionId) {
				http.Error(w, "The request didn't supply all the necessary data", http.StatusBadRequest)
				return
			}

			accountNumber := chi.URLParam(r, string(accountNumberURLParam))
			if !helpers.IsValidAccountNumber(accountNumber) {
				http.Error(w, "The request didn't supply all the necessary data", http.StatusBadRequest)
				return
			}

			ok, err := transactionService.BelongToUser(accountNumber, transactionId)
			if err != nil {
				http.Error(w, "unexpected error occurred", http.StatusInternalServerError)
				return
			}

			if !ok {
				http.Error(w, "The user is not allowed to access this resource", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
