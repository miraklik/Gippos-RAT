package storage

import (
	"fmt"
	"gippos-rat-server/config"
	"gippos-rat-server/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	cfg, err := config.Load()
	if err != nil {
		logger.Log.Fatalf("Failed to load config: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Db.Host, cfg.Db.Port, cfg.Db.Username, cfg.Db.Password, cfg.Db.Name)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Log.Printf("Failed to connect to database: %v", err)
		return &gorm.DB{}, fmt.Errorf("failed to connect to database: %v", err)
	}

	logger.Log.Info("Connected to database")
	return db, nil
}
