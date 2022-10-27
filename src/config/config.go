package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Address string
	Port    string
}

type Config struct {
	Server    ServerConfig
	DebugMode bool
}

func New() *Config {
	return &Config{
		Server: ServerConfig{
			Address: getEnv("SERVER", "localhost"),
			Port:    getEnv("PORT", ""),
		},
		DebugMode: getEnvAsBool("DEBUG_MODE", true),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		value, err := strconv.Atoi(valueStr)
		if err == nil {
			return value
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if valueStr, exists := os.LookupEnv(key); exists {
		value, err := strconv.ParseBool(valueStr)
		if err == nil {
			return value
		}
	}
	return defaultValue
}
