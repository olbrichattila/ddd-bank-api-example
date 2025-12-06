package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId string, secretKey []byte) (string, error) {
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(6 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func GetUserIdFromToken(tokenString string, secretKey []byte) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", jwt.ErrTokenMalformed
	}

	id, ok := claims["id"].(string)
	if !ok {
		return "", jwt.ErrTokenMalformed
	}

	return id, nil
}
