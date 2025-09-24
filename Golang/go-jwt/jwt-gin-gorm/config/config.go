package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port         string
		Host         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}

	JWT struct {
		Secret        string
		AccessExpiry  time.Duration
		RefreshExpiry time.Duration
	}

	Environment string
}

// Load env values
func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error while loading env: ", err)
	}

	cfg := &Config{}
	//server
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")
	cfg.Server.Host = getEnv("SERVER_HOST", "0.0.0.0")
	cfg.Server.ReadTimeout = time.Second * 15
	cfg.Server.WriteTimeout = time.Second * 15

	// Database
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "ak")
	cfg.Database.Password = getEnv("DB_PASSWORD", "4455")
	cfg.Database.DBName = getEnv("DB_NAME", "users_db")
	cfg.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	// JWT
	cfg.JWT.Secret = getEnv("JWT_SECRET", "your-secret-key")
	cfg.JWT.AccessExpiry = time.Minute * 1  // 5 minutes
	cfg.JWT.RefreshExpiry = time.Hour * 168 // 7 days

	cfg.Environment = getEnv("ENV", "development")

	return cfg, nil
}

// Get env variables
func getEnv(key, defaultValue string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// DB dsn
func (c *Config) GetDSN() string {

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode)
}

func (c *Config) ServerAddress() string {
	return c.Server.Host + ":" + c.Server.Port
}
