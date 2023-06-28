package config

import (
	"github.com/feebrian/gotodolist/app/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("database/todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&entities.Todo{},
		&entities.User{},
	)
	if err != nil {
		panic(err)
	}

	DB = db
}
