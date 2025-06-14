package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"probe-challenge/model"
)

var DB *gorm.DB

func Connect() error {
	dsn := "host=localhost user=root password=root dbname=probe_challenge port=5433 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&model.Probe{})
	if err != nil {
		return fmt.Errorf("error doing AutoMigrate: %w", err)
	}

	DB = db
	return nil
}
