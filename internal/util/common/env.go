package common

import "os"

// EnvPort returns the PORT environment variable or default value
func EnvPort() string {
	return GetDefaultEnv("PORT", "8080")
}

func GetDefaultEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
