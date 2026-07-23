package handler

import (
	"goportunities/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() (db *gorm.DB, logger *config.Logger) {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()

	return db, logger
}
