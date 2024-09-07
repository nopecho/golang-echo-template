package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type Config struct {
	Postgres *PostgresConfig
	Port     int
}

var EnvConfig *Config

func init() {
	_ = godotenv.Load()
	postgres := initPostgresConfig()
	port := initOrDefaultPort(10000)

	EnvConfig = &Config{
		Postgres: postgres,
		Port:     port,
	}
}

func initOrDefaultPort(defaultValue int) int {
	strPort := getOrDefaultEnv("PORT", "10000")
	port, err := strconv.Atoi(strPort)
	if err != nil {
		port = defaultValue
	}
	return port
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
