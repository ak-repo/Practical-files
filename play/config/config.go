package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN        string
	SecrectKey string
}

func Load() (*Config, error) {

	// load env
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	//
	dsn := os.Getenv("DSN")
	secret := os.Getenv("SECRET_KEY")
	if dsn == "" || secret == "" {
		return nil, errors.New("env file issue")
	}

	return &Config{DSN: dsn, SecrectKey: secret}, nil
}
