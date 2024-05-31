package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connecting db: %w", err)
	}

	log.Println("Database connected")

	if err := db.AutoMigrate(&Student{}, &Plan{}); err != nil {
		return fmt.Errorf("migrating db: %w", err)
	}

	log.Println("Database migrated")

	DB = db
	return nil
}
