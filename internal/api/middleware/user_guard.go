package middleware

import (
	"eaglebank/internal/domain/shared/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserGuard() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Early validate user id format
			userId := chi.URLParam(r, string(UserIdKey))
			if ok := helpers.IsValidUserId(userId); !ok {
				http.Error(w, "The request didn't supply all the necessary data", http.StatusBadRequest)
				return
			}

			// Validate if user is logged in who owns the resource to protect against URL tampering
			requestUserId := chi.URLParam(r, "userId")
			loggedInUserId := r.Context().Value(UserIdKey)

			if requestUserId != loggedInUserId {
				http.Error(w, "The user is not allowed to access this resource", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
