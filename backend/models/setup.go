package models

import (
	"movie-booking/models"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  func ConnectDatabase() *gorm.DB{
	
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.User{})
	return db
}