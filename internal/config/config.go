package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL     string
	ServerPort      string
	AppEnv          string
	ServerHost      string
	DbHost          string
	DbPort          string
	DbUser          string
	DbPassword      string
	DbName          string
	DbSSLMode       string
	TokenSecret     string
	TokenTTLMinutes int
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ttl, err := strconv.Atoi(os.Getenv("TOKEN_TTL_MINUTES"))
	if err != nil {
		log.Fatal("Error parsing TOKEN_TTL_MINUTES")
	}

	var config *Config = &Config{
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		ServerPort:      os.Getenv("SERVER_PORT"),
		AppEnv:          os.Getenv("APP_ENV"),
		ServerHost:      os.Getenv("SERVER_HOST"),
		DbHost:          os.Getenv("DB_HOST"),
		DbPort:          os.Getenv("DB_PORT"),
		DbUser:          os.Getenv("DB_USER"),
		DbPassword:      os.Getenv("DB_PASSWORD"),
		DbName:          os.Getenv("DB_NAME"),
		DbSSLMode:       os.Getenv("DB_SSL_MODE"),
		TokenSecret:     os.Getenv("TOKEN_SECRET"),
		TokenTTLMinutes: ttl,
	}
	return config, nil
}
