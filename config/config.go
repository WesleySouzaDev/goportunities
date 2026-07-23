package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	// Load environment variables
	LoadEnv()

	// Initialize database
	_, err := InitializeSqlite()

	if err != nil {
		return err
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// Initialize logger
	logger = NewLogger(prefix)
	return logger
}
