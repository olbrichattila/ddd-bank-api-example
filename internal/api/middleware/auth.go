package middleware

import (
	"context"
	"net/http"
)

type contextKey string

const UserIDKey contextKey = "userId"

func JWT(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// auth := r.Header.Get("Authorization")
			// if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			// 	http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			// 	return
			// }

			// tokenStr := strings.TrimPrefix(auth, "Bearer ")

			// token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			// 	return []byte(secret), nil
			// })
			// if err != nil || !token.Valid {
			// 	http.Error(w, "Invalid token", http.StatusUnauthorized)
			// 	return
			// }

			// claims := token.Claims.(jwt.MapClaims)

			// userID := claims["userId"].(string)

			userID := "usr-d187b52cf4ee97e05e65a7ebd4fd7ef7"

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
