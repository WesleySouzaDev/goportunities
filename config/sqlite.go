package config

import (
	"goportunities/schemas"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create and initialize the database
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite database error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{}, &schemas.User{})
	if err != nil {
		logger.Errorf("sqlite migration error: %v", err)
		return nil, err
	}

	return db, nil
}
