package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv         string
	APIKeyUsername string
	APIKeyPassword string
}

func Load() Config {
	log.Println("Loading configuration...")

	if godotenv.Load() != nil {
		log.Fatal("No .env file found")
	}

	cfg := Config{
		AppEnv:         os.Getenv("APP_ENV"),
		APIKeyUsername: os.Getenv("API_KEY_USERNAME"),
		APIKeyPassword: os.Getenv("API_KEY_PASSWORD"),
	}

	if cfg.APIKeyUsername == "" {
		panic("API_KEY_USERNAME not defined")
	}

	if cfg.APIKeyPassword == "" {
		panic("API_KEY_PASSWORD not defined")
	}

	if cfg.AppEnv == "" {
		cfg.AppEnv = "dev"
	}

	log.Println("Configuration loaded successfully")
	log.Println("App Environment: " + cfg.AppEnv)

	return cfg
}
