package middleware

import (
	"context"
	"eaglebank/internal/domain/shared/helpers"
	"net/http"
	"strings"
)

type contextKey string

const (
	UserIdKey contextKey = "userId"
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

			userId, err := helpers.GetUserIdFromToken(tokenStr, []byte(secret))
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIdKey, userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
