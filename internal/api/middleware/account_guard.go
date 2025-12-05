package middleware

import (
	"eaglebank/internal/application/account"
	"eaglebank/internal/domain/shared/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	accountNumberURLParam = "accountNumber"
)

func AccountGuard(accountService account.Account) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accountNumber := chi.URLParam(r, string(accountNumberURLParam))
			if !helpers.IsValidAccountNumber(accountNumber) {
				http.Error(w, "The request didn't supply all the necessary data", http.StatusBadRequest)
				return
			}

			loggedInUserID := r.Context().Value(UserIDKey)
			if loggedInUserID == nil {
				http.Error(w, "The request didn't supply all the necessary data", http.StatusBadRequest)
				return
			}

			ok, err := accountService.BelongToUser(loggedInUserID.(string), accountNumber)
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
