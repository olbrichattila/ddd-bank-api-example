// Package config, concrete implementation of config interface
package config

type Config interface {
	GetPort() string
	GetDBURL() string
	GetJWTSecret() string
}
