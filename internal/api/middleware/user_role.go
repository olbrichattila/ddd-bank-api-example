package middleware

import (
	"eaglebank/internal/domain/shared/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRole() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Early validate user id format
			userID := chi.URLParam(r, string(UserIDKey))
			if ok := helpers.IsValidUserID(userID); !ok {
				http.Error(w, "The request didn't supply all the necessary data", http.StatusBadRequest)
				return
			}

			// TODO comment back
			// // Validate if user is logged in who owns the resource to protect against URL tampering
			// requestUserID := chi.URLParam(r, "userId")
			// loggedInUserID := r.Context().Value(UserIDKey)

			// if requestUserID != loggedInUserID {
			// 	http.Error(w, "The user is not allowed to access the transaction", http.StatusForbidden)
			// 	return
			// }

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
