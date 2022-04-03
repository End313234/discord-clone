package database

import (
	"discord-clone/src/config"
	"discord-clone/src/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	config.DB = db
}
