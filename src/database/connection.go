package database

import (
	"discord-clone/src/config"
	"discord-clone/src/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	config.DB = db
}
