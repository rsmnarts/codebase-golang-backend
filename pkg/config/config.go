package config

import (
	"os"
	"strconv"
)

// Config holds application configuration
type Config struct {
	ServerPort int
	AppName    string
}

// Load loads configuration from environment variables
func Load() *Config {
	port := 8080
	if envPort := os.Getenv("PORT"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			port = p
		}
	}

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "Golang Fiber Backend"
	}

	return &Config{
		ServerPort: port,
		AppName:    appName,
	}
}
