package config

import (
	"Prioritas2/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	username := "s"
	password := "123ABC4d."
	host := "127.0.0.1"
	port := "3306"
	name := "km5"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username,
		password,
		host,
		port,
		name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Books{}, &models.Blog{})
}
