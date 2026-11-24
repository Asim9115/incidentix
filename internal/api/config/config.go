package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_URL string
}

func Load() (*Config, error) {
	if err := godotenv.Load("../../.env"); err != nil {
		return nil, err
	}

	return &Config{
		POSTGRES_URL: getString("POSTGRES_URL", " "),
	}, nil
}

func getString(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}
