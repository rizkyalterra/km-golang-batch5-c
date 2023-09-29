package mysql

import (
	"clean_architecture/drivers/mysql/user"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB(config Config) *gorm.DB {

	// config := Config{
	// 	DB_Username: os.Getenv("DB_USERNAME"),
	// 	DB_Password: os.Getenv("DB_PASSWORD"),
	// 	DB_Port:     os.Getenv("DB_PORT"),
	// 	DB_Host:     os.Getenv("DB_HOST"),
	// 	DB_Name:     os.Getenv("DB_NAME"),
	// }

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	initMigrate(db)

	return db
}

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
