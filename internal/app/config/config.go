package config

import (
	"github.com/joho/godotenv"
	"github.com/nopecho/golang-template/internal/app/infra/database"
	"os"
)

type EnvConfig struct {
	Postgres *database.ConnectionInfo
	Port     string
}

var Env *EnvConfig

func init() {
	_ = godotenv.Load()
	Env = &EnvConfig{
		Port: GetDefaultEnv("PORT", "10000"),
		Postgres: &database.ConnectionInfo{
			Host:     GetDefaultEnv("POSTGRES_HOST", "localhost"),
			Port:     GetDefaultEnv("POSTGRES_PORT", "5454"),
			Database: GetDefaultEnv("POSTGRES_DATABASE", "local"),
			Username: GetDefaultEnv("POSTGRES_USERNAME", "local"),
			Password: GetDefaultEnv("POSTGRES_PASSWORD", "local"),
		},
	}
}

func GetDefaultEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
