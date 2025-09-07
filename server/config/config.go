package config

import (
	"gippos-rat-server/internal/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db Database
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Warnf("Failed to load ENV file: %v", err)
	}

	cfg := Config{
		Db: Database{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", ""),
		},
	}

	return &cfg, nil
}

func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		if defaultVal == "" {
			logger.Log.Fatalf("Failed to get %s from environment", key)
		}
		return defaultVal
	}
	return val
}
