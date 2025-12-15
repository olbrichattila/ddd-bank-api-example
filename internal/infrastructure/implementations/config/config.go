package config

import (
	"fmt"
	"os"

	configInterface "atybank/internal/infrastructure/config"

	"github.com/joho/godotenv"
)

const dbEnvFileName = ".env.migrator"

func New() configInterface.Config {
	c := &cfg{}
	c.loadEnvIfExists(dbEnvFileName)

	return c
}

type cfg struct {
}

func (c *cfg) GetDBURL() string {
	return c.getenv(
		"DATABASE_URL",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			c.GetDBUserName(),
			c.GetDBPassword(),
			c.GetDBHost(),
			c.GetGetDBPort(),
			c.GetDBDatabase(),
		),
	)
}

func (c *cfg) GetJWTSecret() string {
	return c.getenv("JWT_SECRET", "local-dev-secret")
}

func (c *cfg) GetPort() string {
	return c.getenv("PORT", ":8080")
}

func (c *cfg) GetDBHost() string {
	return c.getenv("DB_HOST", "127.0.0.1")
}

func (c *cfg) GetGetDBPort() string {
	return c.getenv("DB_PORT", "5432")
}

func (c *cfg) GetDBDatabase() string {
	return c.getenv("DB_DATABASE", "atybank")
}

func (c *cfg) GetDBUserName() string {
	return c.getenv("DB_USERNAME", "atybank")
}

// GetDBPassword defaults to a test password, use app without .env file and pass variables via linux env variables
// for example AWS Secret Manager, Ansible, ....
func (c *cfg) GetDBPassword() string {
	return c.getenv("DB_PASSWORD", "atybank123")
}

func (c *cfg) getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func (c *cfg) loadEnvIfExists(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("failed to stat %s: %w", path, err)
	}

	if !info.Mode().IsRegular() {
		return fmt.Errorf("%s exists but is not a regular file", path)
	}

	if err := godotenv.Load(path); err != nil {
		return fmt.Errorf("failed to load %s: %w", path, err)
	}

	return nil
}
