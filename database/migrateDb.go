package database

import "myproject/models"

func MigrateDb() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}