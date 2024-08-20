package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AUTH_PORT    string
	MEDICAL_PORT string
	API_GETEWAY  string

	TokenKey string

	LOG_PATH string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.AUTH_PORT = cast.ToString(coalesce("AUTH_PORT", ":8085"))
	config.MEDICAL_PORT = cast.ToString(coalesce("MEDICAL_PORT", ":8088"))
	config.API_GETEWAY = cast.ToString(coalesce("gateway", ":55555"))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	config.TokenKey = cast.ToString(coalesce("TokenKey", "my_secret_key"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
