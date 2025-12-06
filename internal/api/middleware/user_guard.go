package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"eaglebank/internal/shared/helpers"
	"eaglebank/internal/shared/services"
)

func UserGuard(logger services.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Early validate user id format
			userId := chi.URLParam(r, string(UserIdKey))
			if logger != nil {
				logger.Debug("userIdFromParam: " + userId)
			}
			if ok := helpers.IsValidUserId(userId); !ok {
				http.Error(w, "The request didn't supply all the necessary data", http.StatusBadRequest)
				return
			}

			// Validate if user is logged in who owns the resource to protect against URL tampering
			loggedInUserId := r.Context().Value(UserIdKey)
			loggedUserIdAsStr, ok := loggedInUserId.(string)
			if !ok {
				http.Error(w, "The request didn't supply all the necessary data", http.StatusBadRequest)
				return
			}

			if logger != nil {
				logger.Debug("loggedUserId: " + loggedUserIdAsStr)
			}

			if userId != loggedUserIdAsStr {
				http.Error(w, "The user is not allowed to access this resource", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
