package config

import (
	configInterface "eaglebank/internal/infrastructure/config"

	"os"
)

func New() configInterface.Config {
	return &cfg{}
}

type cfg struct {
}

func (c *cfg) GetDBURL() string {
	return c.getenv("DATABASE_URL", "postgres://eaglebank:eaglebank123@localhost:5432/eaglebank?sslmode=disable")
}

func (c *cfg) GetJWTSecret() string {
	return c.getenv("JWT_SECRET", "local-dev-secret")
}

func (c *cfg) GetPort() string {
	return c.getenv("PORT", ":8080")
}

func (c *cfg) getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
