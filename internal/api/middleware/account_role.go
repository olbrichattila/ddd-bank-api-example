package middleware

import (
	"eaglebank/internal/domain/shared/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	accountNumberURLParam = "accountNumber"
)

func AccountRole() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Early validate user id format
			accountNumber := chi.URLParam(r, string(accountNumberURLParam))
			if ok := helpers.IsValidAccountNumber(accountNumber); !ok {
				http.Error(w, "The request didn't supply all the necessary data", http.StatusBadRequest)
				return
			}

			// TODO check if this account number belongs to the user logged in
			// if ok := helpers.IsValidUserID(userID); !ok {
			// 	http.Error(w, "The user is not allowed to access the transaction", http.StatusForbidden)
			// 	return
			// }

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
