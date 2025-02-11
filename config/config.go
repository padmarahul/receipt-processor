package config

import (
	"os"
)

// Config struct stores application configurations
type Config struct {
	ServerPort string
	RedisAddr  string
}

// LoadConfig initializes configuration values. The default values is used to run in local workspace.
//The same keys are used as environment keys while building docker container.
func LoadConfig() Config {
	return Config{
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

// getEnv fetches environment variables with default values
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}