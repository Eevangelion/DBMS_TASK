package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Address string
	Port    int
}

type Config struct {
	Server    ServerConfig
	Database  *Database
	DebugMode bool
}

var Conf *Config = nil

func init() {
	err := godotenv.Load("././.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func GetConfig() *Config {
	if Conf == nil {
		Conf = &Config{
			Server: ServerConfig{
				Address: getEnv("SERVER", "localhost"),
				Port:    getEnvAsInt("PORT", 6969),
			},
			Database:  SetupDB(),
			DebugMode: getEnvAsBool("DEBUG_MODE", true),
		}
	}
	return Conf
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
