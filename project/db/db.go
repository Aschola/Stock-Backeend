package db

import (
	"fmt"
	"log"
	"project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "schola:Newvera@764@tcp(127.0.0.1:3306)/stocksystem?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	DB = database
	fmt.Println("Database connected and migrated")
}

func GetDB() *gorm.DB {
	return DB
}
