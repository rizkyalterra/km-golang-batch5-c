package configs

import (
	"clean_architecture/drivers/mysql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitConfigDB() mysql.Config {
	return mysql.Config{
		DB_Username: os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}
}

func LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}