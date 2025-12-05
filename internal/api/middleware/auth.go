package middleware

import (
	"context"
	"eaglebank/internal/domain/shared/helpers"
	"net/http"
	"strings"
)

type contextKey string

const (
	UserIDKey contextKey = "userId"
	SecretKey string     = "your-secret-key" // TODO move it to config
)

func JWT(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			auth := r.Header.Get("Authorization")
			if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
				http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(auth, "Bearer ")

			userID, err := helpers.GetUserIDFromToken(tokenStr, []byte(SecretKey))
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// userID := "usr-d187b52cf4ee97e05e65a7ebd4fd7ef7"

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
