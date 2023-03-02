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

	GitHubClientID         string `mapstructure:"GITHUB_OAUTH_CLIENT_ID"`
	GitHubClientSecret     string `mapstructure:"GITHUB_OAUTH_CLIENT_SECRET"`
	GitHubOAuthRedirectUrl string `mapstructure:"GITHUB_OAUTH_REDIRECT_URL"`

	TokenLifeTime        int
	RefreshTokenLifeTime int
	PrivateKey           string
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
				Port:    getEnvAsInt("PORT", 8000),
			},
			Database:               SetupDB(),
			DebugMode:              getEnvAsBool("DEBUG_MODE", true),
			GitHubClientID:         getEnv("GITHUB_OAUTH_CLIENT_ID", ""),
			GitHubClientSecret:     getEnv("GITHUB_OAUTH_CLIENT_SECRET", ""),
			GitHubOAuthRedirectUrl: getEnv("GITHUB_OAUTH_REDIRECT_URL", ""),
			TokenLifeTime:          getEnvAsInt("TokenLifeTime", 15),
			RefreshTokenLifeTime:   getEnvAsInt("RefreshTokenLifeTime", 24),
			PrivateKey:             getEnv("PrivateKey", "YouMissSomething"),
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
