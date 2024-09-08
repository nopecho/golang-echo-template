package config

import (
	"github.com/joho/godotenv"
	"os"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type EnvConfig struct {
	Postgres *PostgresConfig
	Port     string
}

var Env *EnvConfig

func init() {
	_ = godotenv.Load()
	postgres := initPostgresConfig()
	port := getOrDefaultEnv("PORT", "10000")

	Env = &EnvConfig{
		Postgres: postgres,
		Port:     port,
	}
}

func initPostgresConfig() *PostgresConfig {
	postgres := &PostgresConfig{
		Host:     getOrDefaultEnv("POSTGRES_HOST", "localhost"),
		Port:     getOrDefaultEnv("POSTGRES_PORT", "5432"),
		Database: getOrDefaultEnv("POSTGRES_DATABASE", "postgres"),
		Username: getOrDefaultEnv("POSTGRES_USERNAME", "postgres"),
		Password: getOrDefaultEnv("POSTGRES_PASSWORD", "postgres"),
	}
	return postgres
}

func getOrDefaultEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
