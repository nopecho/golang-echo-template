package common

import "os"

// EnvPort returns the PORT environment variable or default value
func EnvPort() string {
	return GetEnv("PORT", "8080")
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
